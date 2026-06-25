# PRD: go-loyalty-point — Phase 5 (Partner & Coalition)

**Status:** Gated 🔴 (do not start without a signed partner) | **Target Release:** TBD (12+ wks after Phase 4, or parallel if partner committed) | **Product Lead:** Dandi
**Phase:** 5 | **Depends on:** Phase 4 complete + **signed partner commitment** | **Source:** `product-analysis-v2.md` §5 Phase 5 · `trd-loyalty-platform-v2.md` §3.7

---

## 1. Executive Summary & Problem Statement

**The Problem:** Members earn and redeem within a single brand. There is no way to earn at one partner and redeem at another, which is the stickiness driver coalition programs (supermarket + pharmacy + fuel) rely on.

**The Solution:** Enable **multi-brand earning** into one unified balance, a **partner admin portal**, automated **financial settlement** of earn liability between partners, and **white-label member portals** so each brand keeps its own UX.

**Why Now?** Only after a single-brand program is proven and a **partner is actually committed**. This phase is **gated**: building coalition infrastructure speculatively is wasted effort. The tenant model (TRD §3.7) already keeps this a lift, not a rewrite — so we lose nothing by waiting.

---

## 2. Target Audience

1. **The Member (B2C):** Wants one balance usable across partner brands — earn at the grocer, redeem at the fuel station.
2. **The Partner Operator (B2B):** Each partner brand wants its own admin view and branded member experience.
3. **Finance (partners + us):** Need automated, auditable settlement of who owes whom for earned/redeemed liability.

---

## 3. Success Metrics & KPIs
*Within 90 days of a coalition going live:*

- **Cross-brand usage:** ≥20% of coalition members earn at >1 partner.
- **Cross-brand redemption:** ≥10% of redemptions use points earned at a different partner.
- **Settlement accuracy:** 100% of inter-partner liability reconciles automatically; 0 manual disputes unresolved >7 days.
- **Stickiness:** coalition-member retention > single-brand baseline.

---

## 4. Scope & Non-Goals

**✅ In Scope (Phase 5):**
- **Multi-brand earning** `[PROPOSED]`: points earned across partner network, pooled into one balance.
- **Partner admin portal** `[PROPOSED]`: separate operator view per partner brand.
- **Financial settlement** `[PROPOSED]`: automated earn-liability attribution + reconciliation between partners.
- **White-label member portal** `[PROPOSED]`: each brand maintains its own branded member UX.

**❌ Out of Scope (Non-Goals):**
- Anything before a signed partner — this whole phase is gated.
- Real-time inter-bank settlement (batch reconciliation is sufficient initially).
- Cross-currency coalition (single currency first).

---

## 5. Core Capabilities & User Stories

### A. Member Experience

| Epic | User Story | Acceptance Criteria |
|---|---|---|
| **Unified Balance** | As a member, I want one balance across partners, so points are simple. | 1. Earn at Partner A and redeem at Partner B on one balance. 2. Single identity across partners; tenant/coalition-scoped. |
| **Branded Experience** | As a member, I want each brand to look like itself, so it feels native. | 1. White-label portal per partner. 2. Shared balance, distinct branding. |

### B. Partner / Finance Experience

| Epic | User Story | Acceptance Criteria |
|---|---|---|
| **Partner Admin** | As a partner operator, I want my own view, so I manage my slice. | 1. Per-partner admin scope. 2. Sees own earn/redeem; aggregate-only on shared members (privacy). |
| **Settlement** | As finance, I want automated liability attribution, so partners are squared up. | 1. Earn/redeem attributed to the originating/redeeming partner. 2. Periodic reconciliation report. 3. Full audit trail. |

---

## 6. Key Edge Cases & Error States

- **Member earns at A, redeems at B, then refunds at A.** → Compensating ledger entry attributed to A; settlement adjusts; balance may go negative until re-earned.
- **Partner leaves the coalition with outstanding member balances.** → Defined wind-down: liability frozen/settled per contract; members notified.
- **Privacy across partners.** → Member PII shared only at aggregate level between partners; per-partner views never expose another partner's customer data (isolation tests).
- **Settlement dispute / mismatch.** → Immutable audit trail per transaction; reconciliation flags deltas for review.
- **Double-attribution of one earn.** → Idempotent attribution keyed by transaction ID.

---

## 7. Technical Dependencies

- **Signed partner commitment** — hard gate; nothing starts without it.
- **Phase 4 complete** (or explicit decision to run as a parallel track for a committed partner).
- **Tenant/coalition model** (TRD §3.7) — extend tenant tree to a coalition grouping.
- **Settlement/finance integration** — reconciliation export to partner finance systems.
- **White-label theming** in the member SDK/portal.

---

**Next Steps for the PM Team:**
1. **Do not start** until a partner contract is signed — confirm the gate with leadership.
2. When a partner commits: validate the settlement model with Finance + partner Finance first.
3. Scope white-label theming and per-partner admin isolation with UX + engineering.
