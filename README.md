# Sys Orchestrator

`sys` is a Go CLI for an agent-native monorepo workflow.

It keeps `/system` as the ratified project foundation during design and then coordinates implementation through OpenSpec and Superpowers during build.

## Install From Source

```bash
go run ./cmd/sys --help
```

During development, commands can be run with:

```bash
go run ./cmd/sys <command>
```

## Core Workflow

Initialize a repository:

```bash
sys init
```

This creates:

```text
.sys-orchestrator/
system/
```

The project starts in design phase.

## Design Phase

Design phase does not use OpenSpec for design decisions.

```bash
sys status
sys explore auth
sys capture
```

`sys explore` prints agent guidance based on the current `/system` foundation.

`sys capture` is used after a decision is finalized. The agent should update the relevant `/system` files and add a decision record under:

```text
system/architecture/decisions/
```

Freeze the foundation when ready to build:

```bash
sys design freeze
```

## Build Phase

Build phase uses OpenSpec for implementation changes.

```bash
sys change propose add-login
sys change apply add-login
sys change archive add-login
```

`sys change apply` points agents back to OpenSpec apply semantics and Superpowers implementation discipline.

Foundational `/system` mutations during build require:

```bash
sys design-change change-auth-boundary
```

## Agent Integration

Codex is the first-class v1 integration:

```bash
sys agent install codex
```

This installs project-local skills:

```text
.codex/skills/sys-explore/
.codex/skills/sys-capture/
.codex/skills/sys-apply/
.codex/skills/sys-design-change/
```

Minimal Cursor and Claude Code instruction scaffolds are also available:

```bash
sys agent install cursor
sys agent install claude
```

## Status And Validation

Human dashboard:

```bash
sys status
```

Machine-readable output:

```bash
sys status --json
```

Validation:

```bash
sys validate
```

Watch mode:

```bash
sys status --watch
```
