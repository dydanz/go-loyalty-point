# Sprint 3 — Tenant Scoping + Real-Time Events

> **Phase:** MVP · **Duration:** 2 weeks · **TRD anchor:** FR-1.4, §3.7, §3.5, §3.3 principle 5
> **Depends on:** Sprint 2 complete.

## Goal

Every record, query, cache key, and event is **tenant-scoped** — a B2B2C surface cannot leak one client's members to another. Earn/burn fans out in real time via Kafka, with a PostgreSQL `event_log` fallback so the bus is never a hard dependency.

## Why now

Ledger is correct (S2). Before exposing it to members (S4), enforce isolation and stand up the real-time event path the member dashboard's "<2s" requirement needs.

**Out of scope:** Kafka write-buffer / 202 ingestion (later phase — MVP uses sync-authority path only), notifications UI.

---

## Task 3.1 — `tenant_id` on every record + query

**Type:** backend · **Effort:** L · **TRD:** §3.7
**Progress signal:** all ledger/member/transaction reads + writes filtered by `tenant_id`; no unscoped query remains.

**Design**
- Tenant tree: business/client → Merchant → Program → Customer. Add `tenant_id` propagated through auth claims into every repository call and cache key. Single source: token claim, not request param.

**Implementation**
- Migration: ensure `tenant_id` column + index on member/transaction/ledger tables (append-only).
- Thread `tenant_id` through service → repository signatures; scope cache keys (`{tenant}:...`).

**Testing**
- Repository tests assert WHERE `tenant_id` present on every query.
- Seed two tenants; assert queries return only own-tenant rows.

---

## Task 3.2 — `TenantScope` middleware + isolation tests

**Type:** backend · **Effort:** M · **TRD:** §3.7, §6.5 (tenant isolation mandatory)
**Progress signal:** integration test: member of tenant A gets 403/empty on tenant B data.

**Design**
- Middleware extracts `tenant_id` from the validated token, injects into request context; services read it from context, never from client input.

**Implementation**
- Add `TenantScope` to the middleware chain (after Auth). Reject/zero cross-tenant access.

**Testing**
- **Mandatory isolation test (§6.5):** tenant A token cannot read tenant B member/ledger — asserted in integration suite. This test is a release gate.

**Swagger:** document that all `/v1/*` endpoints are tenant-scoped via bearer token; no `tenant_id` request param accepted.

---

## Task 3.3 — Emit `points.committed` to Kafka

**Type:** backend · **Effort:** L · **TRD:** FR-1.4, §3.5 (sync-authority path)
**Progress signal:** earn/burn produces a Kafka event within 2s carrying tenant id, delta, new balance, txn ref.

**Design**
- Sync-authority model: commit ledger to PostgreSQL **synchronously**, then publish `points.committed` for fan-out. Payload: tenant id, points delta, new balance, transaction reference, trigger event type. Event key = customer id (ordering).

**Implementation**
- Wire kafka-go producer (currently imported, unused) in `pkg/channel`/`pkg/kafka`; publish after successful ledger commit.

**Testing**
- Integration with embedded/test Kafka: earn → event observed <2s with correct payload + tenant scoping.

---

## Task 3.4 — `event_log` fallback when bus down

**Type:** backend · **Effort:** M · **TRD:** FR-1.4, §3.5 (failure handling), PRD §6
**Progress signal:** with Kafka stopped, earn still commits; event lands in PostgreSQL `event_log`, replays when bus returns.

**Design**
- Kafka is fan-out, **not** a hard dependency on the sync path. On publish failure, write the event to PostgreSQL `event_log`; a replayer drains it when Kafka recovers.

**Implementation**
- `event_log` table (append-only migration); fallback write on producer error; replay worker.

**Testing**
- Kill Kafka in test → earn succeeds, row in `event_log`; restart Kafka → row replayed, marked done.

---

## Task 3.5 — Per-tenant member auth (`auth_source = local`)

**Type:** backend · **Effort:** M · **TRD:** §3.10 (local only for MVP)
**Progress signal:** same email registers independently under two tenants; tokens are tenant-scoped.

**Design**
- MVP = **local auth only** (federation deferred). Member credentials scoped to `tenant_id`; same email can exist per tenant. Issue tenant-scoped member token consumed by `/v1/member/*`.

**Implementation**
- Member auth keyed by `(tenant_id, email)`; token carries `tenant_id` + member id.

**Testing**
- Register email under tenant A and B → two distinct members; cross-tenant token rejected.

**Swagger:** `POST /v1/member/auth/login`, `/register` documented (tenant resolved from context/subdomain, not body).

---

## Sprint 3 — Definition of Done

- [ ] `tenant_id` on every record/query/cache key/event.
- [ ] `TenantScope` middleware live; mandatory isolation test green (release gate).
- [ ] `points.committed` emitted <2s on earn/burn with correct payload.
- [ ] Kafka outage → sync commit still succeeds; `event_log` fallback + replay verified.
- [ ] Local per-tenant member auth working; 0 cross-tenant identity leaks.

## Swagger

All `/v1/*` marked tenant-scoped (bearer). Member auth endpoints registered. Event payloads documented in an async/events appendix (not REST, but record the contract). `swag init -g main.go`.

## Risks

| Risk | L | I | Mitigation |
|---|---|---|---|
| Missed unscoped query leaks cross-tenant | M | Critical | Isolation test as gate; repo-level review for WHERE tenant_id |
| Kafka activation destabilizes commit path | M | Med | Fan-out only; `event_log` fallback (§3.5) |

## Dependencies

S2 (ledger commit is the event trigger). Temporal **not** required this sprint (notification workflow deferred to Phase 1 proper / Sprint 4 email). Open question (TRD §9.2): Temporal persistence choice — not blocking MVP.
