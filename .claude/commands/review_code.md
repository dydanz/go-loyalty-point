# /review-code

Performs a comprehensive code review covering correctness, security, performance, maintainability, and adherence to project patterns. Can review a specific file, a PR diff, or a feature branch.

## When Invoked

If no target is specified:

```
What would you like me to review?

1. Current git diff (uncommitted changes)
2. Specific file(s): [provide path(s)]
3. Feature branch: [provide branch name]
4. Recent commits: last N commits

You can also specify a focus area:
  - Security (auth, injection, data exposure)
  - Performance (N+1 queries, inefficient algorithms, memory)
  - Correctness (logic errors, edge cases, error handling)
  - Patterns (adherence to project conventions)
  - Tests (coverage, quality, missing cases)
  - All of the above (default)
```

## Review Process

### Step 1: Understand Context

Before reviewing:
- Read the git diff or target files
- Understand the intent of the change (from commit messages, PR description, or asking the user)
- Note what domain this change is in (frontend / backend / infrastructure)

### Step 2: Multi-Dimension Review

Evaluate each change against all dimensions:

#### Correctness
- [ ] Logic is correct for the happy path
- [ ] Edge cases are handled (empty inputs, nil, zero values, boundary conditions)
- [ ] Error cases return appropriate errors (not silent failures)
- [ ] Concurrent access is safe (no data races)
- [ ] Resource cleanup happens (defer close, cancel context)

#### Security
- [ ] No SQL injection (parameterized queries only)
- [ ] No command injection (no user input in Bash/exec)
- [ ] Authentication is enforced on protected endpoints
- [ ] Authorization checks prevent unauthorized access
- [ ] No sensitive data in logs, errors, or responses
- [ ] Input is validated and sanitized at boundaries

#### Performance
- [ ] No N+1 queries (check loops that call DB/API)
- [ ] Database queries use indexes (check EXPLAIN plan for new queries)
- [ ] No unbounded memory growth (maps/slices that grow without limit)
- [ ] No unnecessary allocations in hot paths
- [ ] Caching is used appropriately for expensive repeated operations

#### Maintainability
- [ ] Function/method does one thing
- [ ] Naming is clear and consistent with project conventions
- [ ] No duplication of existing logic elsewhere in the codebase
- [ ] Complex logic has clarifying comments (not obvious comments)
- [ ] No dead code or commented-out code

#### Tests
- [ ] New logic has corresponding tests
- [ ] Tests cover error/edge cases, not just happy path
- [ ] Test names describe behavior (not implementation)
- [ ] Mocks are appropriate (not mocking what should be integration-tested)

#### Project Patterns
- [ ] Follows established error handling patterns in this codebase
- [ ] Follows logging patterns (structured, context-aware)
- [ ] Follows the architectural layer boundaries
- [ ] Import paths follow project conventions

### Step 3: Produce the Review

Format findings by severity:

```markdown
## Code Review: [Target]
Date: YYYY-MM-DD

### Summary
[1-2 sentence overall assessment]

### 🔴 Critical (must fix before merge)
- `path/to/file.go:45` — SQL query uses string interpolation → SQL injection risk
  **Fix:** Use parameterized query: `db.Query("... WHERE id = $1", id)`

### 🟡 Important (should fix)
- `path/to/file.go:78` — Error from `cache.Set()` is silently ignored
  **Fix:** Log the error: `if err != nil { logger.Warn("cache set failed", "error", err) }`

### 🔵 Suggestions (nice to have)
- `path/to/file.go:102` — This logic duplicates `pkg/utils/validate.go:34`
  **Suggestion:** Extract to shared function or reuse existing

### ✅ Looks Good
- Error handling pattern matches project conventions
- Test coverage includes edge cases
- Concurrency is handled correctly with mutex
```

## Important Guidelines

- Reference every finding with exact file:line — never make vague claims
- Distinguish Critical (security/correctness) from Suggestions (style/improvement)
- Explain WHY something is a problem, not just that it is
- Provide specific fix suggestions, not just "fix this"
- Acknowledge what is done well — reviews should be constructive, not just critical
- If a pattern is used differently than expected, check if it's intentional before flagging
