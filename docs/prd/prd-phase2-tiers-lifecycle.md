# PRD: go-loyalty-point — Phase 2 (Tiers & Member Lifecycle)

**Status:** Draft 🟡 | **Target Release:** Q1 2027 | **Product Lead:** Dandi
**Phase:** 2 | **Depends on:** Phase 1 complete | **Source:** `product-analysis-v2.md` §5 Phase 2 · `trd-loyalty-platform-v2.md` Phase 2 + §3.10

---

## 1. Executive Summary & Problem Statement

**The Problem:** Members earn and redeem, and operators run campaigns — but there is no **status**. Every member is treated identically, so there is no long-term motivation to keep coming back, and no way for a client to reward their best customers differently. Larger B2B clients also want their customers to log in via their own SSO, which the local-only auth from MVP doesn't support.

**The Solution:** A configurable **tier system** (thresholds, qualification engine, benefits, progress display, downgrade grace), plus the **OIDC federation** half of the hybrid member-auth model so SSO-demanding clients can onboard.

**Why Now?** Tiers are the standard retention lever in loyalty — they convert transactional members into status-seeking ones. With the core loop and operator tooling solid, status is the highest-leverage engagement feature, and the first enterprise clients will gate on SSO.

---

## 2. Target Audience

1. **The Member (B2C):** Wants to know their status, the benefits it unlocks, and how close they are to the next tier — and to be warned before losing it.
2. **The Operator (B2B):** Wants to define tiers and benefits that fit their business and reward their best customers.
3. **The Enterprise Client (B2B):** Wants their customers to authenticate via the client's own IdP (SSO), not a separate login.

---

## 3. Success Metrics & KPIs
*Within 90 days of Phase 2 launch:*

- **Retention lift:** Tier-qualified members retained +10pp vs. the non-tiered base.
- **Progression:** >40% of active members view their tier-progress UI monthly.
- **Downgrade fairness:** 100% of at-risk members receive a downgrade warning before tier loss.
- **SSO adoption:** ≥1 client live on OIDC federation; 0 cross-tenant identity leaks.

---

## 4. Scope & Non-Goals

**✅ In Scope (Phase 2):**
- **Tier definition:** configurable tiers with qualifying thresholds (spend / points / visits).
- **Tier qualification engine:** automated promote/demote on rolling or fixed window; runs on transaction + on interval (Temporal).
- **Tier benefits:** multipliers, exclusive rewards, early access by tier.
- **Tier progress display:** visual progress bar to next tier.
- **Downgrade grace period:** configurable warning window + notification.
- **OIDC member federation:** `auth_source = federated` per tenant (hybrid auth, TRD §3.10).

**❌ Out of Scope (Non-Goals):**
- Tier challenges / accelerators `[PROPOSED]` (defer; gamification-adjacent).
- Behavioral/non-purchase earning, referrals (Phase 3).
- AI segmentation/personalization (Phase 4).
- SAML federation (OIDC only unless a client requires it).
- Partner/coalition tiers (Phase 5).

---

## 5. Core Capabilities & User Stories

### A. Member Experience

| Epic | User Story | Acceptance Criteria |
|---|---|---|
| **Tier Status** | As a member, I want to see my tier and its benefits, so I value my status. | 1. Current tier + benefits shown in member UI. 2. Benefits (multiplier/perks) actually applied on earn. |
| **Tier Progress** | As a member, I want a progress bar to the next tier, so I'm motivated to keep going. | 1. Shows metric + amount required to next tier. 2. Celebratory state on upgrade. |
| **Downgrade Warning** | As a member, I want warning before losing my tier, so I can act. | 1. Notification sent within the configured grace window. 2. Clear "do X to keep your tier" message. |
| **SSO Login** | As a member of an enterprise client, I want to log in with my existing account, so I don't make a new one. | 1. OIDC redirect to the client IdP. 2. Signed assertion (`sub` + `tenant_id`) verified against tenant JWKS. 3. No credential stored by us. |

### B. Operator / System Experience

| Epic | User Story | Acceptance Criteria |
|---|---|---|
| **Tier Definition** | As an operator, I want to define tiers and thresholds, so status fits my business. | 1. ≥3 configurable tiers. 2. One primary qualifying metric per tier (anti-confusion). 3. Rolling or fixed window selectable. |
| **Qualification Engine** | As the system, I must promote/demote members automatically and correctly. | 1. Re-evaluates on qualifying transaction + scheduled interval (Temporal). 2. Idempotent; no flapping. 3. Emits tier-change event for notification. |
| **Federation Config** | As an operator (enterprise), I want to register my IdP, so my customers use SSO. | 1. Per-tenant issuer/JWKS/client-id/claim-map. 2. Validated at onboarding. 3. Falls back cleanly if `auth_source = local`. |

---

## 6. Key Edge Cases & Error States

- **Member oscillates around a threshold (promote/demote flapping).** → Hysteresis + grace window; demotion only after the qualifying window closes, not instantly.
- **Retroactive transaction or refund changes qualification.** → Re-evaluation is idempotent; recompute from the ledger, not incremental counters.
- **Benefit timing (when does a new multiplier apply?).** → Multiplier applies from the moment of qualification forward; never retroactively recomputes past earns.
- **OIDC token from wrong tenant / spoofed issuer.** → Reject if issuer not the tenant's registered IdP; `tenant_id` claim must match; isolation test asserts no cross-tenant mapping.
- **Mid-stay tier change for a transaction in flight.** → Tier resolved at transaction commit; consistent within a single transaction.

---

## 7. Technical Dependencies

- **Phase 1 complete:** rule engine + event bus + analytics.
- **Temporal** for scheduled tier evaluation + downgrade-warning workflows.
- **OIDC libraries** + per-tenant IdP config store (TRD §3.10).
- **Notification path** (Temporal/n8n) for tier upgrade/downgrade messages.
- **Resolve Open Question:** tier qualifying window (rolling 12mo vs calendar year) before build.

---

**Next Steps for the PM Team:**
1. Resolve the qualifying-window Open Question with Product before build.
2. Confirm first federated client + that OIDC (not SAML) suffices.
3. Hand tier-progress + status UI and SSO flow to UX for wireframing.
