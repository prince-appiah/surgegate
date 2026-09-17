# SurgeGate

SurgeGate is a learning-first Go project exploring safe admission to scarce resources during bursts of traffic. It starts with Go fundamentals and will grow through experiments in backend engineering, concurrency, distributed systems, reliability, and performance.

## Why it exists

Imagine 50,000 customers trying to buy 1,000 sneakers. Reading inventory, creating an order, and then decrementing stock independently can sell the same final item twice. SurgeGate will investigate that failure and the engineering decisions needed to prevent it.

Correctness comes before throughput. For a fixed stock allocation:

```text
available + actively_reserved + sold = total_inventory
available >= 0
actively_reserved >= 0
sold >= 0
sold + actively_reserved <= total_inventory
```

Here, `available` means unallocated stock, not total stock. Expired reservations release inventory, so lifetime reservation creations may exceed the initial stock; active reservations plus completed sales must not. In a one-shot experiment without expiry or restocking, successful allocations cannot exceed initial inventory.

## Current milestone and architecture

**M00 — Go Fundamentals: study preparation.** The repository contains module metadata, a roadmap, and a learning package. No application, API, database, queue, workers, or completed experiments exist yet. No milestones have been completed.

Start with the [M00 study plan and resources](docs/learning/00-go-fundamentals.md). Implementation follows study and a learning checkpoint.

## Learning Journey

| Milestone | Topic | Status |
| --- | --- | --- |
| M00 | Go Fundamentals | 🚧 In Progress |
| M01 | Domain Modelling | ⏳ Upcoming |
| M02 | HTTP APIs | ⏳ Upcoming |
| M03 | PostgreSQL | ⏳ Upcoming |
| M04 | Concurrency Failure Experiment | ⏳ Upcoming |
| M05 | Safe Concurrent Inventory | ⏳ Upcoming |
| M06 | Reservations | ⏳ Upcoming |
| M07 | Redis | ⏳ Upcoming |
| M08 | Redis Streams | ⏳ Upcoming |
| M09 | Worker Pools | ⏳ Upcoming |
| M10 | Idempotency | ⏳ Upcoming |
| M11 | Reliability | ⏳ Upcoming |
| M12 | Rate Limiting and Admission Control | ⏳ Upcoming |
| M13 | Realtime Queue Updates | ⏳ Upcoming |
| M14 | Observability | ⏳ Upcoming |
| M15 | Load Testing | ⏳ Upcoming |
| M16 | Failure Experiments | ⏳ Upcoming |
| M17 | Docker | ⏳ Upcoming |
| M18 | Kubernetes | ⏳ Upcoming |

✅ Completed · 🚧 In Progress · ⏳ Upcoming

The [detailed roadmap](docs/learning/ROADMAP.md) defines the learning goals, expected evidence, and milestone branches. Completed notes will describe actual investigations, failures, choices, tests, and tradeoffs; study plans do not claim completed learning.

## Architecture evolution

The intended progression is sequential in-memory Go, then a Go HTTP API backed by PostgreSQL, then admission control, Redis Streams, and workers when experiments justify asynchronous processing. These are planned stages, not implemented capabilities. Kubernetes and cloud scaling come after understanding the local system.

The HTTP milestone already introduces concurrent request handling. Its shared state must be discussed before exposure; deliberately unsafe experiments stay isolated. Later milestones deepen these guarantees instead of implying the intermediate system is production ready.

## Local setup and testing

Work from this repository directory, `surgegate/`, inside the parent workspace. Follow the official [Go installation instructions](https://go.dev/doc/install), then verify `go version`. The module currently declares Go 1.26.0 as its baseline; use a supported stable toolchain that satisfies it.

`go.mod` was authored manually because Go was unavailable during bootstrap. `example.com/surgegate` is a temporary module path; replace it with the actual repository import path before publishing. There are no third-party dependencies and no `go.sum` yet.

There is nothing to run or test yet. Once the first package and tests exist, run from the repository root:

```bash
go fmt ./...
go vet ./...
go test ./...
```

Add `go test -race ./...` when concurrency is introduced. A passing race detector is not proof of inventory correctness or database atomicity; behavior and integration tests must check those separately.

## Branching strategy

`main` is intended to represent stable completed learning states after the initial bootstrap. Milestone branches use `learning/mNN-topic` as listed in the roadmap. Historical milestones retain the limitations and failures that motivated later work. Branch creation, switching, and integration are performed only when explicitly requested by the repository owner; commits are made by the owner.

## Project structure

```text
README.md
go.mod
docs/learning/
    ROADMAP.md
    00-go-fundamentals.md
```

Application packages and architecture, decision, and benchmark documents will be added when they have real content. No benchmark measurements exist yet. Future reports will identify the environment, implementation revision, workload, correctness counters, latency, throughput, and resource use.

## Future direction

The first goal is a visible record of learning Go and understanding why concurrent systems need particular guarantees. Open-source tooling, hosted admission infrastructure, and multi-tenant SaaS are possible later directions, not current implementation scope.
