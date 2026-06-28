---
name: github-issue
description: >
  Register tasks from docs/trd/ as GitHub issues for go-loyalty-point, following
  scrum/kanban convention: Epic (sprint) → User Story (task), with estimation
  (high/med/low + story points), tagging (type/feature/fix/chore, component, phase,
  sprint), and GitHub Project board fields. Use when the user says "register tasks",
  "create issues from TRD", "file sprint N as issues", "open issue", or "/github-issue".
---

# GitHub Issue Management — go-loyalty-point

Turns the per-sprint tech dev plans in `docs/trd/phase-*/sprint-*.md` into a tracked
backlog of GitHub issues. Pairs with the `github-pr` skill (a PR `Closes #N`) and the
`sdlc` agent (change-class flow).

## 0. Prime Directive — Draft, don't auto-execute

Adopted from `.claude/agents/sdlc.md`. **Never create, edit, or close a GitHub issue
without explicit operator confirmation.** Parse the source, draft the issue(s), show
them, then ask: *"Create these N issues, or revise?"* Wait for the word. Solo-operator
project (Dandi) — you draft, Dandi decides.

---

## 1. Source → Issue Mapping (scrum/kanban)

Source of truth = `docs/trd/<phase>/sprint-N.md`. Structure already matches agile:

| TRD source element | GitHub object | Convention |
|---|---|---|
| `sprint-N.md` (whole file) | **Epic** issue | one per sprint; `type/epic` |
| `## Task N.M — Title` | **User Story** issue | child of the epic; `type/*` by nature |
| Sprint folder `phase-mvp/` | **Milestone** | `MVP`, `Phase 1`… maps to release |
| `Type:` (backend/frontend/fullstack) | `component/*` label(s) | |
| `Effort:` S/M/L | `size/*` + Estimate (SP) + Priority | §3 |
| `Progress signal:` | Acceptance criterion #1 | testable, observable |
| `Design` / `Implementation` / `Testing` | issue body sections | mandatory |
| `Swagger:` line | Acceptance criterion (API registered) | |

**Hierarchy is via task list + label, not GitHub sub-issues** (keep it portable):
the Epic issue body holds a checklist linking each story (`- [ ] #NN Task 1.1 …`);
each story body has `Epic: #<epic-number>`.

Kanban columns (Project board Status field): `Backlog → Ready → In Progress → In Review → Done`, plus `Blocked`. New issues land in **Backlog**.

---

## 2. Label Taxonomy

One-time bootstrap (see §7). Apply on every issue.

### Type (exactly one)
| Label | When |
|---|---|
| `type/epic` | Sprint container issue |
| `type/feature` | New capability / user-facing behavior |
| `type/fix` | Corrects a defect (e.g. C1–C4, S-findings) |
| `type/chore` | Build/CI/infra/tooling, no product behavior |
| `type/security` | P0/P1 security defect or hardening |
| `type/test` | Test-coverage-only work |
| `type/docs` | Docs/Swagger-only |

### Change class (exactly one — mirrors sdlc + github-pr)
`class:low` · `class:medium` · `class:high` · `class:hotfix`
Default per effort: S→low/medium, M→medium, L→medium/high. Security defects ≥ `class:high`. If unsure, go higher.

### Component (all that apply — go-loyalty-point pkg layout)
`component/handler` · `component/service` · `component/repository` · `component/domain` ·
`component/middleware` · `component/kafka` · `component/database` · `component/frontend` ·
`component/ci` · `component/infra`
Map from `Type:` → backend = service/repository/etc; frontend = `component/frontend`; fullstack = both.

### Phase / Sprint
`phase/mvp` · `phase/1` … `phase/5` — and — `sprint/0` … `sprint/5`

### Size (story-point band — mirrors github-pr)
`size/S` (1–2 SP) · `size/M` (3–5 SP) · `size/L` (8 SP) · `size/XL` (13 SP)

### Estimation confidence (high/med/low — explicit per request)
`est/low` · `est/med` · `est/high` — **effort confidence**, not size.
Default: known refactor with file refs → `est/high`; new infra (Kafka, Temporal, SOPS) → `est/low`.

### Special
`blocked` · `adr-required` (story changes interfaces/contracts/migrations) · `good-first-task`

### Effort → field defaults (one row, copy when drafting)
| `Effort:` | `size/*` | Estimate SP | Priority | `est/*` default |
|---|---|---|---|---|
| S | `size/S` | 2 | Low/Medium | `est/high` |
| M | `size/M` | 5 | Medium | `est/med` |
| L | `size/L` | 8 | High | `est/low` |

---

## 3. Issue Templates

### Epic (sprint)
```markdown
## Goal
[copy sprint Goal]

## Why now
[copy sprint rationale]

## Stories
- [ ] #NN — Task N.1 …
- [ ] #NN — Task N.2 …
(filled after child issues created)

## Definition of Done (sprint)
[copy sprint DoD checklist]

## Source
docs/trd/<phase>/sprint-N.md

## Meta
Milestone: <MVP|Phase k> · Sprint: N · Depends on: sprint N-1
```
Labels: `type/epic`, `phase/*`, `sprint/N`, `class:*` (highest among children).

### User Story (task)
```markdown
## Story
As a [member|operator|platform], I want [capability], so that [value].
(derive from PRD user story if the task maps to one; else state the engineering outcome)

## Design
[copy task Design]

## Implementation
[copy task Implementation — keep file refs, e.g. Dockerfile:37, error_handler.go:116]

## Testing plan
[copy task Testing]

## Acceptance criteria
- [ ] [Progress signal — the observable outcome]
- [ ] Design + implementation done as specified
- [ ] Tests added; `go test -race ./...` green
- [ ] [Swagger: endpoints registered via `swag init -g main.go`]  ← only if API
- [ ] Tenant isolation respected (no cross-tenant read/write)

## Estimation
Effort: <S|M|L> · Story points: <2|5|8> · Confidence: <low|med|high>

## Epic
Epic: #<epic-number>

## Source
docs/trd/<phase>/sprint-N.md → Task N.M

## Change class
<low|medium|high> — [1-sentence justification]
```
Labels: `type/*`, `class:*`, `component/*`, `phase/*`, `sprint/N`, `size/*`, `est/*`,
`adr-required` if it touches interfaces/contracts/migrations.

---

## 4. Workflow

1. **Read the source.** Confirm which sprint(s): "All of phase-mvp, or sprint N?"
2. **Parse tasks.** For each `## Task N.M`, extract Title, Type, Effort, Progress signal, Design, Implementation, Testing, Swagger.
3. **Draft the Epic** + one Story per task, applying §2 labels and §3 templates. Compute size/SP/priority/est from the Effort row in §2.
4. **Present drafts** (titles + labels + estimation table). Ask: *"Create these N issues, or revise?"*
5. **On confirmation:** create Epic first (capture its number), then Stories with `Epic: #N`, then edit the Epic body to fill the `## Stories` checklist with real issue numbers.
6. **Link to Project board** (§5–6, reuse github-pr machinery): add each issue, set Status=Backlog, Priority, Size, Estimate.
7. **Report** the created issue numbers + board link.

**Title convention:** `<type>(<sprint-scope>): <task title>` —
e.g. `feat(s1): wire rule engine into earn path`, `fix(s0): remove .env from Docker image`.

---

## 5. Create Commands

```bash
# Epic
gh issue create \
  --title "epic(s<N>): <sprint theme>" \
  --assignee dydanz \
  --milestone "<MVP|Phase k>" \
  --label "type/epic,phase/mvp,sprint/<N>,class:<...>" \
  --body-file <drafted-epic-body.md>

# Story (repeat per task)
gh issue create \
  --title "<type>(s<N>): <task title>" \
  --assignee dydanz \
  --milestone "<MVP|Phase k>" \
  --label "type/<...>,class:<...>,component/<...>,phase/mvp,sprint/<N>,size/<...>,est/<...>" \
  --body-file <drafted-story-body.md>

# After stories exist, fill the epic checklist
gh issue edit <EPIC_N> --body-file <updated-epic-body.md>
```

Use `--body-file` with a temp file in the scratchpad dir (avoids heredoc-escaping issues
with multi-line Design/Implementation blocks).

---

## 6. Link to GitHub Project + Set Fields

Reuse the field-ID machinery from `.claude/skills/github-pr/SKILL.md` §4–5 (same project,
same `${go-loyalty-point_*}` env vars). After creating each issue:

```bash
# Add to board
gh project item-add ${go-loyalty-point_PROJECT_NUM} --owner dydanz \
  --url "$(gh issue view <N> --json url -q .url)"

# Get item id, then set fields (Status=Backlog, Priority, Size, Estimate)
gh project item-edit --project-id ${go-loyalty-point_PROJECT_ID} --id <ITEM_ID> \
  --field-id ${go-loyalty-point_FIELD_ESTIMATE} --number <SP>
# Priority / Size / Status → single-select-option-id (see github-pr §5 for option IDs)
```

If project env vars are not set, run the one-time discovery in `github-pr` §5 first, or
skip board linking and tell the operator labels-only was applied.

---

## 7. One-Time Label Bootstrap

Run once per repo. Idempotent-ish (`|| true` to ignore "already exists").

```bash
# Type
for t in epic:6f42c1 feature:1d76db fix:d73a4a chore:cfd3d7 security:b60205 test:fbca04 docs:0075ca; do
  gh label create "type/${t%%:*}" --color "${t##*:}" --force; done
# Class
gh label create "class:low"    --color c2e0c6 --force
gh label create "class:medium" --color fbca04 --force
gh label create "class:high"   --color d73a4a --force
gh label create "class:hotfix" --color b60205 --force
# Component
for c in handler service repository domain middleware kafka database frontend ci infra; do
  gh label create "component/$c" --color 0366d6 --force; done
# Phase + sprint
for p in mvp 1 2 3 4 5; do gh label create "phase/$p" --color 5319e7 --force; done
for s in 0 1 2 3 4 5; do gh label create "sprint/$s" --color bfd4f2 --force; done
# Size + estimation confidence
gh label create "size/S" --color c2e0c6 --force; gh label create "size/M" --color fbca04 --force
gh label create "size/L" --color e99695 --force; gh label create "size/XL" --color d73a4a --force
gh label create "est/low" --color ededed --force; gh label create "est/med" --color d4c5f9 --force
gh label create "est/high" --color 0e8a16 --force
# Special
gh label create "blocked" --color 000000 --force
gh label create "adr-required" --color 5319e7 --force
gh label create "good-first-task" --color 7057ff --force
```

Milestones (one-time):
```bash
for m in MVP "Phase 1" "Phase 2" "Phase 3" "Phase 4" "Phase 5"; do
  gh api repos/:owner/:repo/milestones -f title="$m" 2>/dev/null || true; done
```

---

## 8. Guardrails

- **Confirm before any `gh issue create/edit/close`** (Prime Directive).
- **No duplicates:** before creating, `gh issue list --search "<task title> in:title"`; if a match exists, ask whether to update instead.
- **One Epic per sprint.** If the epic already exists, attach new stories to it, don't recreate.
- **`adr-required`** on any story touching interfaces, session/event contracts, or DB migrations (mirrors sdlc Phase 4).
- **Security defects** (C1–C4) → `type/security` + `class:high` minimum, even if effort is S.
- **Acceptance criteria must be testable** — the Progress signal is criterion #1; if a task lacks one, derive an observable check before filing.
- Keep file references intact (e.g. `Dockerfile:37`) — they make the issue actionable.
