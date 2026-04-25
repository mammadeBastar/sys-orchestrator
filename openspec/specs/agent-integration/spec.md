## Purpose

Define the agent-runtime integrations for Codex, Cursor, and Claude Code.

## Requirements

### Requirement: Install Codex Skills
The system SHALL install Codex-native sys skills with `sys agent install codex`.

#### Scenario: Install Codex integration
- **WHEN** a user runs `sys agent install codex`
- **THEN** the system creates project-local `sys-explore`, `sys-capture`, `sys-apply`, and `sys-design-change` skill directories under `.codex/skills/`

### Requirement: Avoid Role-Specific Install Commands
The system SHALL avoid requiring users to install separate design, frontend, or backend agent roles.

#### Scenario: Codex integration installed once
- **WHEN** a user has run `sys agent install codex`
- **THEN** the installed skills infer role from current working directory and project phase

### Requirement: Generate Minimal Cursor Instructions
The system SHALL generate minimal Cursor rules with `sys agent install cursor`.

#### Scenario: Install Cursor integration
- **WHEN** a user runs `sys agent install cursor`
- **THEN** the system writes `.cursor/rules/sys-orchestrator.mdc` with sys workflow rules and file-access guidance

### Requirement: Generate Minimal Claude Code Instructions
The system SHALL generate minimal Claude Code instructions with `sys agent install claude`.

#### Scenario: Install Claude Code integration
- **WHEN** a user runs `sys agent install claude`
- **THEN** the system creates or updates a clearly marked sys-orchestrator section in `CLAUDE.md`

### Requirement: Preserve Existing Agent Files
The system SHALL avoid overwriting unrelated user-authored agent instructions.

#### Scenario: Existing CLAUDE file
- **WHEN** `CLAUDE.md` already contains user-authored content
- **THEN** `sys agent install claude` updates only the marked sys-orchestrator section or appends one if missing
