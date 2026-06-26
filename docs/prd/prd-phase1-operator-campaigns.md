# PRD: go-loyalty-point — Phase 1 (Operator Control & Campaign Tools)

**Status:** Draft 🟡 | **Target Release:** Q4 2026 | **Product Lead:** Dandi
**Phase:** 1 | **Depends on:** MVP complete | **Source:** `product-analysis-v2.md` §5 Phase 1 · `trd-loyalty-platform-v2.md` Phase 3

---

## 1. Executive Summary & Problem Statement

**The Problem:** After MVP, the earn/redeem loop works but every rule and promotion still goes through us (API/engineering). Operators cannot launch a campaign, issue vouchers, or see how a program is performing on their own — and there is no points expiry, so liability never ages off and the dashboard is still fake.

**The Solution:** Give operators self-serve control: a **no-code rule builder**, an **in-house voucher engine** (bulk codes, eligibility, stacking), **campaign scheduling**, **configurable points expiry** with member warnings, **fraud velocity checks**, and a **real analytics dashboard** served from ClickHouse.

**Why Now?** Operator self-service is the product's core promise ("double points on weekends is a toggle, not a deploy"). It removes us from the critical path of every campaign and unlocks the time-to-campaign metric that justifies the platform.

---

## 2. Target Audience

1. **The Campaign Manager (B2B operator, non-technical):** Wants to configure earn rules and issue vouchers visually, schedule campaigns, and watch results — no developer.
2. **The Member (B2C):** Benefits from promotions and must be warned before points expire.
3. **Platform Operator (us):** Needs guardrails so a bad rule can't blow up liability, and fraud controls before promotions go live.

---

## 3. Success Metrics & KPIs
*Within 90 days of Phase 1 launch:*

- **Operational efficiency:** Time-to-campaign <30 min via admin UI (from "API + us" today).
- **Self-service rate:** >80% of new campaigns created without engineering involvement.
- **Redemption:** Checkout redemption drop-off <15%.
- **Expiry hygiene:** 100% of expiring points trigger 30/14/7-day member warnings.
- **Fraud:** 0 successful velocity-abuse incidents on a live campaign.

---

## 4. Scope & Non-Goals

**✅ In Scope (Phase 1):**
- **No-code rule builder** over the existing rule model (if/then, conditions, time-bound), version-controlled (author/timestamp/diff).
- **Voucher engine (in-house):** bulk unique single-use codes (≥10k/batch), eligibility (SKU/category/segment/channel/date), stacking rules, redemption logging.
- **Campaign scheduling:** start/end, timezone, auto-activation.
- **Points expiry:** configurable + 30/14/7-day warnings.
- **Fraud velocity checks:** configurable per-program limits; admin alert + account freeze/reverse.
- **Basic analytics dashboard** (ClickHouse): active members, earn rate, redemption rate.
- **Customer management UI** (member entity is server-only today).

**❌ Out of Scope (Non-Goals):**
- Tiers / status (Phase 2).
- Gamification, referrals, behavioral earning (Phase 3).
- AI/personalization, segmentation (Phase 4).
- OIDC member federation (Phase 2).
- Partner/coalition (Phase 5).

---

## 5. Core Capabilities & User Stories

### A. Operator Experience (Rules, Vouchers, Campaigns)

| Epic | User Story | Acceptance Criteria |
|---|---|---|
| **Rule Builder** | As a marketer, I want to set "2x points on weekends for SKU X," so I run a promo without code. | 1. If/then UI with condition groups (AND/OR), event triggers, time bounds. 2. No deploy required. 3. Change logged with author/timestamp/diff. |
| **Voucher Generation** | As a marketer, I want to bulk-issue unique codes, so I can run a coupon drop. | 1. ≥10,000 unique single-use codes in <60s. 2. Eligibility by SKU/category/segment/channel/date. 3. Stacking rules enforced per transaction. |
| **Campaign Scheduling** | As a marketer, I want start/end dates with auto-activation, so campaigns run unattended. | 1. Timezone-aware schedule. 2. Auto-activates/deactivates. 3. Results visible within 24h. |
| **Analytics** | As an operator, I want active members, earn rate, redemption rate, so I know if it's working. | 1. Dashboard served from ClickHouse. 2. Figures reconcile with the ledger. |

### B. Member & Platform Experience

| Epic | User Story | Acceptance Criteria |
|---|---|---|
| **Expiry Warning** | As a member, I want warning before my points expire, so I don't lose value silently. | 1. Configurable expiry per program. 2. 30/14/7-day in-app + email/push warnings. 3. Expiry runs async (Temporal), off-peak. |
| **Fraud Controls** | As the platform, I want velocity checks + override, so promos aren't abused. | 1. Configurable max earn/redeem per member per window. 2. Alert within 60s. 3. Operator can freeze account + reverse txns without engineering. |
| **Customer Lookup** | As an operator, I want to search/manage members, so I can support them. | 1. Tenant-scoped search. 2. View member balance, ledger, vouchers. |

---

## 6. Key Edge Cases & Error States

- **Mass expiry at a date boundary (e.g., Dec 31).** → Expiry processed asynchronously via Temporal queue during off-peak (e.g., 02:00), batched to avoid DB spikes.
- **Bad rule blows up liability (e.g., stackable 100% off).** → Builder validation + guardrails on save; rule version-control enables instant rollback; preview impact before activation.
- **Voucher double-redeem / shared code.** → Single-use codes enforced atomically; second redeem fails; per-code redemption cap enforced.
- **Fraud threshold mis-tuned → alert fatigue.** → Start conservative, per-program configurable; alerts deduplicated.
- **Campaign overlap (two rules match one transaction).** → Defined precedence/stacking resolution; logged which rule(s) fired.

---

## 7. Technical Dependencies

- **MVP complete:** rule engine wired, event bus active, ledger idempotent, tenant scoping.
- **ClickHouse** analytics store + Kafka→ClickHouse consumer (TRD §3.8) — self-hosted in-cluster.
- **Temporal** workflows for async expiry + scheduled campaign activation.
- **CRM sync** (via n8n) — P1 integration.
- **Admin console** front-end work for builder/voucher/campaign/dashboard.

---

**Next Steps for the PM Team:**
1. Validate rule-builder guardrails + approval flow with engineering (prevent costly misconfig).
2. Confirm ClickHouse retention/sizing (Open Question) before dashboard build.
3. Hand voucher + rule-builder flows to UX for wireframing; design the fraud-alert action surface.
