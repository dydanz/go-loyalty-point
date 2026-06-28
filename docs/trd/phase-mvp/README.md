# TRD — Phase MVP: Sprint Tech Dev Plans

Per-sprint technical development plans for the **MVP (Core Earn & Burn Loop)**.

**Sources:**
- PRD: [`docs/prd/prd-mvp-core-earn-burn.md`](../../prd/prd-mvp-core-earn-burn.md)
- TRD: [`docs/trd-loyalty-platform-v2.md`](../../trd-loyalty-platform-v2.md) — Phase 0 (§5) + Phase 1 (§5) + architecture §3
- Code analysis: [`docs/code-analysis-20260623.md`](../../code-analysis-20260623.md)

**Cadence:** 2-week sprints. Ordering by **dependency, not value** — security + correctness foundations precede the member-facing surface. Tenant scoping threads through from Sprint 1.

---

## Sprint Index

| Sprint | Theme | TRD anchor | Gate |
|---|---|---|---|
| [Sprint 0](./sprint-0.md) | Stabilize & Secure | TRD Phase 0 (C1–C4, S8) | **Hard gate — blocks all release** |
| [Sprint 1](./sprint-1.md) | Rule Engine + Money-Path Tests | FR-1.2, S1 | Depends on S0 |
| [Sprint 2](./sprint-2.md) | Ledger Integrity + Idempotency | FR-1.1, FR-1.3, §3.5 | Depends on S1 |
| [Sprint 3](./sprint-3.md) | Tenant Scoping + Real-Time Events | FR-1.4, §3.7, §3.5 | Depends on S2 |
| [Sprint 4](./sprint-4.md) | Member Surface + Email | PRD §5.A | Depends on S3 |
| [Sprint 5](./sprint-5.md) | Observability + Launch Readiness | §6.2, NFR | Depends on S4 |

## Definition of Done (every story)

- Code merged + `go test -race ./...` green in CI
- Tenant isolation respected (no cross-tenant read/write)
- Observable: error paths logged via `LogError`; key paths metered
- Parameterized SQL only; migrations append-only
- No secrets in code/image

## Sequencing Notes

- **Sprint 0 is a hard gate.** P0 security defects (§2.3 C1–C4) make the system unsafe to build on. Nothing member-facing ships before it.
- **Sprints 1–3 are backend-only** — no UI dependency; runnable by a small backend team.
- **UX wireframes** (member balance/redeem widget, operator ledger view) start in parallel during S0–S1 and lock the Sprint 4 API contract.
- **Open Questions** must close before the sprint that needs them: email provider + starting-balance policy (S4); points-liability accounting stance (S5); Temporal persistence choice — PostgreSQL vs Cassandra (S3, TRD §9.2).
- Compressible to ~4 sprints if S2 + S3 parallelize across two engineers.

## Scope boundary (MVP non-goals — do NOT build here)

No-code rule builder UI · tiers · voucher campaigns · points expiry · gamification · partner/coalition · AI · OIDC federation (local auth only) · real-time analytics dashboards · native mobile SDK. See PRD §4.
