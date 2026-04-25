## 1. Documentation Inventory

- [x] 1.1 Review current `README.md` and identify missing documentation sections against the project-documentation spec.
- [x] 1.2 Review current CLI commands and synced OpenSpec specs to ensure documented behavior matches implementation.

## 2. README Expansion

- [x] 2.1 Rewrite the README introduction with the project purpose and mental model.
- [x] 2.2 Add a table of contents for the expanded documentation.
- [x] 2.3 Document installation and running from source.
- [x] 2.4 Document the core lifecycle from initialization through design, freeze, build, and archive.
- [x] 2.5 Document the canonical `/system` tree and responsibilities of each area.
- [x] 2.6 Document design-phase workflows and examples.
- [x] 2.7 Document build-phase workflows and examples.
- [x] 2.8 Document Codex, Cursor, and Claude Code agent integrations.
- [x] 2.9 Document status, validation, freeze behavior, and troubleshooting.
- [x] 2.10 Add a complete v1 command reference.
- [x] 2.11 Add contributor guidance for keeping README documentation aligned with behavior changes.

## 3. Review And Verification

- [x] 3.1 Check README examples against implemented command names.
- [x] 3.2 Check README claims against current OpenSpec specs.
- [x] 3.3 Run `go test ./...` to confirm documentation-only changes did not disturb code.
- [x] 3.4 Run OpenSpec status and validation for this change.
