#!/usr/bin/env bash
# Purpose: VAI-SEC-002 secret scan, run locally by `make ci-local` (GitHub Actions are disabled
# on vAudience repos since 2026-11-21; this replaces .github/workflows/security.yml).
# Same pinned binary as the retired workflow: gitleaks 8.30.1. The release tarball AND the extracted
# binary are pinned by sha256; the cached binary is re-verified on every run.
# Scope: the full history of every LOCAL ref (`gitleaks git`), so it is stricter than the workflow's
# checkout: unpushed local branches are scanned too. A shallow clone is refused (the workflow forced
# fetch-depth: 0), because a truncated history would scan clean.
# Accepted risk: the binary is verified, then executed; something already running as this user with write
# access to ~/.cache could swap it in between (it could equally edit the Makefile).
# A finding is never "fixed" by rewriting history: rotate the secret, then add its fingerprint to
# .gitleaksignore with a comment naming the rotation (or the reason it is a false positive).
set -euo pipefail

GITLEAKS_VERSION="8.30.1"
GITLEAKS_TGZ_SHA256_LINUX_X64="551f6fc83ea457d62a0d98237cbad105af8d557003051f41f3e7ca7b3f2470eb"
GITLEAKS_BIN_SHA256_LINUX_X64="88f91962aa2f93ac6ab281d553b9e125f5197bbbce38f9f2437f7299c32e5509"

case "$(uname -s)-$(uname -m)" in
  Linux-x86_64) ;;
  *) echo "ci-gitleaks: only linux/x86_64 has a pinned checksum; refusing to run an unverified binary" >&2; exit 2 ;;
esac

for tool in git curl tar sha256sum install shuf base64; do
  command -v "${tool}" >/dev/null || { echo "ci-gitleaks: required tool '${tool}' is missing" >&2; exit 2; }
done

root="$(git rev-parse --show-toplevel)"
if [[ "$(git -C "${root}" rev-parse --is-shallow-repository)" != "false" ]]; then
  echo "ci-gitleaks: shallow clone; the scan needs the full history. Run: git fetch --unshallow" >&2
  exit 2
fi

# The scan's configuration must not be weakened from outside the reviewed tree: the environment cannot redirect
# it (only the source-root .gitleaks.toml is read, gitleaks 8.30.1 has no other auto-loaded name), and a repo config
# may not disable rules (gitleaks reads keys case-insensitively, so refuse any spelling). Whether the EFFECTIVE config
# still detects is proven BY USE below (canaries); allowlists stay possible as reviewed changes like .gitleaksignore.
unset GITLEAKS_CONFIG GITLEAKS_CONFIG_TOML
cfg="${root}/.gitleaks.toml"
if [[ -e "${cfg}" ]] && grep -qiE 'disabled.?rules|disabled\\u0052ules' "${cfg}"; then
  echo "ci-gitleaks: ${cfg##*/} may not disable rules (any spelling of disabledRules)" >&2
  exit 2
fi

cache="${XDG_CACHE_HOME:-$HOME/.cache}/vai-ci/gitleaks-${GITLEAKS_VERSION}"
bin="${cache}/gitleaks"

verify_bin() { echo "${GITLEAKS_BIN_SHA256_LINUX_X64}  ${bin}" | sha256sum -c - >/dev/null 2>&1; }

if [[ ! -x "${bin}" ]] || ! verify_bin; then
  mkdir -p "${cache}"
  tmp="$(mktemp -d)"
  trap 'rm -rf "${tmp}" "${bin}.tmp.$$"' EXIT INT TERM
  url="https://github.com/gitleaks/gitleaks/releases/download/v${GITLEAKS_VERSION}/gitleaks_${GITLEAKS_VERSION}_linux_x64.tar.gz"
  if ! curl -fsSL --max-time 120 --retry 3 -o "${tmp}/gitleaks.tgz" "${url}" \
    || ! echo "${GITLEAKS_TGZ_SHA256_LINUX_X64}  ${tmp}/gitleaks.tgz" | sha256sum -c - >/dev/null \
    || ! tar -xzf "${tmp}/gitleaks.tgz" -C "${tmp}" gitleaks \
    || ! echo "${GITLEAKS_BIN_SHA256_LINUX_X64}  ${tmp}/gitleaks" | sha256sum -c - >/dev/null; then
    rm -rf "${tmp}"
    echo "ci-gitleaks: download or checksum verification failed" >&2
    exit 2
  fi
  # Atomic: a concurrent or interrupted install never leaves a half-written executable.
  install -m 0755 "${tmp}/gitleaks" "${bin}.tmp.$$"
  mv -f "${bin}.tmp.$$" "${bin}"
  rm -rf "${tmp}"
  verify_bin || { echo "ci-gitleaks: installed binary failed its checksum" >&2; exit 2; }
fi

"${bin}" version >/dev/null

# Canary BY USE: the effective config (defaults, or the repo's .gitleaks.toml) must still report one planted secret per
# default rule below, in the SAME scan mode and path shape as the real scan (git mode, repo-relative paths). Each value is
# random per run, so a config cannot be written to miss exactly the published canaries. This catches configs that load
# no rules, same-id [[rules]] overrides, and allowlists or stopwords that match these rules, values or path shapes. A
# repo-specific path allowlist (e.g. ^pkg/) still passes: that stays a reviewed allowlist, like .gitleaksignore and
# gitleaks:allow. Literals are assembled from pieces so the repo's own history scan never flags this file; the scan runs
# from the canary dir, so the repo's .gitleaksignore cannot suppress a canary.
rnd() { local r; r="$(head -c 4096 /dev/urandom | LC_ALL=C tr -dc "$1")"; printf '%s' "${r:0:$2}"; }
canary_cfg=()
[[ -e "${cfg}" ]] && canary_cfg=(--config "${cfg}")
canary_dir=""
trap 'rm -rf "${canary_dir}"' EXIT
trap 'rm -rf "${canary_dir}"; exit 2' INT TERM
# run_canary <dir> prints the rules whose canary was NOT reported (empty = all detected); returns 2 if it cannot run.
# The PARENT creates and removes <dir> (a failed mktemp exits there under set -e); run_canary refuses any dir that is
# not a fresh directory outside the scanned repo, so a git init/add/commit can never land in the caller's repository.
run_canary() {
  local d="${1:-}" k L D i rule missed=""
  [[ -n "${d}" && -d "${d}" && "${d}" == /* ]] || return 2
  case "$(cd -- "${d}" && pwd -P)/" in "$(cd -- "${root}" && pwd -P)/"*) return 2 ;; esac
  [[ -z "$(ls -A -- "${d}")" ]] || return 2
  mkdir -p "${d}/internal/app" "${d}/cmd/svc" "${d}/web/src"
  local p_gh='gh''p_' p_aws='AK''IA' p_key='PRIVATE'' KEY' p_slack='xo''xb' p_api='api''_key' s1 s2 s3
  printf 'package app\n\nconst tok = "%s%s"\n' "${p_gh}" "$(rnd 'A-Za-z0-9' 36)" > "${d}/internal/app/config.go"
  # 16 DISTINCT characters keep the value above the rule's entropy floor every run (repeats made ~3/10,000 miss).
  printf 'aws_access_key_id = %s%s\n' "${p_aws}" "$(printf '%s\n' {A..Z} 2 3 4 5 6 7 | shuf -n 16 | tr -d '\n')" > "${d}/cmd/svc/creds.ini"
  printf -- '-----BEGIN RSA %s-----\n%s\n-----END RSA %s-----\n' "${p_key}" "$(head -c 96 /dev/urandom | base64 -w0)" "${p_key}" > "${d}/cmd/svc/id_rsa"
  s1="$(rnd '0-9' 12)"; s2="$(rnd '0-9' 13)"; s3="$(rnd 'A-Za-z0-9' 24)"
  printf 'const slack = "%s-%s-%s-%s";\n' "${p_slack}" "${s1}" "${s2}" "${s3}" > "${d}/web/src/index.ts"
  # 14 DISTINCT letters interleaved with digits: a digit is guaranteed, no stopword can form and the entropy stays high,
  # so the generic rule fires every run.
  L="$(printf '%s\n' {A..Z} {a..z} | shuf -n 14 | tr -d '\n')"; D="$(rnd '0-9' 14)"; k=""
  for i in {0..13}; do k+="${L:i:1}${D:i:1}"; done
  printf '%s = "%s"\n' "${p_api}" "${k}" > "${d}/internal/app/settings.toml"
  if ! (cd -- "${d}" || exit 2
        # A caller's git environment (a hook sets GIT_DIR / GIT_INDEX_FILE / GIT_WORK_TREE, …) would override the cd and
        # send the canary's init/add/commit into the caller's repository: clear every per-repo git variable first.
        # shellcheck disable=SC2046
        # The user's own git config, template and HOME must not reach the canary repo either: a template can share the
        # caller's refs (commondir), a global config can run hooks or an fsmonitor. Run with none of them.
        unset $(git rev-parse --local-env-vars) GIT_TEMPLATE_DIR XDG_CONFIG_HOME
        export GIT_CONFIG_GLOBAL=/dev/null GIT_CONFIG_NOSYSTEM=1 HOME="${d}"
        git init -q --template= . >/dev/null \
        && git add -A . >/dev/null \
        && git -c user.email=c@c -c user.name=c -c core.hooksPath=/dev/null -c commit.gpgsign=false commit -q --no-verify -m c >/dev/null \
        && "${bin}" git --no-banner --log-level error ${canary_cfg[@]+"${canary_cfg[@]}"} \
             --report-format json --report-path "${d}/.report.json" --exit-code 0 "${d}" >/dev/null 2>&1); then
    return 2
  fi
  for rule in github-pat aws-access-token private-key slack-bot-token generic-api-key; do
    grep -q "\"RuleID\": *\"${rule}\"" "${d}/.report.json" || missed+="${rule} "
  done
  printf '%s' "${missed}"
}
# A real weakening misses deterministically; a random value may (rarely) miss once. Retry once with fresh values.
missed=""
for _ in 1 2; do
  canary_dir="$(mktemp -d)"
  missed="$(run_canary "${canary_dir}")" || { echo "ci-gitleaks: canary scan failed to run" >&2; exit 2; }
  rm -rf "${canary_dir}"; canary_dir=""
  [[ -z "${missed}" ]] && break
done
if [[ -n "${missed}" ]]; then
  echo "ci-gitleaks: the effective config no longer detects: ${missed}(canary not reported twice): the scan is weakened" >&2
  exit 2
fi
trap - EXIT INT TERM

exec "${bin}" git --redact --no-banner --exit-code 1 "${root}"
