# DDIA Mental Model Reference

Use this file as a design checklist inspired by data-intensive system principles. Do not quote it as book notes; translate the models into project-specific questions, risks, and `/system` decisions.

## Data Models

- Identify the natural model: relational, document, graph, log, search index, cache, queue, object storage, or derived view.
- Check whether the access pattern matches the model. A mismatch usually appears as unbounded joins, duplicated fields, ad hoc denormalization, or slow fan-out reads.
- Name the canonical source when the same fact appears in multiple places.

## Storage And Retrieval

- Identify the write path, read path, indexes, cache layers, and compaction or retention behavior.
- Check hot keys, high-cardinality fields, write amplification, read amplification, and whether indexes support the actual queries.
- Define what happens when cached or indexed data is stale.

## Encoding And Evolution

- Check backward compatibility, forward compatibility, unknown fields, defaults, nullable fields, enum expansion, and version negotiation.
- Design rolling deploys where old and new code can coexist.
- Treat externally visible payloads as contracts; do not rely on synchronized client and server upgrades.

## Replication

- Identify whether replicas are synchronous or asynchronous and what stale reads mean for users.
- Check read-your-writes, monotonic reads, failover behavior, split-brain risk, replay, and conflict resolution.
- Define the recovery path after a replica falls behind or diverges.

## Partitioning

- Identify the partition key, skew risk, hot partitions, cross-partition queries, rebalancing behavior, and fan-out cost.
- Avoid keys that concentrate writes around time, tenant size, celebrity users, or a single global aggregate.
- Define how pagination, ordering, and uniqueness behave across partitions.

## Transactions

- Name which invariants require atomicity and which can tolerate eventual repair.
- Check isolation assumptions: lost updates, write skew, phantom reads, uniqueness races, and read-modify-write hazards.
- Prefer explicit idempotency and constraints over trusting clients, timing, or worker ordering.

## Consistency

- State the consistency model users and services can rely on: linearizable, causal, read-your-writes, monotonic, eventual, or best-effort.
- Define what clients should do after ambiguous success, timeout, retry, or failover.
- Avoid presenting eventually consistent state as authoritative without explaining freshness.

## Batch And Stream Processing

- Identify bounded versus unbounded data, event time versus processing time, replay behavior, checkpointing, backfill, and duplicate handling.
- Require idempotent consumers or deduplication for at-least-once delivery.
- Define late event handling, poison-message handling, and operational controls for pause, resume, and replay.

## Derived Data

- Name the source of truth, transformation, destination, freshness target, recomputation path, and repair mechanism.
- Check that search indexes, caches, projections, analytics tables, and materialized views can be rebuilt.
- Define how users and operators detect and handle divergence from the source of truth.
