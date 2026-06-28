# Sprint 5 — Observability + Launch Readiness

> **Phase:** MVP · **Duration:** 2 weeks · **TRD anchor:** §6.2 (observability), NFR (§Phase 1), PRD §3 KPIs, §4.7 (deprecate fake dashboard)
> **Depends on:** Sprint 4 complete.

## Goal

Production-ready: metrics + runtime log control, latency within budget, fake demo dashboard gone, final P0 sweep, launch sign-off. Make every MVP KPI measurable.

## Why now

Loop works end-to-end (S4). Now prove it's observable, fast, and clean before go-live.

**Out of scope:** ClickHouse/real analytics dashboards (Phase 1), OTel full tracing (can start here, completes Phase 1).

---

## Task 5.1 — Metrics + runtime-configurable log level

**Type:** backend · **Effort:** M · **TRD:** §6.2
**Progress signal:** `/metrics` exposes request count/latency/error rate; log level changes without redeploy.

**Design**
- Prometheus metrics (request count, latency histogram, error rate) on key paths. zerolog level runtime-configurable (no redeploy) — closes the §6.2 + PRD observability-baseline gap.

**Implementation**
- Metrics middleware + `/metrics` endpoint. Runtime log-level setter (admin endpoint or signal/config reload).

**Testing**
- `/metrics` returns expected series after traffic.
- Flip log level at runtime → verbosity changes, no restart.

**Swagger:** `/metrics` documented (or noted as Prometheus scrape, non-`/v1`). Log-level admin endpoint registered if HTTP.

---

## Task 5.2 — Latency hardening

**Type:** backend · **Effort:** M · **TRD:** NFR §Phase 1, PRD §3 (Real-time)
**Progress signal:** earn-to-visible <2s p99; point-award API <300ms p99 under load.

**Design**
- Validate NFR targets: API p99 <300ms, earn-to-visible <2s. Profile hot paths (rule eval, ledger insert, balance read). Load test toward MVP 500 TPS (headroom to 1,000).

**Implementation**
- Profile + tune queries/indexes/cache; fix N+1 or slow rule eval found.

**Testing**
- Load test (k6/vegeta): assert p99 budgets at target TPS; capture report.

---

## Task 5.3 — Remove fake demo dashboard

**Type:** frontend · **Effort:** S · **TRD:** §4.7 (deprecate), PRD §4
**Progress signal:** no screen shows template/demo "Sales" data; honest empty state instead.

**Design**
- Fake "Sales" template dashboard removed, not replaced (real analytics is Phase 1). Honest empty state or redirect to the member/operator surfaces from S4.

**Implementation**
- Delete demo dashboard + sample data; replace landing with real surface or empty state.

**Testing**
- No template data renders anywhere (UI smoke + grep for sample fixtures).

---

## Task 5.4 — Close MVP Open Questions (policy)

**Type:** product/backend · **Effort:** S · **TRD:** PRD §10, §6.3 (double-entry audit)
**Progress signal:** points-liability accounting stance + starting-balance policy documented + implemented.

**Design**
- Resolve: points-liability accounting stance (audit-trail expectation per §6.3 — double-entry built now), starting-balance policy. Email provider already settled in S4.

**Implementation**
- Encode policies in config; ensure ledger audit trail satisfies the agreed accounting stance.

**Testing**
- Config drives starting balance; audit trail export reconciles earn/burn.

---

## Task 5.5 — Launch-readiness checklist

**Type:** infra/backend · **Effort:** M · **TRD:** PRD §3 (all KPIs), §6.2 (SLO before Done)
**Progress signal:** signed checklist — 0 known P0, `-race` green, isolation tests pass, runbook exists, SLO dashboard live.

**Design**
- Gate go-live on: 0 known P0 (re-sweep S0 items), `go test -race ./...` green, tenant isolation tests pass, KPIs instrumented (activation, correctness, latency, safety, isolation), liveness/readiness probes (`/healthz/live`, `/healthz/ready`), runbook + rollback plan, SLO dashboard live (§6.2 — no phase Done without it).

**Implementation**
- Split health probes; assemble dashboard from 5.1 metrics; write runbook/ADR (bus-factor mitigation, §8).

**Testing**
- Full `go test -race ./...` green; isolation suite green; probes respond; dashboard shows live KPIs.

**Swagger:** `/healthz/live`, `/healthz/ready` documented.

---

## Sprint 5 — Definition of Done

- [ ] `/metrics` live; log level runtime-configurable.
- [ ] p99 budgets met under load (earn-visible <2s, API <300ms) at target TPS.
- [ ] Fake dashboard removed; no demo data anywhere.
- [ ] Open Questions (accounting, starting-balance) closed + encoded.
- [ ] Launch checklist signed: 0 P0, `-race` green, isolation green, probes + runbook + SLO dashboard.
- [ ] All MVP §3 KPIs instrumented and visible.

## Swagger

`/metrics`, `/healthz/live`, `/healthz/ready`, log-level admin endpoint (if HTTP) registered. Final `swag init -g main.go`; CI enforces up-to-date as a launch gate.

## Risks

| Risk | L | I | Mitigation |
|---|---|---|---|
| Latency target missed at 500 TPS | M | High | Profile early; index/cache tune; 2× headroom load test |
| Single-developer bus factor | H | High | Runbooks + ADRs (TRD §8); small reviewed changes |

## Dependencies

S1–S4 complete. Closes the MVP. Real analytics dashboard + ClickHouse + OTel-complete + Temporal expansion → Phase 1.
