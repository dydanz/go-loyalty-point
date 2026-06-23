# /analyze-codebase

Analyzes the codebase to understand its current structure, patterns, dependencies, and technical health. Use this as the starting point before planning any feature or refactor.

## When Invoked

If no specific focus area is provided, present the following menu:

```
What would you like to analyze?

1. Full codebase overview (architecture, structure, dependencies)
2. Frontend architecture (components, state, data fetching)
3. Backend architecture (Go packages, interfaces, patterns)
4. Infrastructure (Terraform, K8s manifests)
5. Testing coverage and gaps
6. Specific feature or module: [provide name]

Reply with a number or describe what you'd like to understand.
```

## Analysis Process

### Step 1: Discovery

Spawn parallel sub-agents to locate and understand different parts of the codebase:

**Frontend Discovery:**
- Locate package.json, tsconfig.json, framework config files
- Map the `src/` directory structure
- Identify state management (Redux, Zustand, Context usage)
- Find API client patterns
- Count component files and identify shared vs feature-specific

**Backend Discovery:**
- Locate go.mod, identify Go version and key dependencies
- Map the package structure (cmd/, internal/, pkg/)
- Identify main entry points
- Find interface definitions (ports)
- Identify database layer (ORM, raw SQL, migrations)

**Infrastructure Discovery:**
- Locate Terraform files, identify modules and environments
- Find Kubernetes manifests or Helm charts
- Identify CI/CD pipeline files
- Note cloud provider and key services

### Step 2: Pattern Analysis

For each discovered component:
- Identify the architectural pattern in use
- Note deviations from the pattern (inconsistencies)
- Identify technical debt (TODO comments, deprecated patterns)
- Note missing standard components (no tests, no error handling, no logging)

### Step 3: Synthesis

Produce a structured analysis report:

```markdown
## Codebase Analysis: [Project Name]
Date: YYYY-MM-DD
Commit: [current git hash]

### Architecture Summary
[2-3 sentence overview of the overall system design]

### Frontend
- Framework: [framework + version]
- Structure: [feature-based / component-based / mixed]
- State management: [tools used]
- Key observations: [3-5 bullet points with file:line references]
- Gaps/Concerns: [what's missing or inconsistent]

### Backend
- Language: Go [version]
- Architecture: [clean arch / layered / flat]
- Key patterns: [DI approach, error handling, logging]
- Key observations: [3-5 bullet points with file:line references]
- Gaps/Concerns: [missing patterns, inconsistencies]

### Infrastructure
- Cloud: [provider]
- IaC: [Terraform / CDK / CloudFormation]
- Container orchestration: [K8s version, approach]
- Key observations: [3-5 bullet points]

### Testing
- Coverage: [% if available]
- Test types present: [unit / integration / E2E]
- Key gaps: [what's not tested]

### Priority Recommendations
1. [Most critical improvement]
2. [Second priority]
3. [Third priority]
```

## Important Guidelines

- Always reference specific file paths and line numbers in findings
- Distinguish between "this is an intentional pattern" and "this looks like drift"
- Do not change any code during analysis — this is read-only
- If the codebase is large, analyze the most active packages first (git log --name-only)
- Present the analysis as findings, not prescriptions — the team decides what to act on
