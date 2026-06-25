# Loyalty & Rewards Platform — Market Research Report
> Researcher: Claude (AI Agent) — Senior Product Researcher persona
> Date: 2026-06-24
> Version: 1.0

> **Sourcing note.** Pricing and platform-capability claims below were spot-checked against vendor sites and third-party review aggregators (G2, Capterra, GetApp) in June 2026. Where a claim was **not** verified against a primary source, it is tagged **(AI synthesis — verify before publishing)**. Pricing changes quarterly — re-verify before any external use. Comparison-matrix cells use ❓ wherever a capability could not be confirmed from public documentation.

---

## DOCUMENT A: Market Landscape Overview

### A1. What Is a Modern Loyalty Platform?

A loyalty platform is the software that runs a rewards program: it tracks who your customers are, what they do, how many points they have, and what they can get in return. Twenty years ago this was a punch card or a plastic card with a barcode, backed by a simple database that added points overnight in a batch job. Rules were fixed ("1 point per dollar, 500 points = $5 off"), the same for everyone, and changing them meant a developer ticket and a release.

Today's platforms are **API-first** (other systems talk to them through a standard web interface), **event-driven** (they react to a customer action — a purchase, an app open, a review — in milliseconds rather than overnight), and **behavior-driven** (rewards respond to what a person actually does, not just how much they spend). The points engine is increasingly **headless** — the rules and balances live in the platform, but the brand builds its own customer-facing screens on top. This lets a marketer launch a "triple points this weekend on sneakers" campaign through a visual editor, with no engineering involved, and have it live in minutes.

The leading edge in **2025–2026 is AI and personalization**. Vendors now market "AI loyalty" as the headline: predicting which customers are about to churn and auto-triggering a retention offer, recommending the single best reward for each person at the right moment, and using conversational AI agents to help marketers design programs (Antavo's "Timi AI," Loyalty Juggernaut's "Agentic AI Compass," Hologrow's intent-detection engine). The market direction is clear: away from one-size-fits-all point tables, toward real-time, individualized, predictive engagement — with a parallel rise in *composable* approaches where brands assemble a program from best-of-breed parts rather than buying one monolith.

### A2. Market Segmentation

| Segment | Description | Best For | Example Platforms |
|---|---|---|---|
| Enterprise Suite | Full-featured, complex rule engines, multi-market, partner ecosystems | Large retail, airline, bank | Antavo, Annex Cloud, Loyalty Juggernaut |
| Composable / API-first | Headless, developer-driven, highly flexible incentive engines | Tech-forward brands | Talon.One, Voucherify |
| SMB / Plug-and-play | Quick setup, Shopify/WooCommerce native | Small e-commerce | Smile.io, LoyaltyLion, Yotpo |
| Vertical-specific | Built for one industry's nuances | F&B/QSR, hospitality, B2B | Punchh (QSR), Tango (B2B incentives), Clutch (retail) |
| Open-source | Self-hosted, customizable, no licensing fee for core | Teams with dev capacity | Open Loyalty |
| AI-native (emerging) | Intent-driven personalization, predictive rewards | Innovation-led brands | Hologrow, Loyalty Juggernaut (GRAVTY) |

### A3. Architecture Paradigms

- **Traditional (monolithic).** One large, tightly-coupled system handling rules, balances, UI, and reporting together, often updating points in nightly batches. Reliable but rigid: any change touches the whole system, customer feedback is delayed (you earned points yesterday, you see them today), and scaling one part means scaling all of it.

- **MACH** (Microservices, API-first, Cloud-native, Headless). The platform is split into small independent services (microservices), every function is reachable through a documented API, it runs on elastic cloud infrastructure, and the logic is decoupled from any specific frontend (headless). The payoff: you can change or scale one piece without touching the rest, and ship faster.

- **Composable Commerce.** Rather than buying one suite that does everything adequately, you assemble best-of-breed components — a points engine here, a voucher engine there, your own CDP and email tool — connected by APIs. More flexibility and less vendor lock-in, at the cost of needing engineering capacity to integrate and maintain the seams.

- **Event-Driven / Real-time.** The system listens to a stream of events (purchase made, profile completed, cart abandoned) and reacts instantly — awarding points, firing a webhook, triggering an offer. This is what makes "you just earned 340 points" appear in the app within seconds and enables behavioral triggers that feel responsive rather than delayed.

---

## DOCUMENT B: Feature Catalogue

### B1. Core Feature Taxonomy

> The "Who Offers It" column reflects the *typical* tier where a feature is found, synthesized from vendor documentation and review sites. Treat as a general guide, not a per-vendor guarantee. **(Partial AI synthesis — verify per vendor before publishing.)**

#### B1.1 Points & Currency Management
| Feature | Description | Who Offers It |
|---|---|---|
| Multi-currency points | Run separate point types (base, bonus, tier points) simultaneously | Enterprise, Composable |
| Points expiry rules | Configurable expiry: rolling, fixed date, activity-based | All tiers |
| Points transfer | Members gift or pool points | Enterprise |
| Currency conversion | Exchange points for cash value, miles, crypto | Enterprise, Emerging |
| Soft currency | Non-redeemable engagement credits ("stars" for status only) | Enterprise |
| Points bank ledger | Full audit trail of every earn/burn transaction | Enterprise, Composable |

#### B1.2 Earning Mechanics
| Feature | Description | Who Offers It |
|---|---|---|
| Transactional earning | Points per purchase (flat or tiered by amount) | All |
| Behavioral earning | Points for reviews, referrals, social shares, app opens | Enterprise, Composable |
| Partner earning | Earn at partner brands (co-branded) | Enterprise |
| Event-triggered earning | Custom events via API (profile completion, video watched) | Composable |
| Bonus multipliers | 2x/3x during campaigns, on SKUs, or for tiers | All |
| Retroactive earning | Award points for past purchases after joining | Enterprise |
| Offsite earning | Points for purchases outside brand's own channels | Enterprise |

#### B1.3 Redemption Mechanics
| Feature | Description | Who Offers It |
|---|---|---|
| Discount redemption | Points off at checkout (fixed or %) | All |
| Free product redemption | Points exchanged for specific SKUs / catalog items | All |
| Voucher / coupon redemption | Points converted to single-use codes | Composable, Enterprise |
| Cashback redemption | Points paid to a payment method | Enterprise |
| Charity / donation | Redeem to donate to causes | Enterprise, Composable |
| Experiential rewards | Redeem for events, experiences, exclusive access | Enterprise |
| Pay-with-points (PwP) | Partial redemption at checkout (200 pts = $2 off) | All |
| Tiered reward catalog | Rewards unlocked by tier level | Enterprise |

#### B1.4 Tiers & Status
| Feature | Description | Who Offers It |
|---|---|---|
| Threshold-based tiers | Qualify by spend, points, or visits in a period | All |
| Lifetime tiers | Permanent tier on cumulative lifetime value | Enterprise |
| Tier downgrade rules | Configurable grace periods and qualifying windows | Enterprise, Composable |
| Tier benefits | Exclusive perks, multipliers, early access by tier | All |
| Soft tier qualification | Temporary tier access before full qualification | Enterprise |
| Tier challenges | Accelerated paths to next tier via challenges | Enterprise |

#### B1.5 Vouchers & Promotions
| Feature | Description | Who Offers It |
|---|---|---|
| Bulk voucher generation | Thousands of unique codes in one action | Composable, Enterprise |
| Voucher stacking rules | Control which vouchers combine | Composable, Enterprise |
| Contextual eligibility | Valid only for specific SKUs, categories, users, channels | All |
| Referral vouchers | Auto-generated unique codes per referrer | All |
| Welcome / birthday offers | Triggered vouchers on lifecycle events | All |
| Dynamic value vouchers | Value calculated at redemption based on cart/context | Composable |
| Influencer / partner codes | Trackable promo codes for external partners | Composable, Enterprise |

#### B1.6 Gamification
| Feature | Description | Who Offers It |
|---|---|---|
| Challenges & missions | Multi-step tasks that unlock rewards | Enterprise, Composable |
| Streaks | Bonus rewards for consecutive activity | Emerging, Enterprise |
| Badges & achievements | Non-monetary recognition for milestones | Enterprise, Composable |
| Leaderboards | Competitive ranking within a program | Emerging, Enterprise |
| Spin-the-wheel / scratch cards | Randomized reward moments | Enterprise |
| Progress bars | Visual progress toward next reward/tier | All (UI layer) |

#### B1.7 Personalization & AI
| Feature | Description | Who Offers It |
|---|---|---|
| Segment-based offers | Different rules per audience segment | Enterprise, Composable |
| Predictive churn offers | AI flags at-risk members, auto-triggers retention | Enterprise, Emerging |
| Next-best-offer | AI recommends the most relevant reward at the right moment | Emerging, Enterprise |
| Behavioral triggers | Rules fired by user-behavior events in real time | Composable, Enterprise |
| A/B testing | Test reward values/mechanics on subsets | Composable, Enterprise |
| Dynamic personalization | Reward values/types change per individual | Emerging |

#### B1.8 Partner & Coalition Programs
| Feature | Description | Who Offers It |
|---|---|---|
| Multi-brand earning | Earn across a network of brands | Enterprise |
| Partner dashboard | Separate admin views for coalition partners | Enterprise |
| Revenue sharing | Automated settlement between partners | Enterprise |
| White-label member portal | Each brand keeps its own branded experience | Enterprise |

#### B1.9 Analytics & Reporting
| Feature | Description | Who Offers It |
|---|---|---|
| Member lifecycle reports | Acquisition → activation → retention → churn funnel | Enterprise, Composable |
| Earn/burn ratio | Liability-vs-engagement health metric | Enterprise |
| Campaign performance | ROI of individual promotions and challenges | All |
| Cohort analysis | How member groups behave over time | Enterprise |
| CLV modeling | Predictive value per member | Enterprise, Emerging |
| Real-time dashboards | Live views of active campaigns and redemptions | Enterprise, Composable |
| Export / BI integration | Export to Snowflake, BigQuery, Tableau | Enterprise, Composable |

---

## DOCUMENT C: Platform Comparison Matrix

✅ full · ⚠️ partial/limited · ❌ not available · ❓ unverified

### C1. Feature Coverage by Platform

> Cells reflect public documentation and review-site evidence as of June 2026. ❓ = not confirmable from public sources. **(Partial AI synthesis — verify before publishing.)**

| Feature Area | Talon.One | Antavo | Voucherify | Open Loyalty | Smile.io | Punchh | Hologrow |
|---|---|---|---|---|---|---|---|
| Multi-currency points | ✅ | ✅ | ✅ | ✅ | ⚠️ | ✅ | ❓ |
| Event-triggered earning | ✅ | ✅ | ✅ | ✅ | ⚠️ | ✅ | ✅ |
| Voucher engine | ✅ | ✅ | ✅ | ✅ | ⚠️ | ✅ | ⚠️ |
| Tier management | ⚠️ | ✅ | ⚠️ | ✅ | ✅ | ✅ | ❓ |
| AI personalization | ⚠️ | ✅ | ⚠️ | ⚠️ | ❌ | ✅ | ✅ |
| Gamification | ⚠️ | ✅ | ⚠️ | ✅ | ⚠️ | ✅ | ⚠️ |
| Partner / coalition | ⚠️ | ✅ | ⚠️ | ✅ | ❌ | ⚠️ | ❓ |
| Real-time rules engine | ✅ | ✅ | ✅ | ✅ | ⚠️ | ✅ | ✅ |
| Headless / API-first | ✅ | ✅ | ✅ | ✅ | ⚠️ | ⚠️ | ⚠️ |
| Open-source / self-host | ❌ | ❌ | ❌ | ✅ | ❌ | ❌ | ❌ |
| Native mobile SDK | ⚠️ | ✅ | ⚠️ | ⚠️ | ⚠️ | ✅ | ❓ |
| BI / data export | ✅ | ✅ | ✅ | ✅ | ⚠️ | ✅ | ❓ |

Notes:
- **Talon.One** is an incentives/promotion engine first; tiers and loyalty constructs are supported but it is most differentiated on its rules engine and campaign logic. Loyalty-program features expanded over recent releases.
- **Smile.io** is deliberately SMB/Shopify-focused; API access and advanced features are gated to the Enterprise plan.
- **Punchh (PAR)** is QSR/restaurant-specialized; "headless" is partial because it is sold as an integrated guest-engagement suite rather than a pure API product.
- **Hologrow** is early-stage and AI-native; several enterprise constructs are unverified from public docs (hence ❓).

### C2. Pricing Model Comparison

> Verified June 2026 against vendor pricing pages and review aggregators. Prices change quarterly — re-verify.

| Platform | Model | Entry Price (approx.) | Pricing Drivers | Notes |
|---|---|---|---|---|
| Talon.One | Usage-based, custom | Custom (enterprise; no public list price — reports of ~$1,500/mo basic, unconfirmed) | Data volume, active members, API usage | Demo required; all plans unlock full product. **(Entry price ❓ — not officially published.)** |
| Antavo | Contract / enterprise | Custom | Members, markets | Annual commitment; no public pricing |
| Voucherify | Tiered SaaS + free | Free ($0) → Startup $170/mo → Growth $399/mo → Professional $599/mo → Enterprise custom | API calls/month, projects | Transparent public pricing; 60-day trial |
| Smile.io | Tiered SaaS + free | Free → Starter $49/mo → Growth $199/mo → Pro $599/mo → Enterprise from ~$1k/mo | Plan tier; API access only on Enterprise | Shopify-native; VIP tiers on Pro+ |
| Open Loyalty | Open-source + paid support | Self-hosted core free | Dev effort + hosting; enterprise license/support add-on | Full control, needs engineering capacity |
| Punchh (PAR) | Custom | Custom | Locations, members | Restaurant/QSR-focused; quote-based |
| Loyalty Juggernaut (GRAVTY) | Enterprise / cloud | Custom (also AWS Marketplace) | Members, transactions, modules | Serverless, high-TPS; SOC/ISO certified |
| Hologrow | SaaS (AI-native) | Public pricing not confirmed | Likely usage/members | Early-stage; verify in demo. **(❓)** |

### C3. Integration Ecosystem

> Synthesized from vendor integration docs/listings as of June 2026. Confirm specific connectors in a demo. **(Partial AI synthesis.)**

| Platform | Native E-commerce | CRM | CDP | ESP (Email) | POS | Mobile |
|---|---|---|---|---|---|---|
| Talon.One | Shopify, Magento, SAP (via API/connectors) | Salesforce, HubSpot | Segment | Klaviyo, Braze | Via API | SDK / API |
| Antavo | Via API + connectors | Salesforce | Segment, mParticle | Braze, Klaviyo | Multiple | Native iOS/Android SDK |
| Voucherify | Shopify, WooCommerce | Salesforce, HubSpot | Segment | Klaviyo, Mailchimp, Braze | Via webhook/API | SDK |
| Open Loyalty | Via API/webhooks (commercetools accelerator) | Via API | Via API | Via API | Via API | SDK / via API |
| Smile.io | Shopify, BigCommerce, Wix | ❓ (limited) | ❓ | Klaviyo, Mailchimp | ❌ (online-first) | ⚠️ (web widget; no full native SDK) |
| Punchh (PAR) | PAR Ordering, first-party channels | PAR ecosystem | ❓ | Email/SMS built-in | Deep POS (PAR + 3rd party) | Native app + Apple/Google Wallet (Smart Passes) |
| Hologrow | Shopify + e-commerce (per docs) | ❓ | ❓ | ❓ | ❓ | ❓ |

---

## DOCUMENT D: UX / UI Analysis

### D1. Member-Facing Experience Patterns

#### The Enrollment Journey
- **Entry points.** Best programs enroll at the highest-intent moment: at checkout ("Create an account to earn 340 points on this order"), post-purchase confirmation, in-app onboarding, and in-store via QR/staff prompt. Multiple entry points beat a single "join" page.
- **Friction reduction.** Social/SSO login, one-tap enrollment for existing customers, and *progressive profiling* (collect name + email now, birthday and preferences later) instead of a long upfront form.
- **The "welcome moment."** Top programs deliver value within seconds: instant welcome bonus points, a first-order voucher, or immediate tier status. Seeing a non-zero balance on day one drives activation.
- **Mobile vs. web parity.** Members expect identical balance, history, and redemption on web and app. Forcing an app download to participate is an anti-pattern (see D5).

#### The Rewards Discovery Experience
- **Catalog UX.** Card grid with image, name, point cost, and a clear CTA; filter/sort by category, point cost, and affordability ("show what I can afford now").
- **Contextual surfacing.** Show the relevant reward at the relevant moment — "You're 60 points from a free coffee" on the cart page, not buried in a rewards tab.
- **Transparency.** Always show point value in real money ("500 pts = $5"). Hidden value erodes trust.
- **Empty states.** Before a member can afford anything, show *progress and a path*: "Earn 160 more points to unlock your first reward," plus easy earn actions — not a blank "no rewards available."

#### The Redemption Flow
1. **Trigger.** Member sees a reward they can afford, clearly marked as redeemable now.
2. **Selection.** Crystal-clear "you get X, it costs Y points," remaining balance shown.
3. **Confirmation.** A deliberate confirm step before committing points (prevents accidental spend).
4. **Delivery.** Instant for digital (code/voucher applied immediately); clear ETA for physical/experiential.
5. **Payoff.** Receipt + updated balance + an emotional moment (animation, "Enjoy!", celebratory toast). The dopamine hit drives repeat engagement.

#### The Member Dashboard / Profile
Gold-standard dashboard includes:
- Point balance — **current, pending, and expiring soon** (three distinct numbers).
- Tier status + visual progress to next tier.
- Transaction history (earn and burn, chronological, with icons).
- Active vouchers with expiry dates.
- Personalized offers tailored to the member.
- Referral mechanism (share link + tracking).

### D2. Merchant / Admin Console UX Patterns

| Admin Feature | UX Best Practice | Common Pain Points |
|---|---|---|
| Rule builder | Visual, no-code condition→action builder | Complex logic chains become unmaintainable; hard to debug "why did this fire?" |
| Campaign creation | Template library + guided wizard | Blank-canvas fatigue; no guardrails against costly misconfig (e.g. stackable 100% discount) |
| Member search & lookup | Instant full-text search + filters | Slow queries on large member bases |
| Fraud dashboard | Real-time alerts + manual override | Alert fatigue from poorly-tuned thresholds |
| Reporting | Scheduled exports + live dashboards | BI integration friction, data latency |
| Voucher management | Bulk import/export, status at a glance | No audit trail for manual edits |

### D3. UI Component Inventory (Common Patterns)

**Member-facing components**
- Point balance widget (persistent header/nav display)
- Tier progress bar (horizontal or circular indicator)
- Reward card (image, name, point cost, CTA)
- Voucher tile (code, value, expiry, copy-to-clipboard)
- Activity feed / transaction history (chronological list with icons)
- Challenge card (title, progress, deadline, completion reward)
- Referral share widget (unique link, social buttons, tracking)
- Notification toast (points earned, reward unlocked, tier upgrade)

**Admin-facing components**
- Rule builder canvas (condition → action flow)
- Cohort selector (segment picker with live member-count preview)
- Campaign scheduler (calendar picker + recurrence config)
- Voucher batch generator (CSV import / quantity selector)
- Analytics card grid (earn rate, redemption rate, active members, churn)
- Member profile drawer (slide-in panel with full history)

### D4. Mobile Experience Patterns

| Pattern | Description | Why It Matters |
|---|---|---|
| Persistent balance display | Points always visible in app header | Constant value reminder; drives engagement |
| Push notifications | Alerts for points earned, offers expiring | Re-engagement without opening the app |
| Wallet integration | Apple Wallet / Google Wallet membership cards (e.g. Punchh Smart Passes) | Frictionless in-store identification |
| Offline earning | QR / NFC scan at POS without live internet | Critical for physical retail |
| Biometric auth | Face/fingerprint login | Removes friction for high-frequency users |

### D5. UX Anti-Patterns to Avoid

| Anti-Pattern | Why It's Harmful | Better Alternative |
|---|---|---|
| Hidden point value | Members don't know what a point is worth | Always show cash equivalent |
| Expiry without warning | Points vanish silently → anger, churn | 30/14/7-day email + in-app alerts |
| Redemption minimums too high | "10,000 points to redeem" feels unreachable | Minimum ≤ ~1 meaningful earn session |
| Over-complicated tier rules | Multiple qualifying criteria confuse members | One primary qualifying metric per tier |
| Reward catalog mismatch | Rewards don't fit brand or member taste | Audit catalog against actual redemption data |
| Forced app download | Web-only members locked out | Always offer a web fallback |
| Opaque fraud blocks | Account frozen with no explanation | Transparent messaging + human appeal path |

---

## DOCUMENT E: Use Case Stories

### E1. Retail / E-commerce Loyalty
**User Story.** As an online fashion shopper, I want to earn and redeem points seamlessly across purchases and returns, so I feel rewarded without administrative friction.

**Best-Practice Flow.**
1. Member sees balance in cart: "You'll earn 340 pts on this order."
2. Purchase → points awarded instantly, visible in app within seconds.
3. Return processed → points deducted cleanly with a clear notification.
4. 7 days later: "You're 160 pts from Gold tier" nudge email.
5. Member redeems for $10 off at next checkout — no support call.

**What separates great programs.** Real-time balance visibility, clean returns handling (no negative-balance surprises), proactive tier-progress nudges.

### E2. QSR / F&B (Quick Service Restaurant)
**User Story.** As a coffee shop regular, I want a free drink after every 10 visits, tracked on my phone, and to pay faster without fumbling for a card.

**Best-Practice Flow.**
1. App home screen shows current "star" count.
2. Member orders → scans QR at counter or pays in-app.
3. Stars awarded immediately; progress bar animates.
4. At 10 stars: push "Your free drink is ready!" with one-tap redeem.
5. Barista sees redemption on POS — no paper coupon.

**Differentiators.** Scan-and-pay removes friction; instant visual feedback builds habit. Wallet-based passes (e.g. PAR Punchh Smart Passes) skip the app entirely for identification.

### E3. B2B / Channel Partner Incentives
**User Story.** As a reseller rep, I want bonus rewards for hitting quarterly sales targets, so I prioritize this vendor's products.

**Best-Practice Flow.**
1. Rep logs into partner portal, sees quarter progress vs. target.
2. Qualifying sales uploaded via API from CRM; points awarded weekly.
3. Target hit → email + portal notification with reward catalog link.
4. Rep picks reward (gift card, travel, merchandise) from a curated catalog.
5. Fulfilled within 5 business days; confirmation + tracking email.

**Differentiators.** Transparent progress visibility, fast fulfillment, tax-compliant reward delivery (1099/local-tax handling). Platforms like Tango specialize in B2B reward fulfillment.

### E4. Travel & Hospitality
**User Story.** As a frequent business traveler, I want hotel stays, restaurant bills, and spa visits to earn the same points, redeemable for free nights without blackout dates.

**Best-Practice Flow.**
1. Member checks in via app (room key on phone); stay auto-attributed.
2. All on-property spend (F&B, spa, minibar) earns points via POS integration.
3. Mid-stay: dashboard shows points earned so far.
4. Points redeemable against real-time inventory — no blackouts.
5. Checkout email: points earned, balance, next-tier progress.

**Differentiators.** Unified property-spend earning, no blackout dates, mobile-first check-in/out.

### E5. Coalition / Multi-Brand Program
**User Story.** As a supermarket loyalty member, I want to earn at grocery, pharmacy, and fuel — and redeem across all three — with one card and one balance.

**Best-Practice Flow.**
1. One app QR works at any partner (grocery, pharmacy, fuel).
2. Points earned per partner's rules, pooled into one universal balance.
3. Redeem fuel points at the pump via loyalty number.
4. Monthly statement shows earn breakdown by partner.
5. Partners share member data only at aggregate level (privacy preserved).

**Differentiators.** Single identity across partners, cross-category redemption, privacy-preserving architecture, automated inter-partner settlement.

---

## DOCUMENT F: Technical Capabilities Reference

### F1. Core Engine Capabilities

| Capability | Plain Language | Why It Matters for Product |
|---|---|---|
| Rule engine | Configurable "if this, then that" for every earn/burn decision | Complex campaigns without code changes |
| Event streaming | Reacts to actions in milliseconds, not nightly batches | Real-time awards and instant feedback |
| Idempotency | Safely handles duplicate requests (double-tap, retry) | Members don't earn twice on one purchase |
| Webhooks | Instant outbound notifications when something happens | Braze/Klaviyo send "you earned points!" in real time |
| API rate limits | How many transactions/second the platform handles | Determines survival of your peak sale day |
| Multi-tenancy | One platform instance runs multiple brands/regions | Essential for enterprise/regional programs |

> **Benchmark reference.** Loyalty Juggernaut's GRAVTY publicly claims >100M transactions/hour (~30,000 TPS) at <50ms average latency — a useful high-end yardstick when scoring "real-time rules engine" claims.

### F2. Data Architecture Patterns

| Pattern | What It Means | Best Used When |
|---|---|---|
| Event sourcing | Every transaction stored as an immutable event log | Audit, dispute resolution, retroactive recalculation |
| CQRS | Separate write (transactions) and read (dashboards) systems | High-volume programs where read performance matters |
| CDC (Change Data Capture) | DB changes streamed to other systems automatically | Keep CRM/CDP/analytics in sync without batch ETL |
| Ledger model | Points as a bank account with double-entry accounting | Programs with financial-liability implications |

> This repo's own architecture (Postgres + Redis + Kafka, repository/service layering) maps to an event-driven + ledger approach: Kafka as the event stream, a points-transaction table as the ledger. Idempotency keys on earn/burn endpoints are the single most important correctness safeguard.

### F3. Integration Patterns

| Integration Type | How It Works | Common Use Cases |
|---|---|---|
| REST API | Request-response, real-time earn/burn | POS, checkout, member lookup |
| Webhooks | Platform pushes events to you | Trigger emails, update CRM, fraud alerts |
| Event stream (Kafka/SQS) | High-throughput event pipe | Millions of txns/day, real-time analytics |
| File import/export | Batch CSV/SFTP for legacy systems | ERP sync, offline POS reconciliation |
| SDK (mobile/web) | Pre-built code to embed loyalty UI | App loyalty tabs, balance display |

### F4. Security & Compliance Considerations

| Area | What to Look For | Red Flags |
|---|---|---|
| Points fraud prevention | Velocity checks, device fingerprinting, manual override | No controls or only post-hoc detection |
| Data privacy (GDPR/PDPA) | Right-to-erasure, consent management, data minimization | PII stored with no clear deletion process |
| Financial compliance | Breakage accounting, gift-card regs, B2B tax reporting | No audit trail; points treated as pure marketing |
| SSO / IAM | SAML/OAuth for admin, MFA for members | Shared admin passwords; no MFA |
| Pen testing / certs | SOC 2, ISO 27001, annual pen test | No certifications; vague security docs |

> Enterprise/AI-native vendors increasingly lead with certifications — e.g. GRAVTY cites ISO 27001/27018, SOC 1 & 2 Type II, GDPR/CCPA. Treat the absence of named certs at enterprise tier as a red flag.

---

## DOCUMENT G: Recommendations & Decision Framework

### G1. Build vs. Buy vs. Compose

| Scenario | Recommended Approach | Rationale |
|---|---|---|
| MVP, fast time-to-market, limited budget | Buy (SMB SaaS) | Smile.io / LoyaltyLion / Yotpo — live in days |
| Complex rules, multiple markets, in-house tech | Compose (API-first) | Talon.One or Voucherify + your own UI |
| Enterprise, multi-brand, coalition | Buy (Enterprise Suite) | Antavo, Annex Cloud, Loyalty Juggernaut — full partner ecosystem |
| Full ownership, high volume, data sovereignty | Build on open-source | Open Loyalty as engine + custom frontend |
| AI-first, personalization-heavy, digital-native | Emerging platforms | Hologrow, LJI — evaluate carefully (maturity varies) |

> **For this team (go-loyalty-point, a Go/Postgres/Redis/Kafka service):** you are already on the *build* path with a composable foundation. The pragmatic question is *what to build vs. integrate* — e.g. build the core ledger + rules in-house (you have the stack), but consider integrating a specialized voucher engine (Voucherify) or fulfillment provider (Tango for B2B) rather than rebuilding those. **(AI synthesis — recommendation, not a verified benchmark.)**

### G2. Platform Selection Scorecard

Weight and score each criterion (1–5) for your context:

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

**Technical**
- 99th-percentile API response time at peak load?
- How do you handle double-spend / race conditions in redemptions?
- Can we access raw event-level data for our own warehouse?
- Disaster recovery / RTO?

**Commercial**
- What happens to our data if we churn?
- Minimum contract term?
- How are price increases handled at renewal?

**Product**
- Roadmap for AI/personalization features?
- How many engineers do customers typically need to integrate?
- Three reference customers in our industry?

---

## DOCUMENT H: Open Research Questions

- [ ] What loyalty mechanics resonate most with Indonesian / Southeast Asian consumers specifically?
- [ ] What is the realistic earn/burn liability cost for a new program in our vertical?
- [ ] Which platforms have succeeded with programs similar to our use case (EV / mobility)?
- [ ] How do our target members prefer to interact: app, WhatsApp, web, or in-person?
- [ ] Typical activation rate (% of enrollees who earn at least once) in comparable programs?
- [ ] What fraud patterns should we design against first?

---

## DOCUMENT I: Glossary of Loyalty Terms

| Term | Plain-Language Definition |
|---|---|
| Earn rate | Points per currency unit spent (e.g. 1 pt per $1) |
| Burn rate | Point cost of a reward (e.g. 500 pts = $5 off) |
| Breakage | Earned-but-never-redeemed points — a benefit for the operator |
| Soft currency | Points/credits for status only, no redemption value |
| Hard currency | Points with real monetary value → financial liability |
| Tier qualification window | Period to accumulate enough spend/points to reach/keep a tier |
| Coalition program | Program shared across multiple brands; earn and spend at all partners |
| Headless loyalty | Loyalty logic decoupled from the frontend — your UI, their engine |
| Idempotency key | Ensures a transaction can be retried safely without double-awarding |
| Event-driven | Reacts to real-time signals, not nightly batches |
| Voucher stacking | Applying multiple codes/offers at once; controlled by stacking rules |
| Redemption friction | Any barrier to using earned rewards |
| CLV | Total revenue expected from one customer over the relationship |
| Churn prediction | ML model identifying members at risk of going inactive |
| MACH | Microservices, API-first, Cloud-native, Headless |

---

## Sources

- [Talon.One — Pricing](https://www.talon.one/pricing) · [G2](https://www.g2.com/products/talon-one/pricing) · [Capterra](https://www.capterra.com/p/159778/Talon-One/pricing/)
- [Voucherify — Pricing](https://www.voucherify.io/pricing) · [GetApp](https://www.getapp.com/website-ecommerce-software/a/voucherify/pricing/)
- [Smile.io — Pricing](https://smile.io/pricing) · [Flowium plan breakdown](https://flowium.com/blog/understanding-the-smile-io-pricing-plans/)
- [Open Loyalty](https://www.openloyalty.io/) · [Voucherify open-source accelerator](https://www.voucherify.io/open-source-composable-loyalty-accelerator)
- [Antavo — AI Loyalty Cloud](https://antavo.com/) · [Gartner Peer Insights](https://www.gartner.com/reviews/product/antavo-ai-loyalty-cloud) · [Antavo in Forrester Loyalty Platforms Landscape Q3 2025](https://antavo.com/news/antavo-named-in-the-loyalty-platforms-landscape-2025/)
- [Hologrow](https://hologrow.ai/) · [AI loyalty program generator](https://hologrow.ai/loyalty-program-generator)
- [Loyalty Juggernaut (GRAVTY)](https://www.lji.io/) · [Forrester Wave Q4 2025 recognition](https://www.prnewswire.com/news-releases/loyalty-juggernaut-ushers-in-a-new-era-of-modern-loyalty-tech-recognized-as-a-strong-performer-in-the-forrester-wave-for-loyalty-platforms-q4-2025-302637698.html) · [AWS Marketplace](https://aws.amazon.com/marketplace/pp/prodview-7bxvqyb2njh6s)
- [PAR Punchh](https://punchh.com/) · [PAR restaurant loyalty](https://partech.com/solutions/guest-engagement-platform/restaurant-loyalty-software/) · [Wallet-based loyalty / Smart Passes](https://www.verdictfoodservice.com/news/par-technology-wallet-based-loyalty-tool-restaurants/)

*End of Report*
