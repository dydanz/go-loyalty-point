---
name: routing
description: Model routing rules for the AKB48 multi-agent system. Read this before deciding which model tier to use for any task.
---

# Model Routing Rules

## Tier Definitions

| Tier | Model | Use When |
|------|-------|----------|
| **Haiku** | `claude-haiku-4-5` | Fast, cheap, deterministic tasks with no judgment required |
| **Sonnet** | `claude-sonnet-4-6` | Standard implementation, moderate reasoning, code review |
| **Opus** | `claude-opus-4-7` | Architecture decisions, ambiguous requirements, multi-step planning |

**Decision rule:** If in doubt, start one tier lower. Escalate when you detect ambiguity, conflicting requirements, or need to make a judgment call with lasting consequences.

---

## Haiku Tasks

Low-cognitive tasks — no judgment, no design decisions, deterministic output.

### Git & GitHub Operations
- `git commit` — stage files and write commit message
- `git push` — push branch to remote
- Create a GitHub PR — `gh pr create` with title and body
- Add PR labels, assignees, or reviewers
- Close or merge a PR (when explicitly instructed)

### File Operations
- Read a file or directory listing
- Search for a string or symbol in files (`grep`, `glob`)
- Rename a file or directory
- Move a file or directory
- Delete a file (when explicitly instructed)
- Copy a file

### Classification & Triage
- Classify a user message into a category
- Extract structured data from text (fact extraction, JSON parsing)
- Identify which skill file matches a user intent
- Route a task to the correct agent

### Formatting & Boilerplate
- Format code to match style conventions
- Generate boilerplate from a template (struct stubs, test scaffolding)
- Reformat a config file
- Sort imports
- Convert between data formats (JSON ↔ TOML ↔ YAML)

### Simple Lookups
- Look up a type, function, or constant in the codebase
- Report which files changed in a diff
- Summarize a changelog or git log
- Check if a dependency is present in go.mod

---

## Sonnet Tasks

Moderate reasoning — implementation work with clear requirements.

### Code Implementation
- Implement a function or method with a defined interface
- Write tests for an existing function
- Fix a bug with a clear root cause
- Refactor code to match an established pattern
- Add error handling to an existing function
- Wire up a new adapter or handler using existing patterns

### Code Review
- Review a PR diff for correctness, security, and pattern adherence
- Identify missing error handling, data races, or resource leaks
- Check test coverage gaps
- Flag deviations from project conventions

### Moderate Reasoning
- Explain how a piece of code works
- Trace a request through the call stack
- Debug a test failure with a known error message
- Evaluate two concrete implementation options (when requirements are clear)

### Test Writing
- Write unit tests for a new function
- Write integration test scaffolding
- Generate table-driven tests for edge cases

---

## Opus Tasks

High-stakes judgment — design decisions with lasting consequences or genuinely ambiguous inputs.

### Architecture Decisions
- Design a new subsystem from scratch (new PRD, new integration)
- Choose between fundamentally different approaches (e.g., MCP vs REST for GBrain)
- Define interface contracts between major components
- Decide on session management or caching strategy
- Evaluate a proposal that touches multiple PRDs

### Ambiguous or Conflicting Requirements
- Requirements contradict each other or are underspecified
- Multiple stakeholders have conflicting goals
- The task requires inferring unstated constraints
- PRD is missing acceptance criteria

### Multi-Step Planning
- Break a large feature into ordered tickets
- Sequence implementation phases with dependency analysis
- Define a migration plan with rollback strategy
- Draft a technical spec that others will implement from

### Conflict Resolution
- Resolve a design disagreement between agents
- Adjudicate between two valid architectural patterns
- Decide which PRD interpretation to follow when ambiguous

---

## Escalation Protocol

An agent operating at Haiku or Sonnet tier MUST escalate to the next tier when:

1. The task requires a judgment call not covered by existing patterns
2. Requirements are ambiguous and cannot be resolved by reading the PRD
3. The implementation would touch more than 3 files in ways that affect interfaces
4. Two valid approaches exist with significant trade-offs

Escalation means: stop, describe the ambiguity, and request Opus-tier review before proceeding.

---

## AKB48-Specific Routing

| Task | Tier | Reason |
|------|------|--------|
| `git commit` / `git push` | Haiku | Deterministic, no judgment |
| `gh pr create` | Haiku | Templated, no judgment |
| Read / search / move files | Haiku | File I/O, no reasoning |
| Fact extraction for GBrain | Haiku | Per CLAUDE.md design intent |
| Write a new tool handler | Sonnet | Implementation with clear interface |
| Code review (PR or diff) | Sonnet | Pattern matching + correctness check |
| Write tests for a handler | Sonnet | Clear input/output contract |
| Fix a known bug | Sonnet | Bounded scope |
| Design a new PRD component | Opus | Architecture, interface contracts |
| Session compaction strategy | Opus | Memory policy, lasting consequences |
| GBrain MCP integration design | Opus | Multi-system interface |
| Skill resolver algorithm | Opus | Judgment-heavy, affects all skills |
