---
name: github-pr
description: >
  Create and manage GitHub PRs for the go-loyalty-point project with full metadata:
  assignee, labels (class/type/component/phase/size), GitHub Project linking,
  and project field values (Priority, Size, Estimate, Start Date, Target Date).
  Use when the user says "create PR", "open PR", "submit PR", or "link to project".
---

# GitHub PR Management — go-loyalty-point

## 1. Required PR Fields

Every PR must include:

| Field | Value |
|---|---|
| Assignee | `dydanz` |
| Labels | `class:*` + `component/*` + `type/*` + `phase/*` + `size/*` |
| Linked issue | `Closes #N` in body |
| KLW ticket ref | `GLP: #N` in body |
| Body sections | `## Summary`, `## Risk`, `## Rollback`, `## Test plan` |

## 2. Create PR Command

```bash
gh pr create \
  --title "<type>(<scope>): <description>" \
  --base main \
  --assignee dydanz \
  --label "<class:*>,<component/*>,<type/*>,<phase/*>,<size/*>" \
  --body "$(cat <<'EOF'
## Summary
- bullet points

## Risk
Low/Medium/High. One line rationale.

## Rollback
Concrete rollback steps.

## Test plan
- [ ] item

GLP: #<issue-number>

Closes #<issue-number>

🤖 Generated with [Claude Code](https://claude.com/claude-code)
EOF
)"
```

## 3. Label Taxonomy

### Change Class (exactly one)
| Label | When |
|---|---|
| `class:low` | Typo, doc, comment, config-only |
| `class:medium` | Feature, refactor, component-scoped |
| `class:high` | Breaking change, arch, interface contract, migration |
| `class:hotfix` | Production incident — ship first, doc after |

### Type (one)
`type/user-story` · `type/chore`

### Component (all that apply)
`component/runtime` · `component/adapter` · `component/session` · `component/brain` · `component/identity`

### Phase (current active)
`phase/9` · `phase/10` · `phase/11` · `phase/12`

### Size
`size/S` (1-2 SP) · `size/M` (3-5 SP) · `size/L` (8 SP)

### Special
`adr-required` — add when PR changes component interfaces, session contracts, runtime wiring, or conventions

## 4. Link PR to GitHub Project

```bash
# Add PR to project board
gh project item-add ${go-loyalty-point_PROJECT_NUM} \
  --owner dydanz \
  --url "$(gh pr view --json url -q .url)"
```

## 5. Set Project Fields

Requires `read:project` + `project` token scopes. Set after linking.

### Priority
```bash
# Values: "No priority" | "Urgent" | "High" | "Medium" | "Low"
gh project item-edit --project-id ${go-loyalty-point_PROJECT_ID} --id <ITEM_ID> \
  --field-id ${go-loyalty-point_FIELD_PRIORITY} --single-select-option-id <OPTION_ID>
```

### Size
```bash
# Values: "XS" | "S" | "M" | "L" | "XL"
gh project item-edit --project-id ${go-loyalty-point_PROJECT_ID} --id <ITEM_ID> \
  --field-id ${go-loyalty-point_FIELD_SIZE} --single-select-option-id <OPTION_ID>
```

### Estimate (story points)
```bash
# Values: 1 | 2 | 3 | 5 | 8 | 13
gh project item-edit --project-id ${go-loyalty-point_PROJECT_ID} --id <ITEM_ID> \
  --field-id ${go-loyalty-point_FIELD_ESTIMATE} --number <SP>
```

### Start Date / Target Date
```bash
gh project item-edit --project-id ${go-loyalty-point_PROJECT_ID} --id <ITEM_ID> \
  --field-id ${go-loyalty-point_FIELD_START_DATE} --date "YYYY-MM-DD"

gh project item-edit --project-id ${go-loyalty-point_PROJECT_ID} --id <ITEM_ID> \
  --field-id ${go-loyalty-point_FIELD_TARGET_DATE} --date "YYYY-MM-DD"
```

### One-time field ID discovery
```bash
# List your projects
gh api graphql -f query='
{ viewer { projectsV2(first:10) { nodes { id number title } } } }'

# List fields + select options for a project
gh api graphql -f query='
{ node(id:"<PROJECT_ID>") { ... on ProjectV2 {
  fields(first:20) { nodes {
    ... on ProjectV2FieldCommon { id name }
    ... on ProjectV2SingleSelectField { id name options { id name } }
  }}
}}}'
```

Store discovered IDs in `.env`:
```
go-loyalty-point_PROJECT_NUM=1
go-loyalty-point_PROJECT_ID=PVT_xxxxxxxxxxxx
go-loyalty-point_FIELD_PRIORITY=PVTF_xxxxxxxxxxxx
go-loyalty-point_FIELD_SIZE=PVTF_xxxxxxxxxxxx
go-loyalty-point_FIELD_ESTIMATE=PVTF_xxxxxxxxxxxx
go-loyalty-point_FIELD_START_DATE=PVTF_xxxxxxxxxxxx
go-loyalty-point_FIELD_TARGET_DATE=PVTF_xxxxxxxxxxxx
```

## 6. SDLC Gate (run before converting draft to ready)

```bash
go build ./... && go test ./... && go vet ./...
```

- [ ] Build clean, tests green, vet clean
- [ ] Body has: Summary, Risk, Rollback, Test plan, `Closes #N`, `GLP: #N`
- [ ] `class:high` — PRD/TRD linked, manual smoke noted
- [ ] `adr-required` — ADR drafted or linked

## 7. Post-Merge

- `Closes #N` auto-closes linked issue when CI passes
- `adr-required` — write `docs/adr/ADR-*.md` within 48h of merge
- Implementation deviated from plan — update `docs/dev-plan/phase-*.md`
