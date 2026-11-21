// Package unsafeaudit finds the one unsafe pattern the transpiled libopus code
// must never contain: the address of a function-local variable laundered
// through uintptr(unsafe.Pointer(&x)).
//
// A local whose address only ever becomes a uintptr does not escape, so it may
// live on the goroutine stack — and the transpiled code's deep call chains can
// grow, i.e. MOVE, that stack mid-call. Go rewrites real pointers when a stack
// moves; it cannot rewrite a uintptr, so the callee keeps reading and writing
// the abandoned copy. Under -race (different stack sizes) two decoder tests
// return wrong values for exactly this reason.
//
// The check is syntactic: it cannot tell `&p.f` through a pointer local (safe)
// from `&s.f` on a struct local (unsafe), so it over-reports and a zero is
// strong evidence. The encoder package is held at zero by a guard test.
package unsafeaudit

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
)

// Finding is one laundered local address.
type Finding struct {
	Pos      token.Position
	Local    string
	Function string
}

func (f Finding) String() string {
	return fmt.Sprintf("%s: &%s in %s", f.Pos, f.Local, f.Function)
}

// Dir audits every non-test .go file in dir.
func Dir(dir string) ([]Finding, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		return nil, err
	}
	fset := token.NewFileSet()
	var out []Finding
	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") {
			continue
		}
		af, err := parser.ParseFile(fset, f, nil, 0)
		if err != nil {
			return nil, err
		}
		for _, d := range af.Decls {
			if fd, ok := d.(*ast.FuncDecl); ok && fd.Body != nil {
				out = append(out, auditFunc(fset, fd)...)
			}
		}
	}
	return out, nil
}

func auditFunc(fset *token.FileSet, fd *ast.FuncDecl) []Finding {
	locals := map[*ast.Object]bool{} //nolint:staticcheck // SA1019: ast.Object is the only resolver available without type-checking; the audit is syntactic by design
	if fd.Type.Params != nil {
		for _, p := range fd.Type.Params.List {
			for _, id := range p.Names {
				locals[id.Obj] = true
			}
		}
	}
	ast.Inspect(fd.Body, func(n ast.Node) bool {
		switch v := n.(type) {
		case *ast.ValueSpec:
			for _, id := range v.Names {
				locals[id.Obj] = true
			}
		case *ast.AssignStmt:
			if v.Tok == token.DEFINE {
				for _, l := range v.Lhs {
					if id, ok := l.(*ast.Ident); ok && id.Obj != nil {
						locals[id.Obj] = true
					}
				}
			}
		}
		return true
	})

	var out []Finding
	ast.Inspect(fd.Body, func(n ast.Node) bool {
		u := launderedAddr(n)
		if u == nil {
			return true
		}
		if id := rootIdent(u.X); id != nil && id.Obj != nil && locals[id.Obj] {
			out = append(out, Finding{Pos: fset.Position(u.Pos()), Local: id.Name, Function: fd.Name.Name})
		}
		return true
	})
	return out
}

// launderedAddr matches uintptr(unsafe.Pointer(&x...)) and returns the &x.
func launderedAddr(n ast.Node) *ast.UnaryExpr {
	outer, ok := n.(*ast.CallExpr)
	if !ok || len(outer.Args) != 1 {
		return nil
	}
	if id, ok := outer.Fun.(*ast.Ident); !ok || id.Name != "uintptr" {
		return nil
	}
	inner, ok := outer.Args[0].(*ast.CallExpr)
	if !ok || len(inner.Args) != 1 {
		return nil
	}
	sel, ok := inner.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "Pointer" {
		return nil
	}
	if pkg, ok := sel.X.(*ast.Ident); !ok || pkg.Name != "unsafe" {
		return nil
	}
	u, ok := inner.Args[0].(*ast.UnaryExpr)
	if !ok || u.Op != token.AND {
		return nil
	}
	return u
}

func rootIdent(x ast.Expr) *ast.Ident {
	for {
		switch e := x.(type) {
		case *ast.IndexExpr:
			x = e.X
		case *ast.SelectorExpr:
			x = e.X
		case *ast.ParenExpr:
			x = e.X
		case *ast.Ident:
			return e
		default:
			return nil
		}
	}
}
