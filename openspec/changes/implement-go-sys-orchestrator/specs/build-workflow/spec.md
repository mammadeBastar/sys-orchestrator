## ADDED Requirements

### Requirement: Coordinate OpenSpec Changes During Build
The system SHALL provide build-phase commands that coordinate with the existing OpenSpec CLI.

#### Scenario: Propose build change
- **WHEN** a user runs `sys change propose add-login`
- **THEN** the system verifies the project is in build phase and invokes or instructs the equivalent OpenSpec proposal workflow for `add-login`

### Requirement: Require OpenSpec Apply Path
The system SHALL require implementation changes to flow through OpenSpec apply semantics.

#### Scenario: Apply build change
- **WHEN** a user runs `sys change apply add-login`
- **THEN** the system verifies the OpenSpec change exists and prints or invokes the apply workflow that uses `openspec-apply`

### Requirement: Require Superpowers Discipline During Apply
The system SHALL make Superpowers apply, debugging, testing, and verification discipline part of the build apply workflow.

#### Scenario: Codex apply skill invoked
- **WHEN** the Codex `sys-apply` skill is invoked for an OpenSpec change
- **THEN** the agent follows OpenSpec apply requirements and uses Superpowers methods for implementation planning, test-driven work, debugging, and verification

### Requirement: Support Explicit Design Changes During Build
The system SHALL provide `sys design-change` for foundational `/system` mutations during build phase.

#### Scenario: Foundation change requested during build
- **WHEN** a user runs `sys design-change change-auth-boundary`
- **THEN** the system creates or guides a controlled change path that records rationale, affected `/system` files, and impacted active OpenSpec changes

### Requirement: Archive Completed Build Changes
The system SHALL support archiving completed OpenSpec changes.

#### Scenario: Archive change
- **WHEN** a user runs `sys change archive add-login`
- **THEN** the system invokes or instructs the equivalent OpenSpec archive workflow and then runs system validation
