# /design-architecture

Produces a Technical Design Document (TDD) by translating product requirements into system architecture, API contracts, data models, and a phased implementation plan.

## When Invoked

If no PRD or context is provided, ask:

```
I'll help design the technical architecture. Please provide:

1. **PRD or requirements:** Link to the PRD, or describe the feature to build
2. **Current system context:** What services/components currently exist?
3. **Technology constraints:** Any mandatory tech choices (language, framework, cloud)?
4. **Timeline:** Target delivery date or milestone?
5. **Team structure:** Who will implement this (frontend, backend, platform)?
```

## Architecture Design Process

### Step 1: Requirements Analysis

Read the PRD (or intake requirements). Identify:
- All functional requirements → map to technical components
- Non-functional requirements → translate to technical constraints
- Acceptance criteria → translate to testable technical assertions

Flag any requirements that are ambiguous or technically risky:
```
⚠️ Ambiguous requirement: "The system should be fast"
   → Need to define: Fast at what? Under what load? Measured how?
   → Proposed SLO: API p99 < 200ms at 100 req/s
   → Is this accurate?
```

### Step 2: Architecture Options

Present 2-3 architectural options for key decisions:

```
## Key Technical Decision: [Decision Name]

Option A: [Approach] — Simple, lower operational complexity
  + [Advantages]
  - [Disadvantages]
  Best when: [conditions]

Option B: [Approach] — More scalable, higher initial investment
  + [Advantages]
  - [Disadvantages]
  Best when: [conditions]

Recommendation: Option [X] because [rationale]
Agree? Or do you prefer Option [Y]?
```

Get explicit agreement on key decisions before writing the full TDD.

### Step 3: API Contract Design

Design all API endpoints before data model:
- RESTful resources, methods, paths
- Request/response schemas (JSON)
- Error responses
- Authentication requirements
- Pagination strategy

### Step 4: Data Model Design

Design database schema to support the API:
- New tables/columns
- Indexes
- Migration strategy (zero-downtime)
- Rollback plan

### Step 5: Write the TDD

Use the `tech-lead-architect` agent to write the full Technical Design Document.
Save to `.claude/thoughts/architecture/YYYY-MM-DD-[feature-name]-tdd.md`

### Step 6: Risk Review

Present identified risks with mitigations before finalizing:
```
## Technical Risks Identified

1. [Risk]: [Description]
   Likelihood: High/Med/Low | Impact: High/Med/Low
   Mitigation: [Specific action]

2. [Risk]: ...

Are these risks correctly characterized? Any additional risks I should document?
```

### Step 7: Implementation Phase Planning

Break work into phases that can be independently deployed:
- Phase 1: MVP — what delivers the core user value
- Phase 2: Hardening — observability, edge cases, performance
- Phase 3: (Optional) — enhancement features

Present a high-level work breakdown:
```
Estimated work breakdown:
  Backend: X days
  Frontend: Y days
  Infrastructure: Z days
  Testing: W days
  Total: ~N days

Does this scope seem right? Should any scope be cut from Phase 1?
```

## Important Guidelines

- API contracts must be reviewed by frontend/backend leads before finalizing
- All database migrations must be zero-downtime — document the migration strategy
- Technical risks must be explicit — surface unknowns, don't hide them
- Phase 1 must be independently testable and deployable
- After TDD is written, recommend creating implementation tasks: /plan-feature
