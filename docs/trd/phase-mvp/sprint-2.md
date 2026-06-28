# Sprint 2 — Ledger Integrity + Idempotency

> **Phase:** MVP · **Duration:** 2 weeks · **TRD anchor:** FR-1.1, FR-1.3, §3.3 principles 1–2, §3.5
> **Depends on:** Sprint 1 complete.

## Goal

Ledger never double-awards on retry, never drifts negative under concurrency, and reverses refunds cleanly. **The single most critical risk in the whole program is ledger data integrity (TRD §1)** — this sprint owns it.

## Why now

Engine wired (S1) → now make the ledger trustworthy under real-world retries and concurrent access before exposing it to members (S4).

**Out of scope:** events/Kafka (Sprint 3), Kafka write-buffer path (later phase).

---

## Task 2.1 — Idempotency key on earn

**Type:** backend · **Effort:** M · **TRD:** FR-1.1, §3.3 principle 2, §3.5
**Progress signal:** resubmitting the same transaction ID twice yields one ledger entry, one balance change.

**Design**
- Two-layer dedupe: Redis key (fast) + **unique DB constraint** on transaction ID (authoritative). Extend the existing atomic CTE insert. Client supplies transaction ID. Same ID = no-op returning the original result.

**Implementation**
- Migration: unique constraint on `(tenant_id, transaction_id)` in the ledger table (append-only migration).
- Redis idempotency check before insert; on unique-violation, fetch + return original entry (idempotent response, not error).

**Testing**
- Sequential resubmit → identical response, single entry.
- Redis-miss + DB-hit path (simulate Redis eviction) → still deduped by DB constraint.

**Swagger:** document `transaction_id` as required idempotency key on `POST /v1/transactions`; note resubmit semantics (same result, no double credit).

---

## Task 2.2 — Idempotency under concurrency

**Type:** backend · **Effort:** M · **TRD:** FR-1.2 (concurrent), §6.5 race
**Progress signal:** N concurrent identical submits → exactly one credit (race test green).

**Design**
- Concurrent same-txn submits must not both insert. Rely on DB unique constraint as the serialization point; CTE insert stays atomic.

**Implementation**
- Ensure insert path handles unique-violation as success-dedupe under contention.

**Testing**
- `go test -race`: spawn goroutines submitting same txn ID; assert one entry, correct balance.

---

## Task 2.3 — Atomic redemption + double-spend guard

**Type:** backend · **Effort:** M · **TRD:** FR-1.3, PRD §6 (double-spend)
**Progress signal:** two concurrent redeems on one balance → one succeeds, one fails "Insufficient Balance", no negative drift.

**Design**
- Balance validated before commit; deduction atomic via DB locking + idempotency. Second concurrent redeem fails cleanly.

**Implementation**
- Row-lock / atomic CTE on redeem; reject when insufficient; failed redeem deducts nothing.

**Testing**
- Race test: two devices/tabs redeem same balance → exactly one success.
- Assert balance never goes negative on the redeem path.

---

## Task 2.4 — Refund / reversal as compensating entry

**Type:** backend · **Effort:** M · **TRD:** FR-1.1, PRD §6 (refund)
**Progress signal:** refund after earn produces a compensating ledger entry; balance may go negative; fully audited.

**Design**
- Refund reverses the original earn as a new compensating entry (not a mutation — ledger is append-only). Atomic + audited. Balance may go negative until re-earned.

**Implementation**
- Reversal entry references original txn ID; source = refund. Idempotent on refund ID.

**Testing**
- Earn → refund → balance reflects reversal; original entry untouched.
- Double refund of same refund ID → single reversal (idempotent).

**Swagger:** `POST /v1/transactions/{id}/refund` (or reversal event) documented.

---

## Task 2.5 — Immutable ledger entry shape

**Type:** backend · **Effort:** S · **TRD:** FR-1.1, glossary PointsLedger
**Progress signal:** every entry carries delta, source, txn ID, resulting balance, timestamp; none mutable.

**Design**
- Append-only; balance = last entry per customer+program. Entry immutable after write.

**Implementation**
- Confirm/lock schema fields; no UPDATE paths on ledger rows.

**Testing**
- Contract test: entry has all required fields; attempt to mutate is not exposed.

---

## Sprint 2 — Definition of Done

- [ ] Resubmit same txn ID = no-op (Redis + unique DB constraint), verified under concurrency.
- [ ] Concurrent redeem cannot double-spend or drift negative.
- [ ] Refund reverses atomically as a compensating, audited entry.
- [ ] Ledger entries immutable with full required fields.
- [ ] PRD §6 edge cases (double-spend, retry-dup, refund) covered by `-race` tests.

## Swagger

`POST /v1/transactions` (idempotency key documented), `POST /v1/redemptions`, refund/reversal endpoint — all regenerated via `swag init`. Error responses (insufficient balance, duplicate) documented with codes.

## Risks

| Risk | L | I | Mitigation |
|---|---|---|---|
| Balance corruption during constraint rollout | M | Critical | Reconciliation test on staging; rollback plan; dry-run migration |
| Redis/DB dedupe divergence | L | High | DB constraint is authoritative; Redis is optimization only |

## Dependencies

S1 (engine produces the amount being credited). No frontend. Migration must be append-only (CLAUDE.md hard rule).
