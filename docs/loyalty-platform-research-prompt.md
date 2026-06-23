# Prompt: Loyalty & Rewards Platform — Market Research for Product Teams

## Role & Context

You are a **Senior Product Researcher** specializing in consumer engagement platforms. You have been asked to conduct a comprehensive study of the **best-in-class loyalty, rewards, and voucher systems** currently available in the market — from enterprise SaaS platforms to open-source alternatives — so that your Product Management team can make informed decisions about building, buying, or adapting a modern loyalty platform.

Your audience is a **mixed PM and non-technical stakeholder team**. Prioritize clarity. When technical terms are unavoidable, explain them plainly in parentheses on first use.

---

## Research Scope

Cover the following platform categories:
1. **Enterprise loyalty SaaS** — Talon.One, Antavo, Yotpo, LoyaltyLion, Smile.io, Annex Cloud
2. **Open-source / composable** — Open Loyalty, Voucherify, Runa
3. **Vertically focused** — Punchh (QSR/F&B), Tango (B2B incentives), Clutch (retail)
4. **Emerging / AI-native** — Hologrow, Loyalty Juggernaut, and any platforms launched post-2024

For each, extract only what is observable and documented. Do not invent features. Flag where information is unavailable or unverified.

---

## Output Format

Produce the research in **Markdown**, using the structure below exactly. Each section heading is a standalone document suitable for sharing with different teams.

---

```markdown
# Loyalty & Rewards Platform — Market Research Report
> Researcher: [Name / AI Agent]
> Date: [Date]
> Version: 1.0

---

## DOCUMENT A: Market Landscape Overview

### A1. What Is a Modern Loyalty Platform?

Write 2–3 paragraphs explaining the evolution of loyalty systems from punch cards and static point tables to today's API-first, behavior-driven, personalized engagement engines. Use plain language. End with a one-paragraph summary of where the market is heading in 2025–2026.

### A2. Market Segmentation

| Segment | Description | Best For | Example Platforms |
|---|---|---|---|
| Enterprise Suite | Full-featured, complex rule engines, multi-market | Large retail, airline, bank | Antavo, Annex Cloud |
| Composable / API-first | Headless, developer-driven, highly flexible | Tech-forward brands | Talon.One, Voucherify |
| SMB / Plug-and-play | Quick setup, Shopify/WooCommerce native | Small e-commerce | Smile.io, LoyaltyLion |
| Vertical-specific | Built for one industry's nuances | F&B, hospitality, B2B | Punchh, Tango |
| Open-source | Self-hosted, customizable, no licensing fee | Teams with dev capacity | Open Loyalty |
| AI-native (emerging) | Intent-driven personalization, predictive rewards | Innovation-led brands | Hologrow, LJI |

### A3. Architecture Paradigms

Explain the following in plain language (2–3 sentences each):

- **Traditional (monolithic)**: How legacy systems were built and their limitations
- **MACH** (Microservices, API-first, Cloud-native, Headless): What this means for flexibility and speed
- **Composable Commerce**: The shift toward assembling best-of-breed components
- **Event-Driven / Real-time**: How modern platforms react instantly to user behavior

---

## DOCUMENT B: Feature Catalogue

### B1. Core Feature Taxonomy

A comprehensive, categorized list of features found across leading platforms. For each feature, note which platform tier typically offers it.

#### B1.1 Points & Currency Management
| Feature | Description | Who Offers It |
|---|---|---|
| Multi-currency points | Run separate point types (e.g., base points, bonus points, tier points) simultaneously | Enterprise, Composable |
| Points expiry rules | Configurable expiry: rolling, fixed date, activity-based | All tiers |
| Points transfer | Allow members to gift or pool points | Enterprise |
| Currency conversion | Exchange points for cash value, miles, crypto | Enterprise, Emerging |
| Soft currency | Non-redeemable engagement credits (e.g., "stars" for status only) | Enterprise |
| Points bank ledger | Full audit trail of every point earn/burn transaction | Enterprise, Composable |

#### B1.2 Earning Mechanics
| Feature | Description | Who Offers It |
|---|---|---|
| Transactional earning | Points per purchase (flat rate or tiered by amount) | All |
| Behavioral earning | Points for non-purchase actions: reviews, referrals, social shares, app opens | Enterprise, Composable |
| Partner earning | Earn at partner brands (co-branded programs) | Enterprise |
| Event-triggered earning | Custom events via API (e.g., completing a profile, watching a video) | Composable |
| Bonus multipliers | 2x/3x points during campaigns, on specific SKUs, or for specific tiers | All |
| Retroactive earning | Award points for past purchases after joining | Enterprise |
| Offsite earning | Points for purchases made outside the brand's own channels | Enterprise |

#### B1.3 Redemption Mechanics
| Feature | Description | Who Offers It |
|---|---|---|
| Discount redemption | Points off at checkout (fixed or percentage) | All |
| Free product redemption | Points exchanged for specific SKUs or catalog items | All |
| Voucher / coupon redemption | Points converted to single-use codes | Composable, Enterprise |
| Cashback redemption | Points paid to a payment method | Enterprise |
| Charity / donation | Redeem to donate value to causes | Enterprise, Composable |
| Experiential rewards | Redeem for events, experiences, or exclusive access | Enterprise |
| Pay-with-points (PwP) | Partial redemption at checkout (e.g., 200 pts = $2 off) | All |
| Tiered reward catalog | Different rewards unlocked by tier level | Enterprise |

#### B1.4 Tiers & Status
| Feature | Description | Who Offers It |
|---|---|---|
| Threshold-based tiers | Qualify by spend, points, or visit count in a period | All |
| Lifetime tiers | Permanent tier based on cumulative lifetime value | Enterprise |
| Tier downgrade rules | Configurable grace periods and qualifying windows | Enterprise, Composable |
| Tier benefits | Exclusive perks, multipliers, and early access by tier | All |
| Soft tier qualification | Temporary tier access before full qualification | Enterprise |
| Tier challenges | Accelerated paths to the next tier via challenges | Enterprise |

#### B1.5 Vouchers & Promotions
| Feature | Description | Who Offers It |
|---|---|---|
| Bulk voucher generation | Create thousands of unique codes in one action | Composable, Enterprise |
| Voucher stacking rules | Control which vouchers can be combined | Composable, Enterprise |
| Contextual eligibility | Vouchers valid only for specific SKUs, categories, users, or channels | All |
| Referral vouchers | Auto-generated unique codes for each referrer | All |
| Welcome / birthday offers | Triggered vouchers on lifecycle events | All |
| Dynamic value vouchers | Voucher value calculated at redemption based on cart or context | Composable |
| Influencer / partner codes | Trackable promo codes for external partners | Composable, Enterprise |

#### B1.6 Gamification
| Feature | Description | Who Offers It |
|---|---|---|
| Challenges & missions | Multi-step tasks that unlock rewards | Enterprise, Composable |
| Streaks | Bonus rewards for consecutive activity (daily, weekly) | Emerging |
| Badges & achievements | Non-monetary recognition for milestones | Enterprise, Composable |
| Leaderboards | Competitive ranking within a program | Emerging, Enterprise |
| Spin-the-wheel / scratch cards | Randomized reward moments | Enterprise |
| Progress bars | Visual indicators of progress toward next reward or tier | All (UI layer) |

#### B1.7 Personalization & AI
| Feature | Description | Who Offers It |
|---|---|---|
| Segment-based offers | Different earning/redemption rules per audience segment | Enterprise, Composable |
| Predictive churn offers | AI identifies at-risk members and auto-triggers retention offers | Enterprise, Emerging |
| Next-best-offer | AI recommends the most relevant reward at the right moment | Emerging |
| Behavioral triggers | Rules fired by user behavior events in real time | Composable, Enterprise |
| A/B testing | Test different reward values or mechanics on subsets | Composable, Enterprise |
| Dynamic personalization | Reward values and types change per individual based on profile | Emerging |

#### B1.8 Partner & Coalition Programs
| Feature | Description | Who Offers It |
|---|---|---|
| Multi-brand earning | Members earn across a network of brands | Enterprise |
| Partner dashboard | Separate admin views for coalition partners | Enterprise |
| Revenue sharing | Automated financial settlement between partners | Enterprise |
| White-label member portal | Each brand maintains its own branded experience | Enterprise |

#### B1.9 Analytics & Reporting
| Feature | Description | Who Offers It |
|---|---|---|
| Member lifecycle reports | Acquisition, activation, retention, churn funnel | Enterprise, Composable |
| Earn/burn ratio | Financial health metric: liability vs. engagement | Enterprise |
| Campaign performance | ROI of individual promotions and challenges | All |
| Cohort analysis | How different groups of members behave over time | Enterprise |
| CLV (Customer Lifetime Value) modeling | Predictive value per member | Enterprise, Emerging |
| Real-time dashboards | Live views of active campaigns and redemptions | Enterprise, Composable |
| Export / BI integration | Data export to Snowflake, BigQuery, Tableau, etc. | Enterprise, Composable |

---

## DOCUMENT C: Platform Comparison Matrix

### C1. Feature Coverage by Platform

Use ✅ (full), ⚠️ (partial/limited), ❌ (not available), ❓ (unverified) for each cell.

| Feature Area | Talon.One | Antavo | Voucherify | Open Loyalty | Smile.io | Punchh | Hologrow |
|---|---|---|---|---|---|---|---|
| Multi-currency points | | | | | | | |
| Event-triggered earning | | | | | | | |
| Voucher engine | | | | | | | |
| Tier management | | | | | | | |
| AI personalization | | | | | | | |
| Gamification | | | | | | | |
| Partner / coalition | | | | | | | |
| Real-time rules engine | | | | | | | |
| Headless / API-first | | | | | | | |
| Open-source / self-host | | | | | | | |
| Native mobile SDK | | | | | | | |
| BI / data export | | | | | | | |

### C2. Pricing Model Comparison

| Platform | Model | Entry Price (approx.) | Pricing Drivers | Notes |
|---|---|---|---|---|
| Talon.One | Usage-based | ~$700/mo | API calls, active members | No per-seat fees |
| Antavo | Contract / enterprise | Custom | Members, markets | Annual commitment |
| Voucherify | Tiered SaaS | Free tier → $149+/mo | API calls, redemptions | Transparent pricing |
| Smile.io | Tiered SaaS | $49/mo → enterprise | Orders per month | Shopify-native pricing |
| Open Loyalty | Open-source + support | Self-hosted free | Dev effort + hosting | Enterprise license available |
| Punchh | Custom | Custom | Locations, members | Restaurant-focused |

### C3. Integration Ecosystem

| Platform | Native E-commerce | CRM | CDP | ESP (Email) | POS | Mobile |
|---|---|---|---|---|---|---|
| Talon.One | Shopify, Magento, SAP | Salesforce, HubSpot | Segment | Klaviyo, Braze | Via API | SDK |
| Antavo | Via API | Salesforce | Segment, mParticle | Braze, Klaviyo | Multiple | Native iOS/Android |
| Voucherify | Shopify, WooCommerce | Salesforce, HubSpot | Segment | Klaviyo, Mailchimp | Via webhook | SDK |
| *(Continue for others)* | | | | | | |

---

## DOCUMENT D: UX / UI Analysis

### D1. Member-Facing Experience Patterns

#### The Enrollment Journey
Describe the best-practice enrollment flow observed across leading platforms. Include:
- Entry points (checkout, post-purchase, app, in-store)
- Friction reduction techniques (social login, progressive profiling)
- The "welcome moment" — how top programs deliver first value fast
- Mobile vs. web parity

#### The Rewards Discovery Experience
How do members find out what they can earn and redeem?
- Rewards catalog UX patterns (grid, list, filtered browse)
- Contextual surfacing (showing relevant rewards at the right moment)
- Transparency best practices: showing point value in real money terms
- Empty state design: what members see before they have enough points

#### The Redemption Flow
Step-by-step description of best-practice redemption UX:
1. Trigger: member sees redeemable reward
2. Selection: clarity on what they get and what it costs
3. Confirmation: double-check before committing points
4. Delivery: instant vs. delayed, digital vs. physical
5. Confirmation: receipt, balance updated, emotional payoff moment

#### The Member Dashboard / Profile
What does a gold-standard member dashboard include?
- Point balance (current, pending, expiring)
- Tier status + progress to next tier
- Transaction history
- Active vouchers and their expiry
- Personalized offers
- Referral mechanism

### D2. Merchant / Admin Console UX Patterns

What a product manager or marketer expects from the admin side:

| Admin Feature | UX Best Practice | Common Pain Points |
|---|---|---|
| Rule builder | Visual, no-code condition/action builder | Overly complex logic chains become unmaintainable |
| Campaign creation | Template library + guided wizard | Blank-canvas fatigue; no guardrails on bad configs |
| Member search & lookup | Instant full-text search with filter | Slow query on large member databases |
| Fraud dashboard | Real-time alerts, manual override | Alert fatigue from poorly tuned thresholds |
| Reporting | Scheduled exports + live dashboard | BI integration friction, data latency |
| Voucher management | Bulk import/export, status at a glance | No audit trail for manual edits |

### D3. UI Component Inventory (Common Patterns)

List the UI components and design patterns most commonly observed across loyalty frontends:

**Member-facing components**:
- Point balance widget (header/nav persistent display)
- Tier progress bar (horizontal or circular progress indicator)
- Reward card (image, name, point cost, CTA)
- Voucher tile (code, value, expiry, copy-to-clipboard)
- Activity feed / transaction history (chronological list with icons)
- Challenge card (title, progress, deadline, reward on completion)
- Referral share widget (unique link, social share buttons, tracking)
- Notification toast (earned points, unlocked reward, tier upgrade)

**Admin-facing components**:
- Rule builder canvas (condition → action flow)
- Cohort selector (segment picker with real-time member count preview)
- Campaign scheduler (calendar date picker + recurrence config)
- Voucher batch generator (CSV import / quantity selector)
- Analytics card grid (earn rate, redemption rate, active members, churn)
- Member profile drawer (slide-in panel with full member history)

### D4. Mobile Experience Patterns

| Pattern | Description | Why It Matters |
|---|---|---|
| Persistent balance display | Points shown in app header at all times | Constant reminder of value, drives engagement |
| Push notifications | Triggered alerts for points earned, offers expiring | Re-engagement without opening the app |
| Wallet integration | Apple Wallet / Google Pay for digital membership cards | Frictionless in-store identification |
| Offline earning | QR / NFC scan at POS without internet dependency | Critical for physical retail |
| Biometric auth | Face/fingerprint login for returning members | Removes friction from high-frequency interactions |

### D5. UX Anti-Patterns to Avoid

Document what leading platforms (and their users) identify as poor UX in loyalty:

| Anti-Pattern | Why It's Harmful | Better Alternative |
|---|---|---|
| Hidden point value | Members don't know what a point is worth in dollars | Always show equivalent cash value |
| Expiry anxiety without warning | Points expire without sufficient notice | 30/14/7-day email + in-app alerts |
| Redemption minimums too high | "You need 10,000 points to redeem" with slow earn rate | Set minimums at ≤1 meaningful earn session |
| Over-complicated tier rules | Too many qualifying criteria confuse members | One primary qualifying metric per tier |
| Reward catalog mismatch | Rewards don't reflect the brand or member preferences | Regularly audit catalog against redemption data |
| Forced app download for access | Web-only members can't participate | Always offer a web fallback |
| Opaque fraud blocks | Member's account frozen with no explanation | Transparent messaging + human appeal path |

---

## DOCUMENT E: Use Case Stories

Write full use case narratives for each program type. Format: user story + current best-practice flow + what differentiates great from average.

### E1. Retail / E-commerce Loyalty

**User Story**: As an online fashion shopper, I want to earn and redeem points seamlessly across purchases and returns, so that I feel rewarded for my loyalty without administrative friction.

**Best-Practice Flow**:
1. Member sees point balance in cart ("You'll earn 340 pts on this order")
2. Purchases → points awarded instantly, visible in app within seconds
3. Return processed → points deducted cleanly with notification
4. Member receives "You're 160 pts from Gold tier" nudge email 7 days later
5. Member redeems points at next checkout for $10 off without calling support

**What separates great programs**: Real-time balance visibility, frictionless returns handling, proactive tier progress communication.

### E2. QSR / F&B (Quick Service Restaurant)

**User Story**: As a coffee shop regular, I want to earn a free drink after every 10 visits, track my progress on my phone, and pay faster without fumbling for a card.

**Best-Practice Flow**:
1. Mobile app shows current "star" count on home screen
2. Member orders → scans QR code at counter or pays via app
3. Stars awarded immediately, progress bar animates to new state
4. At 10 stars: push notification "Your free drink is ready!" with one-tap redeem
5. In-store barista sees redemption on POS, no paper coupon needed

**Differentiators**: Scan-and-pay integration eliminates friction; immediate visual feedback drives habit formation.

### E3. B2B / Channel Partner Incentives

**User Story**: As a reseller rep, I want to earn bonus rewards for hitting quarterly sales targets with my vendor, so that I'm motivated to prioritize their products over competitors.

**Best-Practice Flow**:
1. Rep logs into partner portal and sees their current quarter progress vs. target
2. Qualifying sales uploaded via API from CRM; points awarded weekly
3. At target hit: email + portal notification with reward catalog link
4. Rep selects reward (gift card, travel, merchandise) from curated catalog
5. Reward fulfilled within 5 business days; confirmation email with tracking

**Differentiators**: Transparent progress visibility, fast fulfillment, tax-compliant reward delivery.

### E4. Travel & Hospitality

**User Story**: As a frequent business traveler, I want my hotel stays, restaurant bills, and spa visits to all earn the same points, and I want to use those points for free nights without blackout dates.

**Best-Practice Flow**:
1. Member checks in via app (room key on phone); stay automatically attributed
2. All on-property spend earns points (F&B, spa, minibar) via POS integration
3. Member checks dashboard mid-stay: sees points earned so far
4. Points redeemable for future bookings at real-time inventory (no blackouts)
5. Checkout email: summary of points earned, current balance, next tier progress

**Differentiators**: Unified property spend earning, no blackout dates, mobile-first check-in/out experience.

### E5. Coalition / Multi-Brand Program

**User Story**: As a supermarket loyalty member, I want to earn points at my grocery store, pharmacy, and fuel station — and redeem them across all three — with one card and one balance.

**Best-Practice Flow**:
1. Member presents one app QR at any partner location (grocery, pharmacy, fuel)
2. Points earned per partner's rules, pooled into one universal balance
3. Member redeems fuel points at petrol station via loyalty number at pump
4. Monthly statement shows earn breakdown by partner
5. Partner brands share member data only at aggregate level (privacy preserved)

**Differentiators**: Single identity across partners, cross-category redemption flexibility, privacy-preserving data architecture.

---

## DOCUMENT F: Technical Capabilities Reference

*Plain-language explanations for non-technical PMs, with depth for engineering handoff.*

### F1. Core Engine Capabilities

| Capability | Plain Language Explanation | Why It Matters for Product |
|---|---|---|
| Rule engine | A configurable "if this, then that" system for every earn/burn decision | Powers complex campaigns without code changes |
| Event streaming | The system reacts to user actions in milliseconds (not in overnight batches) | Enables real-time point awards and instant feedback |
| Idempotency | The system safely handles duplicate requests (e.g., double-tap, network retry) | Prevents members from earning points twice on one purchase |
| Webhooks | Instant outbound notifications to other systems when something happens | Enables Braze/Klaviyo to send "you earned points!" emails in real time |
| API rate limits | How many transactions per second the platform can handle | Determines if the platform survives your peak sale day |
| Multi-tenancy | Running multiple brands or regions from one platform instance | Essential for enterprise brands with regional programs |

### F2. Data Architecture Patterns

| Pattern | What It Means | Best Used When |
|---|---|---|
| Event sourcing | Every point transaction is recorded as an immutable event log | Needed for audit, dispute resolution, and retroactive recalculation |
| CQRS | Separate systems for writing (transactions) and reading (dashboards) | High-volume programs where read performance matters |
| CDC (Change Data Capture) | Database changes are streamed to other systems automatically | Keeps CRM, CDP, and analytics in sync without batch ETL |
| Ledger model | Points treated like a bank account with double-entry accounting | Required for programs with financial liability implications |

### F3. Integration Patterns

| Integration Type | How It Works | Common Use Cases |
|---|---|---|
| REST API | Request-response calls, usually for real-time earn/burn | POS integration, checkout, member lookup |
| Webhooks | Platform pushes events to your system | Triggering emails, updating CRM, fraud alerts |
| Event stream (Kafka/SQS) | High-throughput event pipe for large programs | Millions of transactions/day, real-time analytics |
| File import/export | Batch CSV/SFTP for legacy systems | ERP sync, offline POS reconciliation |
| SDK (mobile/web) | Pre-built code to embed loyalty UI in your app | Mobile app loyalty tabs, balance display |

### F4. Security & Compliance Considerations

| Area | What to Look For | Red Flags |
|---|---|---|
| Points fraud prevention | Velocity checks, device fingerprinting, manual override tools | No fraud controls or only post-hoc detection |
| Data privacy (GDPR/PDPA) | Right-to-erasure, consent management, data minimization | Vendor stores PII without clear deletion process |
| Financial compliance | Breakage accounting, gift card regulations, tax reporting for B2B rewards | No audit trail; points treated as pure marketing not financial liability |
| SSO / IAM | SAML, OAuth for admin console; MFA for member accounts | Shared admin passwords; no MFA option |
| Penetration testing | Vendor SOC 2 certification, annual pen test | No certifications; vague security documentation |

---

## DOCUMENT G: Recommendations & Decision Framework

### G1. Build vs. Buy vs. Compose

| Scenario | Recommended Approach | Rationale |
|---|---|---|
| MVP, fast time-to-market, limited budget | Buy (SMB SaaS) | Smile.io, LoyaltyLion — live in days |
| Complex rules, multiple markets, in-house tech | Compose (API-first) | Talon.One or Voucherify + your own UI |
| Enterprise, multi-brand, coalition | Buy (Enterprise Suite) | Antavo or Annex Cloud — full partner ecosystem |
| Full ownership, high transaction volume, data sovereignty | Build on Open-Source | Open Loyalty as engine, custom frontend |
| AI-first, personalization-heavy, digital-native | Emerging platforms | Hologrow, LJI — evaluate carefully (early-stage) |

### G2. Platform Selection Scorecard

Weight and score each criterion for your context (1–5):

| Criterion | Weight | Platform A | Platform B | Platform C |
|---|---|---|---|---|
| Feature completeness | 20% | /5 | /5 | /5 |
| API flexibility | 15% | /5 | /5 | /5 |
| Time to launch | 15% | /5 | /5 | /5 |
| Pricing model fit | 15% | /5 | /5 | /5 |
| Integration ecosystem | 10% | /5 | /5 | /5 |
| Mobile capability | 10% | /5 | /5 | /5 |
| Analytics / data ownership | 10% | /5 | /5 | /5 |
| Vendor stability | 5% | /5 | /5 | /5 |
| **Weighted total** | 100% | | | |

### G3. Questions to Ask Every Vendor

Use these in demos and RFP responses:

**Technical**
- What is your 99th percentile API response time at peak load?
- How do you handle double-spend / race conditions in point redemptions?
- Can we access raw event-level data for our own data warehouse?
- What is your disaster recovery / RTO?

**Commercial**
- What happens to our data if we churn?
- Is there a minimum contract term?
- How are price increases handled at renewal?

**Product**
- What is your product roadmap for AI/personalization features?
- How many engineers do customers typically need to integrate your platform?
- Can you provide three reference customers in our industry?

---

## DOCUMENT H: Open Research Questions

Areas requiring primary research (user interviews, vendor demos, or pilot programs) that cannot be answered from desk research alone:

- [ ] What loyalty mechanics resonate most with Indonesian / Southeast Asian consumers specifically?
- [ ] What is the realistic cost of the earn/burn liability for a new program in our vertical?
- [ ] Which platforms have succeeded with programs similar to our use case (EV / mobility)?
- [ ] How do members in our target demographic prefer to interact: app, WhatsApp, web, or in-person?
- [ ] What is the typical activation rate (% of enrollees who earn at least once) in comparable programs?
- [ ] What fraud patterns should we design against first?

---

## DOCUMENT I: Glossary of Loyalty Terms

| Term | Plain-Language Definition |
|---|---|
| Earn rate | How many points you get per currency unit spent (e.g., 1 point per $1) |
| Burn rate | The point cost of a reward (e.g., 500 points = $5 off) |
| Breakage | Points that are earned but never redeemed — a revenue benefit for the operator |
| Soft currency | Points or credits used only for status / gamification, with no redemption value |
| Hard currency | Points with real monetary value that create financial liability |
| Tier qualification window | The time period in which a member must accumulate enough spend/points to reach or maintain a tier |
| Coalition program | A loyalty program shared across multiple brands — members earn and spend at all partners |
| Headless loyalty | Loyalty logic (rules, points engine) decoupled from the frontend — your UI, your rules engine behind it |
| Idempotency key | A mechanism that ensures a transaction can be safely retried without double-awarding points |
| Event-driven | The loyalty system reacts to real-time signals (purchases, logins, reviews) rather than nightly batches |
| Voucher stacking | When multiple discount codes or offers are applied simultaneously — platforms control this with stacking rules |
| Redemption friction | Any barrier that makes it harder for a member to use their rewards |
| CLV (Customer Lifetime Value) | The total revenue a business can expect from one customer over the course of their relationship |
| Churn prediction | AI/ML model that identifies members at risk of becoming inactive |
| MACH | Architecture principle: Microservices, API-first, Cloud-native, Headless |

---

*End of Report*
```

---

## Research Methodology Notes

When using this prompt with AI assistance or for live research, follow this sequence:

1. **Ground truth first** — Search vendor documentation, G2/Capterra reviews, and case studies before relying on AI knowledge.
2. **Verify pricing** — Pricing changes quarterly; always web-search current pricing before publishing.
3. **Flag uncertainty** — Use ❓ in comparison tables where data is estimated or outdated.
4. **Primary source preference**: vendor docs > third-party review sites > analyst reports > AI synthesis.
5. **Mark AI-synthesized content** — Any section generated without primary source verification should note "(AI synthesis — verify before publishing)".
6. **For the full step-by-step research workflow** — vendor scoring, demo checklist, primary source verification, and stakeholder deliverable formats — see `loyalty-research-methodology.md`.
7. **Output file** - put in folder ./docs as markdown format files
