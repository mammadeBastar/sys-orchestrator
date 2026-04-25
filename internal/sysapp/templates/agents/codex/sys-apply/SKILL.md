---
name: sys-apply
description: Apply OpenSpec changes using sys workflow, openspec-apply, and Superpowers discipline.
---

## Purpose

Use this skill during build phase to implement an OpenSpec change while preserving `/system` as the foundation truth. This skill orchestrates local sys rules; it does not replace OpenSpec or Superpowers.

## Initial Checks

1. Run or read `sys status --json`.
2. Confirm the project is in build phase.
3. Confirm the named OpenSpec change exists.
4. Read the OpenSpec proposal, design, specs, and tasks for the change.
5. Read the relevant `/system` files allowed for the current role before editing implementation code.

## Phase Rules

- Build phase is required for implementation.
- Design phase work should use `sys-explore` and `sys-capture` instead.
- OpenSpec owns change planning and task tracking during build.
- Superpowers owns implementation discipline: planning, test-driven development, systematic debugging, and verification.
- Frozen /system files are not implementation files.

## Role And File Access

- Infer role from the current working directory.
- Read the allowed `/system` files for that role before deciding how to implement.
- Frontend agents should treat `system/contracts/`, `system/flows/`, and `system/modules/frontend.md` as their build context.
- Backend agents should treat `system/contracts/`, `system/flows/`, `system/modules/backend.md`, `system/data/`, and `system/obs/` as their build context.
- Change agents may read OpenSpec change files and the `/system` files required by that change.

## Workflow

1. Invoke or follow the local OpenSpec apply workflow for the named change.
2. Use Superpowers skills for implementation planning, TDD, debugging, and verification when the environment provides them.
3. Work through OpenSpec tasks in order and mark each task complete only after implementation and verification.
4. Keep edits scoped to the change and the current task.
5. Compare implementation needs against `/system` truth before changing behavior.
6. If implementation reveals design drift, pause and explain the mismatch.
7. Escalate to `sys design-change` before mutating foundational truth.

## Validation

- Run focused tests for the changed behavior.
- Run broader tests required by the OpenSpec change before completion.
- Re-read modified code and relevant `/system` files to check alignment.
- Confirm no frozen /system files changed accidentally.
- Confirm OpenSpec task checkboxes accurately reflect completed work.

## Stop Conditions

- Stop if `sys status --json` does not show build phase.
- Stop if the OpenSpec change is missing or blocked.
- Stop if the requested implementation contradicts `/system` truth.
- Stop if a foundation mutation is required; use `sys design-change`.
- Stop if tests fail and systematic debugging has not isolated the cause.

## Do Not

- Do Not implement outside an OpenSpec change during build phase.
- Do Not mutate frozen /system files as part of normal apply.
- Do Not copy full OpenSpec or Superpowers instructions into this skill; invoke or follow them.
- Do Not mark tasks complete without fresh verification.
- Do Not hide design drift by forcing code to fit an outdated spec.
