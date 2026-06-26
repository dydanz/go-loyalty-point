# PRD: go-loyalty-point — Phase 4 (Personalization & AI)

**Status:** Draft 🟡 | **Target Release:** Q3 2027 | **Product Lead:** Dandi
**Phase:** 4 | **Depends on:** Phase 3 complete + data-volume threshold met | **Source:** `product-analysis-v2.md` §5 Phase 4 · `trd-loyalty-platform-v2.md` Phase 4

---

## 1. Executive Summary & Problem Statement

**The Problem:** Every member in the same tier sees the same offers. Rewards are not matched to individual behavior, so high-value and at-risk members are treated identically — wasting reward budget on those who'd buy anyway and missing those about to churn.

**The Solution:** Move from rule-based to **intent-driven**: member **segmentation** (RFM, tier, category affinity), **segment-based offers**, a **churn-prediction trigger** that auto-sends retention offers, **next-best-offer** recommendations, and **A/B testing** to validate reward mechanics.

**Why Now?** Only after Phases 1–3 has the platform accumulated enough behavioral data (purchases, redemptions, behavioral events) to train models. Personalization is gated on data volume — doing it earlier would produce unreliable models on thin data.

---

## 2. Target Audience

1. **The Member (B2C):** Wants relevant offers — rewards that match what they actually want, surfaced at the right moment.
2. **The Operator (B2B):** Wants to spend reward budget efficiently — retain at-risk members, not subsidize loyal ones.
3. **Data/Platform team (us):** Needs reliable models, a clear churn definition, and a way to measure lift.

---

## 3. Success Metrics & KPIs
*Within 90 days of Phase 4 launch (post data-threshold):*

- **Churn model quality:** ≥0.70 precision @ ≥0.50 recall on held-out validation; **beats a recency heuristic baseline**.
- **Retention lift:** at-risk members receiving an auto-offer retained +X pp vs. control (set per tenant).
- **Personalization lift:** segment-based offers show higher redemption rate vs. blanket offers (A/B proven, 95% confidence).
- **Budget efficiency:** reward spend per retained member down vs. Phase 3 baseline.

---

## 4. Scope & Non-Goals

**✅ In Scope (Phase 4):**
- **Member segmentation:** rule-based cohorts (RFM, tier, category affinity); daily recompute, real-time for high-value segments.
- **Segment-based offers:** different earn/redemption rules per segment.
- **Churn-prediction trigger** `[PROPOSED]`: at-risk members auto-receive a retention offer.
- **Next-best-offer** `[PROPOSED]`: AI recommends the most relevant reward at the right moment.
- **A/B testing:** test reward values/mechanics on member subsets with significance reporting.

**❌ Out of Scope (Non-Goals):**
- Partner/coalition (Phase 5).
- Real-time per-impression bidding / deep personalization beyond next-best-offer.
- Building before the **data-volume threshold** is met (gate, not a goal).

---

## 5. Core Capabilities & User Stories

### A. Member Experience

| Epic | User Story | Acceptance Criteria |
|---|---|---|
| **Relevant Offers** | As a member, I want offers that fit me, so rewards feel personal. | 1. Two members in the same tier can receive different offers based on behavior. 2. Offer surfaced at a relevant moment. |
| **Retention Offer** `[PROPOSED]` | As an at-risk member, I want a reason to stay, so I don't drift away. | 1. At-risk segment auto-receives a retention offer. 2. No manual intervention. 3. Suppressed if member recently active. |

### B. Operator / System Experience

| Epic | User Story | Acceptance Criteria |
|---|---|---|
| **Segmentation** | As an operator, I want dynamic cohorts, so I target precisely. | 1. Segments by RFM/tier/affinity/behavioral history. 2. Recompute ≥ daily; real-time for high-value. |
| **A/B Testing** | As an operator, I want to test reward mechanics, so I ship what works. | 1. Define audience split. 2. Auto-report at test end with statistical significance (95%). |
| **Churn Model** `[PROPOSED]` | As the system, I must flag at-risk members reliably. | 1. ≥0.70 precision @ ≥0.50 recall on held-out set. 2. At-risk surfaced as a targetable segment. 3. Beats recency baseline. |

---

## 6. Key Edge Cases & Error States

- **Thin data per tenant (below threshold).** → Feature gated: ship rule-based segments only; hold ML churn until ≥6mo history + ≥10k active members + ≥100k transactions per tenant.
- **Churn false positives waste budget.** → Precision bar enforced; cap retention-offer spend; require beating the heuristic baseline before auto-trigger goes live.
- **Member gets retention offer right after buying.** → Suppress if recently active; recompute at-risk on latest behavior.
- **Segment overlap / conflicting offers.** → Defined priority; one offer per member per moment.
- **PII in model training.** → Training data anonymized/pseudonymized; no raw PII in features.
- **A/B contamination (member in two tests).** → Mutually exclusive assignment; logged.

---

## 7. Technical Dependencies

- **Data-volume threshold met** (TRD): ≥6mo history + ≥10k active members + ≥100k txns/tenant.
- **ClickHouse** analytics store live (from Phase 1) as the feature/training source.
- **Churn definition** agreed: "no qualifying activity in N days," N = 2–3× median inter-purchase interval per tenant.
- **CDP / data-warehouse export** (P2 integration) for model pipelines.
- **Behavioral event history** (from Phase 3) as model features.

---

**Next Steps for the PM Team:**
1. Confirm the churn definition with Data + Product before Phase 3 ends.
2. Verify the data-volume gate per target tenant — do not start modeling below it.
3. Decide which `[PROPOSED]` AI features ship first (churn vs next-best-offer); hand A/B + segmentation UI to UX.
