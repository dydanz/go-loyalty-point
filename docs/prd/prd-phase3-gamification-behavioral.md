# PRD: go-loyalty-point — Phase 3 (Gamification & Behavioral Earning)

**Status:** Draft 🟡 | **Target Release:** Q2 2027 | **Product Lead:** Dandi
**Phase:** 3 | **Depends on:** Phase 2 complete | **Source:** `product-analysis-v2.md` §5 Phase 3 · `trd-loyalty-platform-v2.md` Phase 4 (behavioral triggers)

---

## 1. Executive Summary & Problem Statement

**The Problem:** Members only earn when they spend money. Engagement is therefore bounded by purchase frequency — there's no reason to interact between purchases, and no viral acquisition loop.

**The Solution:** Let members earn for **non-purchase actions** — reviews, referrals, app opens — via custom API events, plus a **referral program** and **challenges/missions**. Add streaks and badges as proposed engagement boosters.

**Why Now?** With tiers driving status motivation, behavioral earning multiplies the *number of touchpoints* and adds a referral-driven acquisition channel — lowering CAC, which is the original business driver. It depends on the active event bus (from MVP) and the fraud controls (from Phase 1).

---

## 2. Target Audience

1. **The Member (B2C):** Wants more ways to earn and to be rewarded for engagement, not just spend; wants to invite friends and benefit.
2. **The Operator (B2B):** Wants to drive engagement frequency and acquire members cheaply via referrals.
3. **Platform Operator (us):** Must keep non-purchase earning from becoming a fraud/abuse vector.

---

## 3. Success Metrics & KPIs
*Within 90 days of Phase 3 launch:*

- **Engagement frequency:** +25% non-purchase earning events per active member per month.
- **Referral acquisition:** ≥10% of new members arrive via a referral code.
- **Challenge completion:** >30% of members who start a challenge complete it.
- **Abuse control:** Referral/behavioral fraud rate <2% of awarded events.

---

## 4. Scope & Non-Goals

**✅ In Scope (Phase 3):**
- **Behavioral event earning:** points from custom API events (review, referral, app open) via the event bus, real-time.
- **Referral program:** unique per-member codes; referrer + referee both earn on referee activation.
- **Challenges & missions:** multi-step tasks with a defined completion reward; progress visible.
- **Streak bonuses** `[PROPOSED]`: consecutive-activity rewards (daily/weekly).
- **Badges / achievements** `[PROPOSED]`: non-monetary milestone recognition.

**❌ Out of Scope (Non-Goals):**
- AI segmentation, churn, next-best-offer (Phase 4).
- Leaderboards, spin-the-wheel (not evidenced as needed; defer).
- Partner/coalition earning (Phase 5).

---

## 5. Core Capabilities & User Stories

### A. Member Experience

| Epic | User Story | Acceptance Criteria |
|---|---|---|
| **Behavioral Earning** | As a member, I want points for a review/app-open, so I'm rewarded for engaging. | 1. Custom API event triggers earn within 5s. 2. Event types extensible via API (not just purchases). 3. Ledger records source = event type. |
| **Referral** | As a member, I want a code that rewards me and my friend, so I invite people. | 1. Unique per-member code. 2. Both earn only on referee **activation** (not signup). 3. Referrals capped per member. |
| **Challenges** | As a member, I want a multi-step mission with a reward, so I have a goal. | 1. Challenge = steps + reward + deadline. 2. Progress visible in member UI. 3. Completion auto-grants reward + notifies. |
| **Streaks** `[PROPOSED]` | As a member, I want a bonus for consecutive activity, so I build a habit. | 1. Configurable daily/weekly streak. 2. Reset rules clear; progress shown. |

### B. Operator / System Experience

| Epic | User Story | Acceptance Criteria |
|---|---|---|
| **Event Configuration** | As an operator, I want to define which actions earn and how much, so I shape engagement. | 1. Event-type → reward mapping via rule builder. 2. Time-bound + segment-targetable. |
| **Abuse Control** | As the platform, I must stop referral/behavioral farming. | 1. Velocity checks (Phase 1) extended to non-purchase events. 2. Self-referral / duplicate-device blocked. 3. Suspicious activity alerts. |

---

## 6. Key Edge Cases & Error States

- **Self-referral / circular referral.** → Block referee = referrer; detect shared device/payment fingerprint; require genuine activation.
- **Referral awarded then referee refunds/churns immediately.** → Award gated on activation threshold (e.g., first qualifying purchase), reversible if reversed within window.
- **Behavioral event spam (script hits "app open" 1000×).** → Velocity caps per event type per window; idempotency on event ID.
- **Challenge completion race (two events complete simultaneously).** → Idempotent completion; reward granted exactly once.
- **Streak timezone ambiguity.** → Streak window evaluated in the member's/tenant's configured timezone, consistently.

---

## 7. Technical Dependencies

- **MVP event bus** active (Kafka) for real-time behavioral events.
- **Phase 1 fraud velocity checks** — extended to non-purchase events.
- **Temporal** for challenge/streak evaluation + reward granting.
- **Member SDK** must emit custom events (app open, review) from the client app.

---

**Next Steps for the PM Team:**
1. Decide which `[PROPOSED]` mechanics (streaks/badges) make the cut vs. defer.
2. Define the referee "activation" event with operators (signup vs first purchase).
3. Hand challenge + referral flows to UX; design anti-abuse messaging.
