# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

---

## Project Overview

**go-loyalty-point** is a Go HTTP API service for managing loyalty points — earning, redeeming, and tracking points for users across transactions. Built with Gin, PostgreSQL, Redis, Kafka, and zerolog.

Stack: `Go 1.23` · `Gin` · `PostgreSQL` · `Redis` · `Kafka` · `zerolog` · `Swagger` · `Docker`

---

## Running the Project

```bash
# Build and run
go build -o server ./...
./server

# Run without building
go run ./main.go

# Run with docker compose
docker compose up --build

# Run tests
go test -race ./...

# Generate Swagger docs
swag init -g main.go
```

Environment variables (never commit secrets):
```bash
export DATABASE_URL=...
export REDIS_URL=...
export KAFKA_BROKERS=...
```

---

## Architecture

```
HTTP Request (Gin)
      │
      ▼
Middleware (auth, logging, CORS, recovery)
      │
      ▼
Handler (pkg/handler/) — thin HTTP layer, validates input
      │
      ▼
Service (pkg/service/) — business logic, orchestrates repos
      │
      ▼
Repository (pkg/repository/) — DB queries via lib/pq
      │
      ▼
PostgreSQL / Redis / Kafka
```

**Package layout:**
```
pkg/
  handler/        # HTTP handlers (Gin) — thin layer, delegates to service
  service/        # Business logic
  repository/     # DB/cache access
  domain/         # Domain structs, value objects, errors
  middleware/      # Auth, logging, CORS
  router/         # Route registration
  config/         # Config struct + env loader
  database/       # DB connection setup
  channel/        # Kafka producers/consumers
  kafka/          # Kafka helpers
  logging/        # zerolog setup
  mocks/          # Testify mocks
  migrations/     # SQL migration files
  bootstrap/      # App bootstrap / wiring
  util/           # Shared helpers
  docs/           # Swagger generated docs
```

---

## Key Design Patterns

### Error Handling
Typed domain errors in `pkg/domain/`. Infrastructure errors are wrapped with context before returning up the stack. HTTP handlers map domain errors to HTTP status codes.

### Logging
`zerolog` via `pkg/logging/`. Always use structured fields — never `fmt.Sprintf` log messages. Pass logger via context or dependency injection.

### Config
Loaded from env vars at startup. Fail fast on missing required config. Never store secrets in code or config files.

### Kafka
Producers and consumers in `pkg/channel/` and `pkg/kafka/`. Events are the integration mechanism between bounded contexts.

### Database Migrations
SQL files in `pkg/migrations/`. Use `golang-migrate`. Never modify a migration that has already been applied — always add a new one.

---

## Hard Rules (Non-Negotiable)

- **Never push to main** — always open PRs
- **Never auto-deploy to production** — require explicit approval
- **Secrets via env vars only** — no secrets in code or committed files
- **Always run `go test -race ./...`** before marking work done
- **Migration files are append-only** — never edit applied migrations
- **No global state** — pass dependencies via constructor injection

---

## What to Avoid

- Global variables for dependencies (DB connections, loggers)
- Skipping error handling — every error must be logged or returned
- String-interpolated SQL — always use parameterized queries
- Direct `fmt.Println` / `log.Print` — use zerolog
- Business logic in handlers — handlers are thin HTTP adapters only

---

## SDLC

This project follows a lightweight multi-phase SDLC. Full agent: `.claude/agents/sdlc.md`.

**Three rules always active:**
1. **Draft, don't auto-execute.** Propose every GitHub action (issue, PR, comment). Wait for explicit operator confirmation.
2. **Event-driven.** Act on invocation or hook events. Never poll.
3. **Proportional.** Read `class:low/medium/high/hotfix` from issue labels. Scale ceremony to that class.

| Class | Spec | Gate | Test | Branch |
|---|---|---|---|---|
| `class:low` | None | Self-approval | CI | `fix/<id>-slug` |
| `class:medium` | Spec in issue body | Self `/approve-spec` | `go test ./...` | `feature/<id>-slug` |
| `class:high` | TRD in issue / doc | Self-approval + 2nd review | Full suite + manual | `feature/<id>-slug` |
| `class:hotfix` | Skip; post-merge ≤24h | Self-approval + smoke | Smoke | `hotfix/<id>-slug` |

**PR creation skill:** `.claude/skills/github-pr/SKILL.md` — invoke via `/github-pr` or say "create PR".

---

## graphify

If `graphify-out/graph.json` exists, use it for codebase questions:
- `graphify query "<question>"` — scoped subgraph
- `graphify path "<A>" "<B>"` — relationships
- `graphify explain "<concept>"` — focused concept
- Read `graphify-out/GRAPH_REPORT.md` only for broad architecture review.
- After modifying code, run `graphify update .` to keep the graph current.
