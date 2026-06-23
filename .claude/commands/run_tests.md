# /run-tests

Runs the appropriate test suite, analyzes failures, and provides actionable fix guidance. Orchestrates testing agents based on what tests are needed.

## When Invoked

If no scope is specified:

```
Which tests would you like to run?

1. All tests (unit + integration)
2. Unit tests only (fast, no external dependencies)
3. Integration tests (requires DB/services)
4. E2E tests (full stack, slowest)
5. Specific package/feature: [provide path or name]
6. Failed tests from last run
7. Benchmark tests

Or describe what you want to test and I'll select the right command.
```

## Test Execution Process

### Step 1: Determine Test Type

Based on the request:

**Unit tests:**
```bash
# Go
go test -race -short -coverprofile=coverage.out ./...

# Frontend
pnpm test --run
```

**Integration tests:**
```bash
# Go (requires TEST_DATABASE_URL set)
go test -race -run Integration ./...

# Or with docker-compose for dependencies
docker compose -f docker-compose.test.yaml up -d
go test -race ./...
docker compose -f docker-compose.test.yaml down
```

**E2E tests:**
```bash
npx playwright test
```

**Benchmarks:**
```bash
go test -bench=. -benchmem -benchtime=10s ./...
```

### Step 2: Run and Monitor

Execute the tests and capture:
- Total: passed / failed / skipped count
- Duration
- Coverage percentage
- Any flaky patterns (same test failing intermittently)

### Step 3: Failure Analysis

For any test failures, analyze and categorize:

```markdown
## Test Run Summary
Date: YYYY-MM-DD HH:MM
Command: `go test -race ./...`
Result: ❌ FAILED

### Failed Tests (3)

#### 1. TestOrderService_CreateOrder_WhenDBDown
File: internal/domain/order/service_test.go:45
Error:
  expected error "connection refused", got nil

Root cause: The test expects DB failure, but the mock is returning success.
The mock setup at line 38 returns nil error unconditionally.

Fix:
  Line 38: Change `mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil, nil)`
  To: `mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil, errors.New("connection refused"))`

#### 2. TestUserHandler_POST_ValidatesEmail
...

### Coverage Report
Package                           Coverage  Change
internal/domain/order             82.3%     +2.1%
internal/infrastructure/http      71.0%     -0.5% ⚠️
internal/domain/user              68.2%     No change

⚠️ Coverage dropped in internal/infrastructure/http — new code added without tests
```

### Step 4: Guidance

For each failure, provide:
1. Root cause explanation (not just what failed, but why)
2. Specific fix with file:line reference
3. Whether it's a test bug or a code bug

For coverage drops:
- Identify the new uncovered code
- Suggest which test cases to add
- Provide test case template

### Quick Commands Reference

```makefile
# Add to project Makefile

test:          ## Run all tests
	go test -race ./...

test-unit:     ## Run unit tests (fast, no deps)
	go test -race -short ./...

test-cover:    ## Run tests with coverage report
	go test -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

test-watch:    ## Watch mode (requires gotestsum)
	gotestsum --watch --format=testname -- -short ./...

test-race:     ## Run with race detector (CI-grade)
	go test -race -count=3 ./...

bench:         ## Run benchmarks
	go test -bench=. -benchmem ./...
```

## Important Guidelines

- Always run with `-race` flag — race conditions are silent without it
- Integration tests require the test environment to be running — check docker status first
- Coverage below 70% overall or below 85% for domain packages is a concern — report it
- Never modify tests to make them pass by removing assertions — fix the code or test logic
- If tests are flaky (pass sometimes, fail sometimes), flag it — don't just re-run until it passes
- Benchmark results need comparison against a baseline to be meaningful
