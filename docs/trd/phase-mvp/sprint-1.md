# Sprint 1 — Rule Engine + Money-Path Tests

> **Phase:** MVP · **Duration:** 2 weeks · **TRD anchor:** FR-1.2, S1 (§2.3), §3.3 principle 3
> **Depends on:** Sprint 0 complete.

## Goal

Earn computed by the **real rule engine**, not hardcoded constants. Money paths (`TransactionService.Create`, `RedemptionService.Create`) get real test coverage so the loop is trustworthy.

## Why now

§2.3 S1: the rules engine (`point_rewards_engine.go`) is dead code; `transaction_service.go` uses hardcoded math — the core differentiator is a façade. §2.4: money paths have zero tests. Wire + cover before adding ledger guarantees.

**Out of scope:** idempotency/concurrency (Sprint 2), events (Sprint 3), member UI (Sprint 4).

---

## Task 1.1 — Wire rule engine into earn path

**Type:** backend · **Effort:** L · **TRD:** FR-1.2
**Progress signal:** same transaction under two different active rules yields two different point totals.

**Design**
- `TransactionService.Create` must consult active `ProgramRule`s via the rewards engine for every earn. Fetch rules through `ProgramRuleRepository.GetActiveRules(programID)`; pass domain `Transaction` + rules into `calculatePoints`/`evaluateRule`. Rule precedence + no-matching-rule fallback (0 or program default) defined explicitly.

**Implementation**
- Inject `ProgramRuleRepository` into `TransactionService` (constructor DI — no global state).
- Replace hardcoded constant computation with engine call.
- Handle: no active rule, multiple matching rules (precedence), rule references missing field.

**Testing**
- Unit (table-driven): amount/type/category inputs × rules → expected points.
- Regression: fixture transaction + rule A vs rule B → different totals (the FR-1.2 DoD test).
- Edge: zero amount; no active rule; overlapping rules.

**Swagger:** earn endpoint (`POST /v1/transactions`) response schema unchanged in shape, but document that `points_earned` is rule-derived; regenerate swag.

---

## Task 1.2 — Delete shadow types

**Type:** backend · **Effort:** S · **TRD:** FR-1.2, S1
**Progress signal:** engine compiles against domain types only; duplicate types gone, tests still green.

**Design**
- `point_rewards_engine.go` defines local shadow `Transaction`/`Program`/`ProgramRule` duplicating domain types — bug source. Engine must use `pkg/domain` types.

**Implementation**
- Remove shadow types; update engine signatures to domain types; fix call sites.

**Testing**
- Compilation + existing engine tests pass against domain types.
- No remaining references to shadow types (grep guard in review).

---

## Task 1.3 — Test coverage: `TransactionService.Create`

**Type:** backend · **Effort:** M · **TRD:** §2.4 S8, §6.5 (≥80%)
**Progress signal:** coverage report ≥80% on the service; CI gate green.

**Design**
- Use existing testify/suite template (`points_service_test.go` style) + mocks from `pkg/mocks`. Cover earn happy path, rule-driven amounts, error mapping (typed domain errors).

**Implementation**
- Tests with mocked repos; assert ledger insert called with engine-computed amount; assert typed errors propagate (not swallowed).

**Testing (self)**
- `go test -race -cover` ≥80% on `TransactionService`.

---

## Task 1.4 — Test coverage: `RedemptionService.Create`

**Type:** backend · **Effort:** M · **TRD:** FR-1.3, §2.4 S8
**Progress signal:** redeem coverage ≥80%; failed-redeem-deducts-nothing test passes.

**Design**
- FR-1.3: validate balance before commit; failed redemption deducts nothing; rollback atomic. (Full concurrency in Sprint 2; here cover logical correctness.)

**Implementation**
- Tests: sufficient balance → deduct; insufficient → reject with typed error, no deduction; mid-failure → rollback.

**Testing (self)**
- `go test -race -cover` ≥80% on `RedemptionService`.

---

## Sprint 1 — Definition of Done

- [ ] Earn 100% computed by rule engine; hardcoded constants removed.
- [ ] Regression proves rule change → point change.
- [ ] Shadow types deleted; engine uses domain types.
- [ ] Both money services ≥80% coverage; `go test -race` green.
- [ ] Swagger regenerated.

## Swagger

`POST /v1/transactions` (earn) and `POST /v1/redemptions` (redeem) documented with request/response schemas + typed error codes. `swag init -g main.go`; CI enforces up-to-date.

## Risks

| Risk | L | I | Mitigation |
|---|---|---|---|
| Rules wiring regresses ledger correctness | M | Critical | Regression + reconciliation tests before merge |
| Hidden hardcoded paths remain | M | High | Grep for constant math; assert via rule-change test |

## Dependencies

S0 (CI `-race`, `LogError`). Rule model assumed present in DB (TRD: already modeled). Frontend not involved.
