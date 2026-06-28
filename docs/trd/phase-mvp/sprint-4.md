# Sprint 4 — Member Surface + Email

> **Phase:** MVP · **Duration:** 2 weeks · **TRD anchor:** PRD §5.A, §3.7 (member API + SDK), FR-1.4 (notification)
> **Depends on:** Sprint 3 complete.

## Goal

Member can join, see balance + ledger, and redeem — through a thin embeddable surface. Transactional email reliable (OTP + earn/redeem confirmation). First sprint with **frontend**.

## Why now

Backend correct, isolated, real-time (S0–S3). Now expose the loop. UI is thin: wireframes (started in parallel during S0–S1) lock the API contract built here.

**Out of scope:** no-code rule builder, tiers, dashboards, native mobile (all post-MVP).

**Blocker before start:** email provider selection + starting-balance policy (Open Questions) must be closed. Redeem-tap flow depends on POS/checkout integration model — settle first.

---

## Task 4.1 — Enrollment + optional welcome balance

**Type:** backend · **Effort:** M · **TRD:** PRD §5.A Enrollment
**Progress signal:** new member created under correct tenant; optional starting balance applied; confirmation email sent.

**Design**
- Member created under `tenant_id` (from S3.5 auth context). Optional welcome balance per program config → seeds an earn ledger entry (rule-engine or config-driven). Confirmation email queued.

**Implementation**
- `POST /v1/member/enroll`: create member, apply starting balance as a ledger entry (idempotent), trigger confirmation email.

**Testing**
- Enroll → member under correct tenant; starting balance = config; email enqueued.
- Edge: enroll without welcome-balance config → 0 balance, no spurious entry.

**Swagger:** `POST /v1/member/enroll` request/response + error codes registered.

---

## Task 4.2 — Transactional email (OTP + confirmations)

**Type:** backend · **Effort:** M · **TRD:** PRD §5.A, §3.6 (Temporal/n8n delivery), PRD §6 (OTP not delivered)
**Progress signal:** OTP email arrives; "resend code" works; earn/redeem send confirmation.

**Design**
- Reliable delivery via selected provider; route through Temporal activity (durable retry) or n8n integration. OTP for enrollment/login; confirmations on earn/redeem. "Resend code" so enrollment never silently stuck.

**Implementation**
- Email service abstraction + provider adapter; OTP issue/verify; confirmation hooks on `points.committed`.

**Testing**
- OTP delivered + verifies; resend issues a fresh valid code; undelivered → resend path, not stuck.
- Confirmation fires on earn and redeem.

**Swagger:** `POST /v1/member/auth/otp/request`, `/otp/verify`, `/otp/resend` registered.

---

## Task 4.3 — Balance + ledger display (member widget)

**Type:** frontend + backend · **Effort:** L · **TRD:** PRD §5.A Point Visibility, §3.7 (embeddable SDK)
**Progress signal:** widget renders balance + history + cash value; reflects a completed transaction <2s.

**Design**
- Embeddable web widget (B2B2C surface). Shows balance, cash equivalent ("500 pts = Rp 5,000"), ledger history. Reads tenant-scoped member API. Real-time <2s via the S3 event path (poll or push).
- Wireframe (low-fi, locked in S0–S1) drives the API response shape — fields the widget needs define the contract.

**Implementation**
- Backend: `GET /v1/member/balance`, `GET /v1/member/ledger` (paginated), include cash-value conversion.
- Frontend: widget consuming those endpoints; loading/empty/error states.

**Testing**
- Backend: endpoints tenant-scoped, correct cash conversion, pagination.
- Frontend: renders balance/history; updates within 2s after a seeded transaction; empty-state.

**Swagger:** `GET /v1/member/balance`, `GET /v1/member/ledger` registered with schemas.

---

## Task 4.4 — Redeem at checkout

**Type:** frontend + backend · **Effort:** M · **TRD:** PRD §5.A Redemption, FR-1.3
**Progress signal:** ≤3 taps from trigger to confirmation; balance validated; atomic deduction (reuses S2 guards).

**Design**
- Redeem flow at checkout reuses the atomic, double-spend-safe redemption from Sprint 2. ≤3 taps trigger→confirm. Validated balance; failed redeem deducts nothing; confirmation email/screen.

**Implementation**
- Backend: `POST /v1/member/redeem` (delegates to `RedemptionService.Create`).
- Frontend: redeem UI in the widget — select/confirm, success + error ("Insufficient Balance") states.

**Testing**
- Backend: redeem success/insufficient/duplicate (idempotent) — already covered by S2 service tests; add member-endpoint integration.
- Frontend: tap-count ≤3; success + insufficient-balance UX.

**Swagger:** `POST /v1/member/redeem` registered.

---

## Task 4.5 — Operator ledger view

**Type:** frontend + backend · **Effort:** M · **TRD:** PRD §5.B Ledger Visibility
**Progress signal:** operator sees a member's tenant-scoped earn/burn history with delta/source/balance/timestamp.

**Design**
- Operator (admin surface) views per-member ledger for support/verification. Tenant-scoped. Read-only for MVP.

**Implementation**
- Backend: `GET /v1/admin/members/{id}/ledger` (tenant-scoped via S3.2).
- Frontend: simple operator ledger table.

**Testing**
- Backend: tenant-scoped; cannot view other tenant's member (isolation test).
- Frontend: renders entries; empty-state.

**Swagger:** `GET /v1/admin/members/{id}/ledger` registered.

---

## Sprint 4 — Definition of Done

- [ ] Full loop demoable: join → earn → see balance → redeem.
- [ ] Member widget renders balance/ledger/cash-value; reflects txn <2s.
- [ ] Redeem ≤3 taps, atomic, double-spend-safe.
- [ ] OTP + confirmation emails reliable; resend works.
- [ ] Operator ledger view live, tenant-scoped.
- [ ] All new endpoints in Swagger; `go test -race` green; isolation tests pass.

## Swagger

New endpoints registered: enroll, OTP (request/verify/resend), balance, ledger, redeem, admin member-ledger. `swag init -g main.go`; CI enforces up-to-date.

## Risks

| Risk | L | I | Mitigation |
|---|---|---|---|
| Email provider undecided blocks OTP | M | High | Close Open Question before sprint start |
| Real-time <2s not met via polling | M | Med | Use S3 event path; tune poll/push; load-test |
| Redeem UX coupling to POS undefined | M | Med | Settle checkout integration model before 4.4 |

## Dependencies

S2 (redeem guards), S3 (tenant scoping + real-time events + member auth). Email provider + starting-balance policy resolved. UX wireframes from S0–S1.
