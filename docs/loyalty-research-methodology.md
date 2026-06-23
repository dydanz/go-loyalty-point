# Research Approach & Workflow Guide
## How to Conduct Loyalty Platform Research as a Non-Technical PM

> Companion to: `loyalty-platform-research-prompt.md`
> Audience: Product Manager, Product Researcher, Business Analyst
> Reading time: ~10 minutes

---

## The Core Problem This Approach Solves

Loyalty platform research is deceptively hard for non-technical PMs because:

1. **The feature names are deceptive** — "points engine" means very different things at Talon.One vs. Smile.io
2. **Vendor marketing is uniform** — every platform claims "AI-powered personalization" and "enterprise-grade"
3. **The real complexity is in the edge cases** — what happens during a return? a fraud event? a timezone mismatch? Demos never show this
4. **The technical and product decisions are entangled** — you can't pick features without understanding architecture limits

This guide gives you a structured 4-phase research process that produces trustworthy, actionable documents — even if you can't read code.

---

## Phase 1 — Landscape Mapping (Days 1–3)

**Goal**: Know who the players are, what category they play in, and which ones are worth investigating.

### Step 1.1 — Build your vendor longlist

Start with these sources:
- **G2.com** → search "loyalty management software" → filter by category and rating
- **Capterra** → same search
- **Gartner Peer Insights** → "Loyalty Marketing" category
- **Forrester / IDC reports** (if you have access) → check the Wave or MarketScape for loyalty platforms
- **ProductHunt** → search "loyalty rewards" → sort by newest (catches emerging platforms)
- **LinkedIn** → search "loyalty platform" in company pages, filter by industry

**Output**: A spreadsheet with 20–30 platform names, their category, and one-line description.

### Step 1.2 — Score and narrow to 6–10

Score each platform on:
- Relevance to your vertical / use case (1–5)
- Market presence / customer count (1–5)
- Evidence of innovation / recency (1–5)

Drop anything below 10/15 total. You should have 6–10 platforms to research deeply.

**Prompt to use at this stage**:
> "I have this list of loyalty platforms: [paste list]. Based on publicly available information, can you help me categorize them by architecture type (traditional vs. composable vs. SMB SaaS vs. vertical-specific) and identify which ones are most commonly mentioned in enterprise buying decisions?"

---

## Phase 2 — Deep Platform Research (Days 3–10)

**Goal**: For each shortlisted platform, extract a standardized profile using the Research Prompt (Document B through F).

### Step 2.1 — Run the AI research prompt per platform

For each platform, use this prompt structure:

```
Platform: [NAME]
Website: [URL]
Category: [from your Phase 1 mapping]

Using publicly available documentation, case studies, and review sites, complete the following profile:

1. Core feature list (what the platform actually does, with evidence)
2. Pricing model and entry price
3. Integration ecosystem (CRM, POS, e-commerce, CDP, email)
4. Architecture approach (monolith / MACH / headless)
5. Best-fit use cases (who are their reference customers?)
6. Verified limitations (from user reviews, not vendor copy)
7. Notable UX patterns (member-facing and admin-facing)

Flag anything unverified with ❓.
```

### Step 2.2 — Cross-reference with primary sources

After AI synthesis, validate with:

| Source Type | What to Verify | Where to Find It |
|---|---|---|
| Vendor docs | Feature names and limits | docs.[vendor].com |
| G2 reviews | Real user pain points | g2.com/products/[vendor]/reviews |
| Case studies | Actual customer outcomes | vendor website → "Customers" |
| LinkedIn posts | Recent product updates | LinkedIn company page |
| YouTube demos | Actual UI and admin console | Search "[vendor] demo 2024/2025" |
| Pricing page | Current pricing tiers | vendor website → "Pricing" |

**Rule**: Anything marked ✅ in your research must have a source link. ❓ means AI-synthesized with no primary source.

### Step 2.3 — Conduct vendor demos strategically

Request demos from your top 4–5 platforms. Use this demo checklist:

**Before the demo**:
- [ ] Send your use case in writing (they'll tailor the demo)
- [ ] Ask for a sandbox environment or trial access
- [ ] Prepare 5 scenario-based questions (not feature questions)

**During the demo — ask to see**:
- [ ] The rule builder: "Show me how I'd set up a 'buy 3, earn double points' campaign"
- [ ] Edge case handling: "What happens when a member returns an item they earned points on?"
- [ ] Admin fraud tools: "Show me what I'd see if a member abused a referral code"
- [ ] Data export: "How would I get all my member transaction data into our data warehouse?"
- [ ] API documentation: "Can I look at your API reference for point award endpoints?"

**After the demo**:
- [ ] Ask for a reference customer in your industry
- [ ] Ask for a written response to your 5 technical questions

---

## Phase 3 — Synthesis & Documentation (Days 10–14)

**Goal**: Turn raw research into PM-team-ready documents.

### Step 3.1 — Use the Research Prompt to structure output

The `loyalty-platform-research-prompt.md` file produces 9 documents (A through I). Treat each as a standalone shareable artifact:

| Document | Primary Audience | Format |
|---|---|---|
| A: Market Landscape | Executive / Stakeholder | Slide or 1-pager |
| B: Feature Catalogue | PM Team + Design | Reference wiki page |
| C: Comparison Matrix | PM Team | Spreadsheet |
| D: UX/UI Analysis | Design Team | Design brief input |
| E: Use Case Stories | PM + Engineering | User story backlog |
| F: Technical Capabilities | Engineering Lead | Technical spec input |
| G: Recommendations | CPO / Decision-maker | Decision memo |
| H: Open Questions | Research team | Next sprint input |
| I: Glossary | All teams | Shared reference |

### Step 3.2 — AI-assisted document drafting

For each document, use this approach:

1. **Feed your raw research** into the AI with the relevant section of the prompt
2. **Ask for a first draft** with uncertainty clearly flagged
3. **You edit for accuracy** based on what you verified in Phase 2
4. **Add source links** for every claim you want stakeholders to trust

**Prompt for synthesis**:
```
Here is my raw research on [platform A], [platform B], and [platform C]:
[paste your notes]

Using Document C format from the Loyalty Research Prompt, create a feature comparison matrix.
- Use ✅ only where I have a verified source
- Use ❓ where the feature is claimed by the vendor but I haven't verified it
- Use ❌ where a review or limitation clearly confirms absence
- Keep feature descriptions to one plain sentence
```

### Step 3.3 — Internal review before publishing

Run each document through this checklist:

- [ ] Every ✅ has a source link or note
- [ ] No vendor marketing language appears without attribution
- [ ] Pricing figures are dated (prices change)
- [ ] Technical explanations have been reviewed by one engineer
- [ ] Use case stories have been validated against at least one real customer example

---

## Phase 4 — Stakeholder Communication (Days 14–16)

**Goal**: Deliver findings in formats different audiences can actually use.

### Recommended deliverable stack

| Deliverable | Audience | Format | Length |
|---|---|---|---|
| Executive summary | CEO / CPO | Slide (3–5 slides) | 5 min read |
| Market landscape brief | Full PM team | PDF or Notion page | 10 min read |
| Feature comparison matrix | PM + Engineering | Spreadsheet | Reference |
| UX pattern library | Design team | Figma or PDF | Reference |
| Platform recommendation memo | Decision-maker | 1-pager | 2 min read |
| Full research archive | Internal wiki | Notion / Confluence | Reference |

### The 1-page recommendation memo structure

```
SUBJECT: Loyalty Platform Recommendation

RECOMMENDATION: [Platform X] for [reason in one sentence]

CONTEXT: We evaluated [N] platforms over [timeframe].

TOP 3 FINDINGS:
1. [Finding with evidence]
2. [Finding with evidence]
3. [Finding with evidence]

ALTERNATIVES CONSIDERED:
- [Platform Y]: Stronger at [X], weaker at [Y], ruled out because [Z]
- [Platform Z]: Considered for [reason], excluded because [reason]

NEXT STEPS:
- [ ] [Action] by [date] — Owner: [name]
- [ ] [Action] by [date] — Owner: [name]

OPEN RISKS:
- [Risk]: [Mitigation approach]
```

---

## Prompt Sequencing for AI-Assisted Research

If you're using Claude or another AI assistant throughout this process, here is the recommended sequence:

| Step | Prompt Goal | What to Feed In | What You Get Back |
|---|---|---|---|
| 1 | Category mapping | Your vendor longlist | Categorized table with rationale |
| 2 | Platform profile | Platform name + website | Structured profile with uncertainty flags |
| 3 | Feature comparison | Multiple platform profiles | Comparison matrix |
| 4 | UX synthesis | Demo notes + screenshots | UX pattern analysis |
| 5 | Use case drafting | Your business context | Tailored user stories |
| 6 | Recommendation | Your priorities + matrix | Scored recommendation |
| 7 | Exec summary | Full research output | 3-slide narrative |

**Always tell the AI what you already know** — "I've verified X from vendor docs, I need you to fill in Y from public sources, and flag Z as unverified."

---

## Common Research Mistakes to Avoid

| Mistake | Why It Happens | How to Avoid |
|---|---|---|
| Trusting vendor feature lists | Vendor marketing is always complete | Verify with G2 reviews and demo scenarios |
| Comparing features out of context | "Gamification" means different things to Punchh vs. Talon.One | Always define the feature before comparing |
| Picking by price alone | Cheapest platform often costs most in integration effort | Total cost = license + integration + maintenance |
| Skipping the admin UX | PMs focus on member experience, forget ops complexity | Always get a demo of the admin/merchant console |
| Ignoring data ownership terms | Your member data may be locked in the vendor's system | Read the data portability and exit clause before signing |
| Over-indexing on AI features | "AI personalization" is often just basic segmentation | Ask: "Show me the AI output for a real member, not a demo account" |
| Research without a use case | Generic research produces generic insights | Start with 3 specific scenarios you need the platform to handle |

---

## Time Budget (Suggested)

| Phase | Duration | Hours/Week |
|---|---|---|
| Phase 1: Landscape mapping | 3 days | ~6 hrs |
| Phase 2: Deep platform research | 7 days | ~12 hrs |
| Phase 3: Synthesis & docs | 4 days | ~10 hrs |
| Phase 4: Stakeholder delivery | 2 days | ~6 hrs |
| **Total** | **~16 days** | **~34 hrs** |

This is realistic for a solo PM researcher. With a team of 2, compress Phase 2 by splitting platforms.

---

*This guide should be used alongside `loyalty-platform-research-prompt.md` which provides the detailed output templates for each document type.*