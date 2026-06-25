# PRD: go-loyalty-point — MVP (Core Earn & Burn Loop, Stabilized)

**Status:** Draft 🟡 | **Target Release:** Q3 2026 | **Product Lead:** Dandi
**Phase:** MVP | **Source:** `product-analysis-v2.md` §5 MVP · `trd-loyalty-platform-v2.md` Phases 0–1

---

## 1. Executive Summary & Problem Statement

**The Problem:** The platform has a loyalty backend but the core loop does not actually work as a product. Earned points are computed with **hardcoded constants** — the configurable rules engine is dead code, so "set an earn rule" does nothing. Members have **no UI** to see a balance or redeem. The money paths (earn/redeem) have **zero tests**, and the build ships **P0 security defects** (DB credentials baked into the published Docker image, a public endpoint leaking bcrypt hashes). It cannot go live in any real environment.

**The Solution:** Ship the smallest *working, safe* loop: a member joins, earns points **computed by the real rule engine**, sees a balance, and redeems — all idempotent, real-time, tenant-scoped, and observable. Fix the P0s and add tests on the money paths so the loop is trustworthy.

**Why Now?** Every later phase (campaigns, tiers, gamification, AI) builds on a correct earn/burn ledger. Stabilizing and wiring the engine now is the prerequisite for all of it; shipping features on top of an untested, insecure, rules-decorative base would compound risk.

---

## 2. Target Audience

1. **The Member (B2C, via our B2B client):** A shopper using the client's branded app/web. Wants to see balance + cash value, earn when they transact, and redeem with minimal friction.
2. **The Operator (B2B client):** Business owner/marketer. For MVP, needs to see earn/redeem happen per member (ledger view); rule configuration is via API/us, not yet a no-code UI.
3. **Platform Operator (us):** Needs the loop to be correct, idempotent, observable, and tenant-isolated.

---

## 3. Success Metrics & KPIs
*MVP is successful if, within 90 days of launch:*

- **Activation:** >60% of enrolled members earn at least once within 7 days.
- **Correctness:** 100% of earn is computed by the rule engine (0% hardcoded); **0 balance-integrity incidents**.
- **Real-time:** Earn-to-visible latency <2s (p99); point-award API <300ms (p99).
- **Safety:** 0 known P0 security defects in the released build; earn/redeem covered by automated tests (`go test -race`).
- **Isolation:** 0 cross-tenant data-leak incidents.

---

## 4. Scope & Non-Goals

**✅ In Scope (MVP):**
- **Stabilization (TRD Phase 0):** SOPS secrets (remove `.env` from image), remove/gate the hash-leak endpoint, implement `LogError`, stop the startup session-wipe, add `-race` to CI, tests on `RedemptionService.Create` + `TransactionService.Create`.
- **Rule-driven earning:** wire `point_rewards_engine` into `TransactionService.Create`; delete shadow types.
- **Member loop:** enrollment, balance + ledger display (new member UI), basic redemption at checkout.
- **Idempotency:** transaction-ID dedupe (Redis key + unique DB constraint).
- **Real-time event:** `points.committed` to Kafka on earn/burn; `event_log` fallback when bus down.
- **Tenant scoping:** `tenant_id` on every record/query/event; member auth = per-tenant **local** (`auth_source = local`).
- **Reliable transactional email** (OTP + earn/redeem confirmation).
- **Observability baseline:** structured 5xx logging, metrics, runtime-configurable log level.

**❌ Out of Scope (Non-Goals for MVP):**
- No-code rule builder UI (rules wired but configured by API/us).
- Tiers · Voucher campaigns · Points expiry · Gamification · Partner/coalition · AI/personalization.
- OIDC member federation (local auth only).
- Real-time analytics dashboards (fake demo dashboard is **removed**, not replaced yet).
- Native mobile SDK.

---

## 5. Core Capabilities & User Stories

### A. Member Experience (member API + embeddable UI)

| Epic | User Story | Acceptance Criteria |
|---|---|---|
| **Enrollment** | As a member, I want to join a program and (optionally) get a welcome balance, so I start with value. | 1. Member created under correct `tenant_id`. 2. Optional starting balance applied per program config. 3. Confirmation email delivered. |
| **Point Visibility** | As a member, I want to see my balance and its cash value, so I know my purchasing power. | 1. Balance + ledger history render in member UI. 2. UI shows cash equivalent (e.g. "500 pts = Rp 5,000"). 3. Reflects a completed transaction <2s. |
| **Rule-driven Earning** | As a member, I want correct points when I transact, so the program is trustworthy. | 1. Points computed by the **rule engine**, not constants. 2. Same txn + different active rule → different points (regression test). 3. Ledger entry immutable with txn ID, source, resulting balance. |
| **Redemption** | As a member, I want to redeem points for a discount/reward at checkout, so I get instant value. | 1. Balance validated before commit. 2. Deduction atomic; failed redemption deducts nothing. 3. ≤3 taps from trigger to confirmation. |

### B. Operator / Platform Experience

| Epic | User Story | Acceptance Criteria |
|---|---|---|
| **Ledger Visibility** | As an operator, I want to view a member's earn/burn history, so I can support and verify. | 1. Per-member ledger, tenant-scoped. 2. Each entry shows delta, source, balance, timestamp. |
| **Idempotent Ingestion** | As the platform, I must never double-award on a retry. | 1. Re-submitting the same txn ID is a no-op (Redis + unique DB constraint). 2. Verified under concurrent submissions. |
| **Safe & Observable** | As the platform, I must run with no P0 defects and full error visibility. | 1. Published image has no secrets (`docker run … cat /app/.env` fails). 2. 5xx logged via `LogError`. 3. `go test -race ./...` green in CI. |

---

## 6. Key Edge Cases & Error States

- **Double-spend (two tabs/devices redeeming one balance).** → Ledger enforces DB locking + idempotency keys; second concurrent redeem fails with "Insufficient Balance," no negative drift.
- **Retry / network duplicate on earn.** → Same transaction ID → idempotent no-op; balance credited exactly once.
- **Return / refund after earn.** → Refund reverses the earn as a compensating ledger entry; balance may go negative (member earns out of "debt" before next redeem). Reversal is atomic and audited.
- **Kafka unavailable at earn time.** → Sync ledger commit still succeeds (Kafka is fan-out, not a hard dependency); event written to PostgreSQL `event_log` and replayed.
- **OTP email not delivered.** → "Resend code" available; enrollment not silently stuck.
- **Cross-tenant access attempt.** → `TenantScope` middleware rejects; integration test asserts tenant A cannot read tenant B.

---

## 7. Technical Dependencies

- **Rules engine wiring** — `ProgramRuleRepository.GetActiveRules` into `TransactionService.Create` (TRD Phase 1, S1).
- **Event bus (Kafka)** activated with `event_log` fallback (currently imported, unused).
- **Redis** for idempotency keys (in addition to sessions/cache).
- **SOPS** secret management + credential rotation; CI `-race` + `govulncheck`.
- **Email/push provider** selection (Open Question — block before launch).
- **POS / checkout** earn-trigger via REST API.
- **Member SDK / embeddable web widget** (B2B2C surface) for balance/redeem.

---

**Next Steps for the PM Team:**
1. Confirm the Non-Goals boundary with engineering (esp. "rules wired but no builder UI yet").
2. Resolve MVP Open Questions before build: email provider, starting-balance policy, points-liability accounting stance.
3. Hand Core Capabilities to UX for low-fi wireframes of the member balance/redeem widget + operator ledger view.
