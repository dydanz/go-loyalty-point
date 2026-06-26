# Product Requirements Document — Loyalty & Rewards Platform
> Version: 2.0 (replaces product-analysis-20260623.md)
> Author: Claude (Opus 4.8) — Senior Product Manager
> Date: 2026-06-24
> Inputs: codebase analysis (`code-analysis-20260623.md`) · `loyalty-platform-research-prompt.md` (market research) · `loyalty-research-methodology.md` (Build/Buy/Compose framework) · cross-referenced with `trd-loyalty-platform-v2.md` for resolved technical decisions

---

## 1. Product Vision

go-loyalty-point is a **B2B2C loyalty engine for Southeast Asian businesses** — a platform our B2B clients use to run points-and-rewards programs *and* to power their own customers' earn/redeem experience. We optimize for one outcome: **a member earns and redeems points with zero friction, and an operator configures the rules behind it with zero engineering**. The product bet: most SMB and mid-market merchants in Indonesia don't want to build a loyalty ledger or hire engineers to change an earn rule — they want a reliable, embeddable engine where "double points on weekends" is a toggle, not a deploy.

---

## 2. Current State Audit

### 2.1 What Is Already Built
Derived from `code-analysis-20260623.md`. "Built" = server logic exists and is reachable; UI status noted separately.

| Feature | Status | Notes |
|---|---|---|
| **Account & Access** | | |
| Business sign-up | Built | Name/email/password/phone |
| Email OTP verification | Partial | Server issues code; email delivery unreliable (test endpoint exists to fetch code) |
| Login / logout | Built | bcrypt + random-token session (not JWT) |
| Account lockout | Built | 5 failed logins → 30 min lock (server-side) |
| **Merchant (Store) Management** | | |
| Store CRUD (add/edit/deactivate/list) | Built | Full admin UI exists |
| **Programs & Rules** | | |
| Program rule CRUD | Built (data only) | Tables + API exist; rules **never applied** to point math |
| View program rules | Partial | Read-only UI |
| **Points** | | |
| Transactional earning | Broken-by-design | Works, but uses **hardcoded constants**, not the rule engine |
| Points ledger | Built | Atomic CTE-based balance insert; sound. Server-only, no UI |
| Earn / redeem points | Built (untested) | Server logic exists; `RedemptionService.Create` has **zero tests** |
| **Rewards & Redemptions** | | |
| Reward CRUD | Built | Server-only, no UI |
| Redemption flow | Built (untested) | Server-only, no UI |
| **Customers** | | |
| Customer (member) management | Built | First-class entity, server-only, no UI |
| **Supporting** | | |
| Dashboard | Broken-as-product | Shows fake "Sales" template demo data |
| Billing / Profile pages | Placeholder | Template sample content |
| API docs (Swagger) | Built | `/swagger` |
| Health check | Built | `/ping` checks PG primary/replica/Redis |
| Event log | Built | Synchronous writes to PostgreSQL `event_log` (Kafka imported but unused) |

### 2.2 What Is Missing vs. Market Standard
Cross-referenced against `loyalty-platform-research-prompt.md` Document B. Gaps only.

| Gap | Market Standard | Impact if Missing |
|---|---|---|
| Rules actually drive point math | Runtime-configurable earn rules (Doc B1.2) | **High** — core differentiator is decorative; "set a rule" does nothing |
| Member-facing UI for points/rewards/redeem | Member dashboard + redemption flow (Doc D1) | **High** — core value unreachable by end users |
| Real-time earn feedback | Event-driven instant award (Doc F1) | **High** — Kafka unused; no instant "you earned X" |
| Points expiry | Configurable expiry + warnings (Doc B1.1) | **High** — no expiry logic; liability never ages off |
| Tiers & status | ≥3 configurable tiers (Doc B1.4) | Medium — no long-term engagement hook |
| Voucher engine | Bulk codes, stacking, eligibility (Doc B1.5) | Medium — no promo mechanism |
| Real analytics | Earn/burn ratio, lifecycle, cohort (Doc B1.9) | **High** — dashboard is fake demo data |
| Gamification | Challenges, streaks, badges (Doc B1.6) | Low — engagement, not core |
| Personalization / AI | Segments, churn, next-best-offer (Doc B1.7) | Low — later-stage |
| Reliable transactional email | OTP / lifecycle email (Doc D) | **High** — blocks enrollment today |

### 2.3 Technical Debt Affecting Product Decisions
Only debt that **constrains feature delivery** (evidence in `code-analysis-20260623.md`):

- **Rules engine is dead code** (`point_rewards_engine.go` never called; S1). → Configurable earning cannot ship until wired. **Blocks MVP.**
- **No event system active** (Kafka imported, unused). → No real-time awards, no async notifications, no analytics stream until activated (TRD Phase 1).
- **Money paths untested** (`RedemptionService.Create`, `TransactionService.Create` are 1-line stubs; S8). → Cannot trust earn/redeem correctness; **must add tests before MVP** is credible.
- **P0 security defects** (creds in image C1; bcrypt-hash leak C2; silent 5xx C3; session wipe on restart C4). → Cannot go to any real environment until fixed. **Blocks everything.**
- **No observability** (no metrics/tracing; 5xx silently swallowed). → Cannot operate or debug a live program.
- **Single developer.** → Sequencing must be ruthless; parallel tracks unrealistic.

---

## 3. Target Users

### User Type 1: Member (end customer of our B2B client)
- **Who they are:** A shopper who earns/redeems at our client's business. They interact with the client's branded app/web (B2B2C), powered by our engine.
- **Core jobs to be done:** See my balance; understand what it's worth; earn points when I transact; redeem for something I want with minimal friction.
- **Success metric:** Activation rate — % of enrolled members who earn at least once within 7 days (>60%).

### User Type 2: Operator (our B2B client's business owner / marketer)
- **Who they are:** The business owner or marketer who signs up, registers stores, defines programs and earning rules, and runs campaigns.
- **Core jobs to be done:** Set up a program; change earn rules without engineering; run a time-limited campaign; see whether it's working.
- **Success metric:** Time-to-campaign — launch a bonus-points campaign in <30 min via admin UI, no engineering.

### User Type 3: Platform Operator (us — internal ops/eng)
- **Who they are:** Our team operating the multi-tenant platform.
- **Core jobs to be done:** Keep earn/redeem correct and observable; detect fraud; onboard tenants (incl. optional SSO federation).
- **Success metric:** 99.9% ledger uptime; zero balance-integrity incidents.

---

## 4. Goals & Success Metrics

| Goal | Metric | Target | Phase |
|---|---|---|---|
| Members earn their first points | Activation rate | >60% within 7 days of enrollment | MVP |
| Earn/redeem is correct and safe | Balance-integrity incidents | 0 | MVP |
| Real-time feedback | Earn-to-visible latency (p99) | <2s | MVP |
| Operators launch a campaign without engineering | Time-to-campaign | <30 min via admin UI | Phase 1 |
| Redemption is frictionless | Redemption drop-off at checkout | <15% | Phase 1 |
| Configurable rules actually work | % earn computed by rule engine (not constants) | 100% | MVP |
| Members stay engaged via status | Tier-qualified member retention vs. base | +10pp | Phase 2 |
| Tenant isolation holds | Cross-tenant data-leak incidents | 0 | MVP |

---

## 5. Phased Roadmap

Sequence basis: **dependency order → member value → operator control**. Phases map to TRD: PRD MVP ≈ TRD Phase 0+1; PRD Phase 1 ≈ TRD Phase 3; PRD Phase 2 ≈ TRD Phase 2 tiers; PRD Phases 3–4 ≈ TRD Phase 4. (Stabilization/security from TRD Phase 0 is folded into MVP release criteria — non-negotiable.)

---

### MVP — Core Earn & Burn Loop (stabilized)
**Goal:** A member can join, earn points (computed by real rules), and redeem them. An operator can see it happen. The system is safe and observable.
**Timeframe:** ~6–8 weeks (includes ~2 weeks stabilization).
**Release criteria:** A real member completes earn → redeem in one session with no manual intervention; points are computed by the rule engine; zero known P0 security defects; earn/redeem covered by automated tests.

#### Features
| # | Feature | Description | User |
|---|---|---|---|
| 1 | Stabilization (P0 fixes) | Remove creds-from-image (SOPS), kill hash-leak endpoint, implement `LogError`, stop startup session-wipe | Platform |
| 2 | Member enrollment | Member joins a program under a tenant; optional starting balance | Member |
| 3 | Rule-driven transactional earning | Points awarded on a transaction, **computed by the wired rule engine** (fixes the dead-code gap) | Member |
| 4 | Points balance + ledger display | Member sees current balance + earn/burn history (new member UI) | Member |
| 5 | Basic redemption | Member redeems points for a discount / reward at checkout; atomic, idempotent | Member |
| 6 | Real-time earn event | `points.committed` emitted on earn/burn; balance visible <2s | Member |
| 7 | Idempotent earn/burn | Transaction-ID dedupe (Redis + DB constraint) — no double-award on retry | System |
| 8 | Operator ledger view | Operator views earn/burn history per member | Operator |
| 9 | Tenant scoping | Every record/query/event scoped by `tenant_id`; member auth = per-tenant local (`auth_source = local`) | System |
| 10 | Reliable transactional email | OTP + earn/redeem confirmation actually delivered | Both |

#### Out of Scope for MVP
- Tiers · Voucher campaigns · No-code rule builder UI (rules wired but configured by us/API in MVP) · Gamification · Partner earning · AI/personalization · OIDC federation (local auth only) · Real-time analytics dashboards · Native mobile SDK.

#### Risks
| Risk | Mitigation |
|---|---|
| Rules-engine wiring corrupts ledger balances | Idempotency + ledger-reconciliation tests before merge; dry-run on staging |
| P0 creds already leaked from published image | Rotate all creds; audit registry access |
| Adding tests surfaces latent money-path bugs | Treat as MVP scope, fix as found — do not defer |
| Email delivery still unreliable | Verify provider end-to-end + add "resend code" before launch |

---

### Phase 1 — Operator Control & Campaign Tools
**Goal:** Operators configure and launch earning campaigns and vouchers without writing code.
**Timeframe:** ~8–12 weeks after MVP.
**Release criteria:** An operator creates, publishes, and monitors a time-limited bonus-points campaign entirely via admin UI.

#### Features
| # | Feature | Description | User |
|---|---|---|---|
| 1 | No-code rule builder | If/then campaign logic (e.g., 2x points on weekends) over the existing rule model | Operator |
| 2 | Voucher generation (in-house) | Bulk creation of unique single-use codes (≥10k/batch) — **built, not Voucherify** | Operator |
| 3 | Voucher eligibility rules | Restrict by SKU, category, segment, channel, date range | Operator |
| 4 | Voucher stacking rules | Control which vouchers combine per transaction | Operator |
| 5 | Campaign scheduling | Start/end dates, timezone, auto-activation | Operator |
| 6 | Basic analytics dashboard | Active members, earn rate, redemption rate (served from ClickHouse) | Operator |
| 7 | Points expiry | Configurable expiry + 30/14/7-day member warnings | Both |
| 8 | Fraud velocity checks | Configurable max earn/redeem per member per window; admin alert + freeze/reverse | Operator |
| 9 | Customer management UI | Add/look up members (entity is server-only today) | Operator |

#### Dependencies
- MVP complete: rule engine wired, event bus active (Kafka), ledger idempotent, tenant scoping.
- ClickHouse analytics path stood up (TRD §3.8) for the dashboard.

#### Risks
| Risk | Mitigation |
|---|---|
| No-code builder produces unmaintainable/costly rule logic | Guardrails + rule version-control (author/timestamp/diff); validation on save |
| Bulk voucher generation performance | Target 10k codes <60s; load-test |
| Fraud thresholds mis-tuned → alert fatigue | Start conservative; make per-program configurable |

---

### Phase 2 — Tiers & Member Lifecycle
**Goal:** Members have a status level unlocking differentiated benefits and motivating long-term engagement.
**Timeframe:** ~10–14 weeks after Phase 1.
**Release criteria:** A member qualifies for a tier, sees new status + benefits, and gets a downgrade warning before losing tier.

#### Features
| # | Feature | Description | User |
|---|---|---|---|
| 1 | Tier definition | Configurable tiers with qualifying thresholds (spend/points/visits) | Operator |
| 2 | Tier qualification engine | Automated promote/demote on rolling or fixed window; runs on transaction + interval (Temporal) | System |
| 3 | Tier benefits | Multipliers, exclusive rewards, early access by tier | Member |
| 4 | Tier progress display | Visual progress bar to next tier | Member |
| 5 | Downgrade grace period | Configurable warning window before tier loss + notification | Both |
| 6 | OIDC member federation | `auth_source = federated` for SSO-demanding clients (hybrid auth fast-follow, TRD §3.10) | Member |
| 7 | Tier challenge `[PROPOSED]` | Accelerated path to next tier via a defined challenge | Member |

#### Risks
| Risk | Mitigation |
|---|---|
| Over-complex tier rules confuse members | One primary qualifying metric per tier (research D5) |
| Qualification window ambiguity | Resolve Open Question §9 before build |

---

### Phase 3 — Gamification & Behavioral Earning
**Goal:** Members earn for non-purchase actions — reviews, referrals, streaks — raising engagement frequency.
**Timeframe:** ~8–10 weeks after Phase 2.
**Release criteria:** A member completes a non-purchase action and sees points awarded in real time.

#### Features
| # | Feature | Description | User |
|---|---|---|---|
| 1 | Behavioral event earning | Points from custom API events (review, referral, app open) via event bus | Member |
| 2 | Referral program | Unique per-member codes; referrer + referee earn on activation | Member |
| 3 | Challenges & missions | Multi-step tasks with a defined completion reward | Member |
| 4 | Streak bonuses `[PROPOSED]` | Consecutive-activity rewards (daily/weekly) | Member |
| 5 | Badges / achievements `[PROPOSED]` | Non-monetary milestone recognition | Member |

#### Risks
| Risk | Mitigation |
|---|---|
| Behavioral earning becomes a fraud vector | Velocity checks (from Phase 1) extended to non-purchase events |
| Referral abuse | Cap referrals/member; require referee activation before award |

---

### Phase 4 — Personalization & AI
**Goal:** Offers/rewards dynamically matched to individuals by behavior and predicted intent.
**Timeframe:** ~12–16 weeks after Phase 3.
**Release criteria:** System surfaces a different offer to two members in the same tier based on their behavior profiles.

#### Features
| # | Feature | Description | User |
|---|---|---|---|
| 1 | Member segmentation | Rule-based cohorts (RFM, tier, category affinity) | Operator |
| 2 | Segment-based offers | Different earn/redemption rules per segment | Operator |
| 3 | Churn prediction trigger `[PROPOSED]` | At-risk members auto-receive a retention offer | System |
| 4 | Next-best-offer `[PROPOSED]` | AI recommends most relevant reward at the right moment | Member |
| 5 | A/B testing | Test reward values/mechanics on member subsets | Operator |

#### Prerequisites
- **Min data volume** (TRD): ≥6 months history + ≥10k active members + ≥100k transactions per tenant. Below this, rule-based segments only.
- ClickHouse analytics store live (from Phase 1).
- **Churn definition** agreed: "no qualifying activity in N days," N = 2–3× median inter-purchase interval per tenant.

#### Risks
| Risk | Mitigation |
|---|---|
| AI underperforms on thin data | Gate behind min-volume threshold; ship rule-based segments first |
| Churn model false positives waste retention spend | Bar: ≥0.70 precision @ ≥0.50 recall; must beat a recency heuristic |

---

### Phase 5 — Partner & Coalition (if applicable)
**Goal:** Members earn/redeem across more than one brand, increasing stickiness.
**Timeframe:** 12+ weeks after Phase 4, or a parallel track **only** if a partner is already committed.
**Release criteria:** A member earns at Partner Brand A and redeems at Brand B on one unified balance.

#### Features
| # | Feature | Description | User |
|---|---|---|---|
| 1 | Multi-brand earning `[PROPOSED]` | Points across partner network, pooled into one balance | Member |
| 2 | Partner admin portal `[PROPOSED]` | Separate operator view per partner | Partner Operator |
| 3 | Financial settlement `[PROPOSED]` | Automated earn-liability attribution between partners | Finance |
| 4 | White-label member portal `[PROPOSED]` | Per-brand branded UX | Member |

#### Gate Condition
Do not begin without a **signed partner commitment**. Do not architect for it in MVP. (Note: the tenant model in TRD §3.7 already keeps this a lift, not a rewrite.)

---

## 6. UX Requirements
From market research Document D. Apply across all phases unless noted.

### Member-Facing Principles
- Point balance visible at all times (persistent in nav/header).
- Always show points in equivalent cash value alongside the point count.
- Redemption ≤3 taps/clicks from trigger to confirmation.
- Expiry warnings at 30 / 14 / 7 days via in-app + push/email.
- Web fallback always available — never force app download (B2B2C: clients embed via API/SDK).

### Admin Console Principles
- Campaign creation completable without engineering.
- Every destructive action (delete campaign, bulk-expire points) requires confirmation + audit-log entry.
- Fraud alerts actionable from the alert itself.

### Anti-Patterns to Avoid (Document D5) — do not ship
- Hidden point monetary value.
- Redemption minimums requiring >2 earn sessions to reach.
- Silent fraud blocks with no member communication path.
- Forced app download to access loyalty features.
- Fake/demo data on real screens (current dashboard — remove in MVP).

---

## 7. Integration Requirements
From research Document C3 + codebase audit + TRD decisions.

| Integration | Type | Required By | Priority |
|---|---|---|---|
| POS / checkout | REST API (earn trigger) | MVP | P0 |
| Email / push provider | Webhook via Temporal/n8n (notifications, OTP) | MVP | P0 |
| Event bus (Kafka) | Internal event stream | MVP | P0 |
| Client app / SDK (B2B2C member surface) | REST API + JS/web widget | MVP→Phase 2 | P0 |
| Client IdP (OIDC) | Federated member auth | Phase 2 | P1 |
| CRM | Bidirectional sync (via n8n) | Phase 1 | P1 |
| ClickHouse (analytics) | Event-stream sink | Phase 1 | P1 |
| CDP / data warehouse export | Stream/batch from ClickHouse | Phase 4 | P2 |
| Partner POS | REST API | Phase 5 | P3 |

---

## 8. Non-Functional Requirements

| Requirement | Target | Notes |
|---|---|---|
| Point award latency | <300ms p99 (sync path) | Real-time member feedback |
| Earn-to-visible latency | <2s | Via event bus |
| Transaction throughput | 500 TPS MVP; load-tested to 1,000 TPS | Per TRD §9 |
| API availability | 99.9% uptime | SLA-aligned with checkout dependency |
| Idempotency | All earn/burn endpoints | Redis key + unique DB constraint |
| Tenant isolation | Enforced + tested | Member of tenant A cannot read tenant B |
| Data retention | Member ledger: indefinite (double-entry audit trail) | Analytics events: 90d hot / 13mo raw (ClickHouse TTL) |
| GDPR / PDPA | Right to erasure, consent log | Scoped into MVP data model; **active compliance deferred per §9, audit trail built now** |
| Fraud: velocity checks | Configurable per program | Required before Phase 1 launch |
| Secrets management | SOPS, runtime-injected | No secrets in images/committed files |

---

## 9. Open Questions

| Question | Blocks | Owner | Due |
|---|---|---|---|
| Qualifying window for tiers (rolling 12mo vs. calendar year)? | Phase 2 | Product | Before Phase 1 ends |
| Will points carry monetary liability on the balance sheet? (regulation deferred per TRD, but accounting stance needed) | MVP | Finance + Legal | Before MVP launch |
| Churn definition for AI personalization? | Phase 4 | Data + Product | Before Phase 3 ends |
| Email/push provider selection (for OTP + lifecycle)? | MVP | Eng | Before MVP launch |
| First federated client + protocol (OIDC assumed; any SAML need)? | Phase 2 | Eng + Product | When first SSO client lands |
| Starting-balance policy on enrollment (welcome bonus or zero)? | MVP | Product | Before MVP build |
| Self-hosted Temporal persistence (reuse PostgreSQL vs Cassandra)? | MVP/Phase 1 infra | Eng | Before Phase 1 |

---

## 10. Glossary

| Term | Definition |
|---|---|
| Earn | A member receiving points for a qualifying action |
| Burn | A member spending points to obtain a reward |
| Breakage | Points earned but never redeemed — a financial benefit to the operator |
| Ledger | Immutable record of every point earn and burn transaction |
| Idempotency | System safely handles duplicate requests without double-awarding |
| Soft currency | Points with no redemption value, used for status/gamification only |
| Qualifying window | Time period used to calculate tier status (e.g., rolling 12 months) |
| Coalition | A loyalty program shared across multiple partner brands |
| Tenant | A B2B client and its isolated data subtree (Merchant → Program → Customer) |
| B2B2C | We serve B2B clients and power their end-customers' experience |
| `auth_source` | Per-tenant member-auth mode: `local` (we hold identity) or `federated` (client IdP via OIDC) |
| Rule engine | Configurable if/then logic that computes points — currently dead code, wired in MVP |

---

*This document replaces product-analysis-20260623.md. Do not maintain both in parallel.*
*Next review: at MVP release milestone.*
