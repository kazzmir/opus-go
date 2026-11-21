# THE local CI gate. GitHub Actions are disabled on this repo (owner, 2026-11-21); `make ci-local`
# must pass before every merge and every release tag. It replaces:
#   - .github/workflows/go-test.yml: `go test ./...` on linux/amd64, linux/arm64, macos, plus
#     `GOARCH=386 go test ./...`. The arm64 and macos runners cannot be reproduced on this host; the
#     closest honest substitute is the cross-compile `go vet` steps below (type-check + build of every
#     package for darwin/arm64 and linux/arm64, tests included). linux/amd64 and 386 tests run for real.
# It also adds (new for this repo, no workflow existed): the full-history gitleaks scan (scripts/ci-gitleaks.sh).
# Each step is its own $(MAKE) so the order is fail-fast even under `make -j`.

.PHONY: ci-local gitleaks fmt-check mod-tidy-check govulncheck vet lint test test-386 vet-cross

ci-local:
	$(MAKE) gitleaks
	$(MAKE) fmt-check
	$(MAKE) mod-tidy-check
	$(MAKE) govulncheck
	$(MAKE) vet
	$(MAKE) lint
	$(MAKE) test
	$(MAKE) test-386
	$(MAKE) vet-cross
	@echo "ci-local: all gates passed"

gitleaks:
	./scripts/ci-gitleaks.sh

fmt-check:
	@unformatted=$$(gofmt -l .) || exit 1; \
	if [ -n "$$unformatted" ]; then echo "not gofmt'd:"; echo "$$unformatted"; exit 1; fi

mod-tidy-check:
	go mod tidy -diff

govulncheck:
	govulncheck ./...

# This module is a C-to-Go transpilation that carries C pointers as uintptr by design, so vet's `unsafeptr`
# analyzer flags thousands of intended sites (libcshim, opus, opuscc, opusccenc); that discipline is held by the
# internal/unsafeaudit tests under `make test`. The transpiled packages (opuscc, opusccenc) additionally
# trip `unreachable` and `assign` on generated control flow. Every other vet analyzer runs everywhere.
VET_TRANSPILED := ./opuscc/... ./opusccenc/...
VET_HAND_WRITTEN := $(shell go list ./... | grep -v -E '/(opuscc|opusccenc)$$')

vet:
	go vet -unsafeptr=false $(VET_HAND_WRITTEN)
	go vet -unsafeptr=false -unreachable=false -assign=false $(VET_TRANSPILED)

lint:
	golangci-lint run ./...

# -race is added over the retired workflow's plain `go test ./...`; -count=1 defeats the test cache.
test:
	go test -race -count=1 ./...

# The retired workflow's 32-bit run (-race is unsupported on 386, so none here).
test-386:
	GOARCH=386 go test -count=1 ./...

# Substitute for the arm64 and macos runners that cannot run here: type-check every package, tests included.
vet-cross:
	GOOS=darwin GOARCH=arm64 $(MAKE) vet
	GOOS=linux GOARCH=arm64 $(MAKE) vet
