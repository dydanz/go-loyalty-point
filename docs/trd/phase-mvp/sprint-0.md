# Sprint 0 — Stabilize & Secure

> **Phase:** MVP · **Duration:** 2 weeks · **TRD anchor:** Phase 0 (§5), §2.3 C1–C4 + S8, §6.1
> **Gate:** HARD — published build is unsafe until this lands. No member-facing work ships before sign-off.

## Goal

Make the system safe to build on. Kill the P0 security defects, stop silent error-swallowing, give CI teeth (`-race` + vuln scan).

## Why first

§2.3 lists four 🔴 Critical defects: creds in the published image, a public endpoint leaking bcrypt hashes, an unimplemented `LogError` swallowing all 5xx, a startup routine wiping every Redis session. Fix before building anything.

**Out of scope:** rule engine, ledger, member UI.

---

## Task 0.1 — Secrets out of image (SOPS)

**Type:** backend / infra · **Effort:** M · **TRD:** C1, §3.4, §6.1, §7
**Progress signal:** `docker run --rm <img> cat /app/.env` fails — image provably clean.

**Design**
- Threat: any image puller reads DB/Redis creds. Remove secrets from build context; inject at runtime only.
- Choose SOPS with `age` key (POC simplicity; KMS later). One encrypted file per env (`config/secrets.<env>.enc.yaml`), decrypted into env / K8s Secrets at deploy.

**Implementation**
- Remove `.env` COPY from `Dockerfile` (TRD cites `Dockerfile:37`); add `.env` to `.dockerignore`.
- Add SOPS encrypted config + decrypt step in deploy/entrypoint; app reads from env (`pkg/config`).
- **Rotate every credential** shipped in a published image (DB, Redis); audit GHCR pull logs if available.

**Testing**
- Image test (CI): assert `.env` absent and no plaintext secret strings in layers.
- Boot test: app starts from injected env only (no `.env` file present).
- Negative: missing required secret → fail-fast with clear error (per CLAUDE.md config rule).

---

## Task 0.2 — Remove hash-leak endpoint

**Type:** backend · **Effort:** S · **TRD:** C2
**Progress signal:** `/api/auth/test/random-user` returns 404; no bcrypt hash retrievable unauthenticated.

**Design**
- Endpoint (`internal_load_test_handler.go` `GetRandomVerifiedUser`) exposes password hashes for offline cracking. Prefer outright removal for MVP; if load-test fixture needed, internal-gate + drop password field.

**Implementation**
- Delete route registration + handler. Remove now-dead repo method if unused.

**Testing**
- Integration: route returns 404.
- Grep/contract test: no handler returns a `password`/hash field in any response body.

---

## Task 0.3 — Stop startup session-wipe

**Type:** backend · **Effort:** S · **TRD:** C4
**Progress signal:** live session token still valid after a rolling restart.

**Design**
- `DeleteAllSession` at boot (TRD cites `cmd/api/main.go:76`) force-logs-out all users on every deploy. Remove; rely on Redis TTL for expiry.

**Implementation**
- Remove the startup call. Confirm session TTL set on write.

**Testing**
- Restart test: create session → restart process → assert token still authenticates.

---

## Task 0.4 — Implement `LogError`

**Type:** backend · **Effort:** S · **TRD:** C3, §6.2
**Progress signal:** triggered 5xx appears as structured JSON log with trace id + route.

**Design**
- Stub at `server/util/error_handler.go:116` swallows 5xx → zero prod visibility. Replace with structured zerolog emit (level=error; fields: correlation id, route, status, err). Every 5xx in the HTTP mapper routes through it.

**Implementation**
- Implement `LogError`; wire into the error mapper for all 5xx mappings.

**Testing**
- Unit: `LogError` emits expected fields (capture zerolog output).
- Integration: forced handler panic/5xx produces one structured error log.

---

## Task 0.5 — CI hardening (`-race` + `govulncheck`)

**Type:** infra · **Effort:** S · **TRD:** §6.5
**Progress signal:** CI fails red on an injected data race or a known vuln.

**Design**
- CI runs `go test` without `-race` (S-level finding) → races invisible. Add `-race` + `govulncheck`, block on High/Critical.

**Implementation**
- Edit `.github/workflows/go.yml`: `go test -race ./...`; add `govulncheck ./...` step.

**Testing**
- Sanity: temporary race in a branch makes CI red; revert.
- `govulncheck` clean on main.

---

## Sprint 0 — Definition of Done

- [ ] Zero known P0 defects; rotated creds for anything shipped in a published image.
- [ ] Image contains no secrets (verified command fails).
- [ ] Hash-leak endpoint removed/gated + password field dropped.
- [ ] `LogError` implemented; 5xx emit structured logs.
- [ ] `go test -race ./...` + `govulncheck` green in CI.
- [ ] Sessions survive rolling restart.

## Swagger

No new public API. **Action:** regenerate `swag init -g main.go` after removing the test endpoint so docs no longer advertise it; enforce swagger-up-to-date in CI.

## Risks

| Risk | L | I | Mitigation |
|---|---|---|---|
| Creds already leaked via published image | M | Critical | Rotate all C1/C2 creds; audit pull logs |
| SOPS slows deploy | M | Med | Single encrypted env file first |

## Dependencies

Blocks every later sprint. No upstream dependency. Decide SOPS key backend (age for POC).
