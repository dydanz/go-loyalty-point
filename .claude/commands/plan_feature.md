# /plan-feature

Creates a detailed implementation plan for a feature or task. Researches the codebase to understand impact, designs the approach, and produces an actionable plan with success criteria.

## When Invoked

If no feature is described, ask:

```
I'll help you create a detailed implementation plan. Please provide:

1. **Feature or task description** — what needs to be built or changed
2. **Context** — any relevant background, requirements, or constraints
3. **Scope** — is there a PRD or technical design to reference?
4. **Timeline** — any deadline or target milestone?

The more context you provide, the better the plan.
```

## Planning Process

### Phase 1: Research (Parallel)

Spawn multiple agents simultaneously:
- **Codebase locator** — find files most relevant to this feature
- **Pattern analyzer** — understand existing patterns to follow or extend
- **Dependency mapper** — identify what this feature depends on and what depends on it
- **Test pattern finder** — find existing tests for similar features as templates

Synthesize findings before proceeding. Ask clarifying questions if needed.

### Phase 2: Design

Based on research findings:
1. Identify the layer where each piece of work lives (frontend / backend / infra / tests)
2. Design the implementation approach (extend existing vs new component)
3. Identify risks and unknowns — surface them explicitly
4. Sequence work: what must be done first?

Present the approach and ask for feedback:

```
Based on my analysis, here's my proposed approach:

**Approach:** [summary of the plan]
**Key files affected:** [file:line list]
**Estimated complexity:** Simple / Medium / Complex
**Main risk:** [identified risk]

Does this align with your expectations? Any concerns before I write the full plan?
```

### Phase 3: Write the Plan

After approval, write the full implementation plan to `.claude/thoughts/plans/YYYY-MM-DD-[feature-name].md`:

```markdown
# Plan: [Feature Name]
Date: YYYY-MM-DD
Status: Approved

## Goal
[What done looks like]

## Approach
[Technical approach with rationale]

## Implementation Steps

### Phase 1: [Backend / Foundation]
- [ ] Step 1: [Specific task] — `internal/path/file.go`
  - Implementation details
  - What to watch out for
- [ ] Step 2: [Specific task]

### Phase 2: [Frontend / Integration]
- [ ] Step 3: [Specific task] — `src/features/...`
- [ ] Step 4: [Specific task]

### Phase 3: [Tests & Documentation]
- [ ] Step 5: [Unit tests for X]
- [ ] Step 6: [Integration tests for Y]

## Success Criteria

### Automated (can be verified by running commands)
- [ ] `go test ./...` passes
- [ ] `npm run test` passes
- [ ] `go vet ./...` passes
- [ ] No linting errors

### Manual (requires human verification)
- [ ] Feature works end-to-end in local environment
- [ ] Edge case: [specific scenario] behaves correctly
- [ ] Performance: [specific operation] responds in < Xms

## Out of Scope
- [What this plan does NOT cover]

## Open Questions
- [ ] [Question that needs resolution before starting]
```

## Important Guidelines

- Never start implementing until the plan is approved by the user
- Plans must have no open questions before they are marked "Approved"
- Each step must reference a specific file path — no vague steps
- Success criteria must distinguish automated (CI) from manual verification
- Flag technical risks explicitly — do not assume they're understood
- If the feature is large, propose phasing: deliver value incrementally
