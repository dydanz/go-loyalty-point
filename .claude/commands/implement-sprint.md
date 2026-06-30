---
description: Implement the go-loyalty-point TRD sprint by sprint as a principal fullstack engineer (Go + ReactJS) — one task per run, test-first, draft PR. Routine-friendly.
model: claude-sonnet-4-6
argument-hint: "[sprint N | task N.M | --daily] (optional; defaults to next unstarted task)"
allowed-tools: Read, Write, Edit, Grep, Glob, Bash, Skill, Agent, TodoWrite
---

# /implement-sprint

Act as a **principal fullstack software engineer** (Go backend + ReactJS frontend) and
implement the next piece of the TRD, **one task per run**, test-first, following existing
patterns. Runs on **Sonnet**. Designed to be invoked manually or by a daily Claude Routine.

**First action:** invoke the `implement-sprint` skill (Skill tool) and follow it exactly.
It is the full playbook (task selection, Go/React rules, TDD, Swagger, draft-PR wrap-up).

## Argument

`$ARGUMENTS`

- empty → pick the next dependency-met, signal-unmet task from `docs/trd/<phase>/sprint-*.md`.
- `sprint N` → next unstarted task within Sprint N.
- `task N.M` → that specific task.
- `--daily` → **Daily-Routine mode** (skill §8): exactly one task, fully local + branch +
  draft PR, clean tree, standup-style log. Never merge, never push `master`, never deploy.

## Guardrails (from CLAUDE.md + agents/sdlc.md)

- Never push to `master`; branch (`fix/…` low, `feature/…` medium+) and open a **draft** PR.
- Never auto-merge or auto-deploy. GitHub writes need confirmation in interactive mode; in
  `--daily` mode, branch + draft PR are the only standing authorization.
- Always `go test -race ./...` before claiming done; FE tests if frontend touched.
- `swag init -g main.go` after any API change. Migrations append-only. No global state.
  Parameterized SQL. zerolog structured fields. Tenant isolation enforced + tested.
- One task per run. Stop and report on any unresolved dependency or open question — don't guess.

## Output

End with a short log: **task id · files touched · test result (with evidence) · PR link ·
next task · blockers.** Suitable as a daily standup entry.

## Wiring into a daily Claude Routine

Schedule this command (via `/schedule` or a routine) to run daily, e.g.:

```
/implement-sprint --daily
```

The routine advances the TRD by one tested, reviewed, draft-PR'd task per day. Review and
merge the draft PR yourself — the routine never merges.
