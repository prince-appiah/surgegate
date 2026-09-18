# SurgeGate learning roadmap

This roadmap combines completed and planned work. M00 was completed on 2026-09-19; all later milestones are upcoming. Timing follows understanding, not a deadline.

For each substantial concept: explain the problem and relevance, research at least five videos and five readings, discuss 3–7 checkpoint questions, and pause for readiness. Then build a small increment, test it, explain the results, and document actual learning. New substantial concepts within a milestone get their own learning checkpoint. Basic programming knowledge is assumed; prioritize what differs in Go.

## Milestones and evidence

| Milestone | Learning and smallest meaningful scope | Evidence needed before completion | Planned branch |
| --- | --- | --- | --- |
| M00 | Environment; modules/packages/visibility; variables/constants; functions/multiple returns; structs/methods/pointers; slices/maps; errors; interfaces/composition; formatting/testing. Generics only where helpful. | Explain core Go code, complete tiny exercises, and run tests and standard checks. | `learning/m00-go-fundamentals` |
| M01 | Sequential in-memory product, sale, inventory, start-sale and purchase operations. | Boundary and sale-state tests: successful purchases never exceed stock. | `learning/m01-domain-model` |
| M02 | Standard-library HTTP handlers, routing, JSON, validation, middleware, status codes, contexts and graceful shutdown. | Handler tests for valid/invalid requests and cancellation; explain shared-state safety before concurrent HTTP use. | `learning/m02-http-api` |
| M03 | Explicit SQL from Go, PostgreSQL constraints, migrations, pooling and transactions. | Real-database integration tests and an explanation of transaction boundaries and inventory authority. | `learning/m03-postgresql` |
| M04 | Isolated naive in-memory/SQL concurrency failure experiments with goroutines. | Record observed outcomes for 100 stock / 10,000 attempts; distinguish Go data races from database logical races. Preserve failure evidence before fixing. | `learning/m04-concurrency-failure` |
| M05 | Compare mutexes, row locks, conditional atomic SQL updates and optimistic versions. | Concurrent invariant tests, measured baseline, and explanation of why a process-local mutex cannot coordinate replicas. | `learning/m05-safe-inventory` |
| M06 | Reservation state transitions, time/expiry, simulated payment completion and stock release. | Tests for expiry versus completion, repeated release, and nonnegative conserved inventory. | `learning/m06-reservations` |
| M07 | Redis data structures, atomic operations, TTL, persistence and caching responsibilities. | A justified small use case, real Redis tests, and failure/source-of-truth discussion. | `learning/m07-redis` |
| M08 | Redis Streams producer, consumer group, acknowledgements, pending entries and redelivery. | Demonstrate worker failure/redelivery and durable duplicate protection; specify when acknowledgement is safe. | `learning/m08-redis-streams` |
| M09 | Explicit bounded worker pools; channels, buffers, select, WaitGroups, cancellation, backpressure and saturation. | Compare 1/5/10/50/100 workers when feasible; test shutdown and bounds, measure rather than assume improvements. | `learning/m09-worker-pools` |
| M10 | End-to-end idempotency keys and at-least-once processing; uniqueness and replay semantics. | Duplicate HTTP and message tests, including concurrent duplicates and conflicting payloads. | `learning/m10-idempotency` |
| M11 | Safe retries/backoff, dead-letter handling, timeouts, recovery and circuit breaker concepts. | Controlled failure tests and explanation of retries after ambiguous outcomes. | `learning/m11-reliability` |
| M12 | Rate vs concurrency limits, queues and admission control; implement a token bucket manually first. | Deterministic limit/burst tests and documented overload policy; distributed version only if justified. | `learning/m12-rate-limiting` |
| M13 | Queue status with polling first, then SSE if justified; ordering and fairness. | Explain position/estimate uncertainty under parallel processing, retries and network latency. | `learning/m13-realtime-queue` |
| M14 | Structured logs, correlation, Prometheus and Grafana; OpenTelemetry later. | Observe requests, queue depth/wait, inventory, reservations, workers, retries, errors and latency; avoid unbounded metric labels. | `learning/m14-observability` |
| M15 | k6 scenarios, latency percentiles, throughput, resource use and bottleneck identification. | Repeatable reports from actual runs, progressively 100 to 100,000 users only as equipment permits; separate users, requests and achieved concurrency. | `learning/m15-load-testing` |
| M16 | Local/test-only worker crashes, Redis outage/latency, slow PostgreSQL, pool exhaustion, duplicates and timeouts. | Reproducible fault injection, correctness accounting and observed recovery limits. | `learning/m16-failure-experiments` |
| M17 | Containerize the understood API/workers/dependencies/observability stack; Docker Compose. | Repeatable startup, networking/configuration, persistence, health and shutdown checks. | `learning/m17-docker` |
| M18 | Pods, Deployments, Services, ConfigMaps, Secrets, probes, resource requests/limits, replicas and autoscaling. | Demonstrate scaling and recovery with resource evidence; worker scaling by queue pressure only when justified. | `learning/m18-kubernetes` |

## Dependencies that can move learning earlier

- M02 introduces concurrency through HTTP even before the dedicated M04 experiment. Discuss protection and test the boundary first; never silently expose a naive shared map/inventory implementation as safe.
- M03 transactions alone do not prove protection against overselling. Clearly label interim guarantees; M04 can reproduce a separate deliberately unsafe strategy.
- M08 redelivery already requires duplicate safety. Teach the minimum prerequisite before queue implementation; M10 expands it into the complete public contract.
- A small local PostgreSQL/Redis dependency container can be introduced when needed after a learning checkpoint. M17 is complete application containerization, not a ban on earlier local tooling.
- Cancellation, graceful shutdown and tests arrive when first needed and deepen later. Any substantial early introduction still needs the resource/readiness gate.

## Continuing questions

Where is inventory authoritative? What happens when two processes reserve the last item? What if a worker commits to PostgreSQL but crashes before ACK? How are duplicates and ambiguous retries handled? Who releases expired reservations? What happens if Redis is unavailable? What does fairness mean? Where is backpressure applied? How is saturation observed? Which guarantees hold under each stated failure assumption?

## Evidence and documentation

Each completed milestone note should explain what was actually learned/implemented, observed failures, tests, tradeoffs, architecture changes and remaining questions. Keep intentionally failing examples isolated and document how to reproduce them. Meaningful decisions get ADRs with context, decision, alternatives, tradeoffs and consequences. Do not add empty directories or speculative decision records.

Benchmark reports must include environment, implementation revision, workload/concurrency, initial stock, successful outcomes, duplicates, overselling, negative inventory, throughput, P50/P95/P99, errors and relevant resource/queue measurements. Mark unavailable metrics honestly. A future target experiment is 1,000 stock and 50,000 attempts with zero overselling, duplicate orders and negative stock; it is not a measured result.

No branches have been created beyond the initial main branch. Preserve milestone history; the repository owner controls commits and requests other Git changes explicitly.
