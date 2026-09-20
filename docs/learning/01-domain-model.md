# M01 — Basic SurgeGate domain model: study plan

Status: learning package prepared on 2026-09-19; no M01 implementation has started. Resource availability was checked during preparation. Dates are publication dates where the source exposed one; unavailable dates are labeled rather than inferred.

## The concept

Domain modelling turns the language and rules of a problem into code. For SurgeGate, the first model should make concepts such as product, flash sale, inventory, sale state, and purchase attempt understandable without knowing anything about HTTP, PostgreSQL, or Redis.

This milestone uses a small part of domain-driven thinking: use the problem's language, put behavior beside the state it protects, and prevent invalid state where practical. It does not introduce full Domain-Driven Design, repositories, aggregates, Clean Architecture, dependency-injection layers, CQRS, or event sourcing.

## Why SurgeGate needs it

The M00 exercise exposes `Inventory.Available`, so any caller can construct `Inventory{Available: -10}` or directly bypass `Purchase`. That was useful while learning Go syntax, but it cannot reliably protect the core rule.

M01 asks where each rule belongs:

- A product has identity and descriptive data.
- A flash sale has identity, a product, configured stock, and a lifecycle.
- A sale cannot accept a purchase before it starts.
- A successful purchase consumes exactly one available unit in this milestone.
- Available inventory cannot become negative.
- Sequential successful purchases cannot exceed configured stock.

The model should make valid operations easy and invalid transitions explicit. It remains entirely in memory and sequential; concurrency safety belongs to later experiments.

## What to understand

- **Domain language:** names in code should match the problem being discussed.
- **State and behavior:** structs hold state; methods express operations such as `Start` and `AttemptPurchase` rather than exposing arbitrary setters.
- **Invariants:** rules that must remain true after every public operation.
- **Valid construction:** constructors can validate input and initialize unexported fields, especially when a type's zero value is not useful.
- **Identity and values:** a product or sale keeps identity over time; quantities and states describe it. We only need this distinction informally in M01.
- **State transitions:** operations can be allowed or rejected according to the current sale state.
- **Package cohesion:** group code that changes for the same domain reason; avoid a folder for every type or a generic `models` package.
- **Error vocabulary:** callers need stable errors for conditions such as invalid inventory, sale not started, already started, and out of stock.
- **Behavior tests:** test rules and observable state transitions, not private implementation details.

## Planned scope after readiness

The smallest useful version may support creating a product and flash sale, setting initial stock during valid construction, starting the sale, and attempting one-unit purchases sequentially. Exact types and package boundaries will be chosen during implementation based on the rules, not copied from a template.

The existing M00 `inventory` package is a learning artifact. We will decide explicitly whether to evolve, move, or replace it; historical Git commits preserve the exercise either way.

Out of scope: HTTP, persistence, repositories, interfaces added only for mocking, customer accounts, orders, reservations, expiry, payments, queues, goroutines, locks, rate limits, and distributed coordination.

## Videos

The talks are older because the relevant Go design principles are stable. Watch the first talk as the overview, then use selected talks to compare viewpoints. Do not copy any speaker's complete project layout.

| Resource / source | Date and classification | What it teaches and why selected |
| --- | --- | --- |
| [How Do You Structure Your Go Apps?](https://www.youtube.com/watch?v=oL6JBUk6tj0) — Kat Zień / Gopher Academy | 2018-09-11. FOUNDATIONAL. | Compares flat layouts, grouping by context, DDD, and hexagonal architecture. Selected because it explains that structure follows needs and has tradeoffs. For M01, focus on flat/domain grouping and the conclusion; treat the advanced architectures as context only. |
| [Package Oriented Design](https://www.youtube.com/watch?v=spKM5CyBwJA) — William Kennedy / GopherCon India | 2017; exact upload date unavailable. FOUNDATIONAL. | Explains package purpose, cohesion, naming, and dependency direction. Selected for thinking about what belongs together. Its vendoring and prescribed layout details predate modern modules and are not instructions for this project. |
| [SOLID Go Design](https://www.youtube.com/watch?v=zzAdEt3xZ1M) — Dave Cheney / Golang UK Conference | 2016-09-07. FOUNDATIONAL. | Reinterprets design principles through small Go interfaces and composition. Selected to contrast Go with class hierarchies. Do not add interfaces mechanically; M01 begins with concrete types. |
| [Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg) — Peter Bourgon / GopherCon Europe | 2018-08-25. FOUNDATIONAL. | Covers judgment in repository structure, program design, dependencies, and testing. Selected for its practical emphasis and warning against dogmatic rules. Watch program-design and testing portions first. |
| [Simplicity is Complicated](https://www.youtube.com/watch?v=rFejpH_tAHM) — Rob Pike / dotGo | 2015-12-02. FOUNDATIONAL. | Explains why simplicity is an active design goal rather than absence of thought. Selected as a counterweight to importing elaborate TypeScript/Java architectures into a small Go model. |
| [Domain Driven, Data Oriented Architecture](https://www.youtube.com/watch?v=bQgNYK1Z5ho) — William Kennedy / Golang Charlotte | Listed by Ardan Labs as 2024; exact upload date unavailable. CURRENT COMPARISON. | A modern, opinionated architecture view covering domains, validation, data flow, and package boundaries. Selected for comparison after the basics. It is much broader than M01; do not reproduce its service layers or infrastructure here. |

## Written resources

| Resource / source | Date and classification | What it teaches and why selected |
| --- | --- | --- |
| [Organizing a Go module](https://go.dev/doc/modules/layout) — Go team | Living documentation; page date unavailable. CURRENT IMPLEMENTATION. | Official guidance for starting small, adding supporting packages when useful, and using `internal` for server code. Selected as the current authority over older folder-layout talks. |
| [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments) — Go project wiki | Living guidance; page date unavailable. CURRENT IMPLEMENTATION. | Practical naming, interfaces, copying, errors, and receiver guidance. Selected as a review checklist, not a rigid style specification. |
| [Effective Go](https://go.dev/doc/effective_go) — Go team | Written for Go's 2009 release; explicitly not actively updated. FOUNDATIONAL. | Read the names, data, methods, interfaces, and embedding sections for idioms. Selected for stable language principles; use the current module-layout guide for project structure and ignore its gaps around modern Go. |
| [Errors are values](https://go.dev/blog/errors-are-values) — Rob Pike / Go blog | 2015-01-12. FOUNDATIONAL. | Shows that error behavior is part of API design rather than repetitive exception substitution. Selected to help define clear domain failures and keep control flow readable. |
| [Introduction to DDD Lite in Go](https://threedots.tech/post/ddd-lite-in-go-introduction/) — Robert Laszczak / Three Dots Labs | 2020-07-01. FOUNDATIONAL. | Demonstrates behavior-oriented methods, private state, valid constructors, invariants, and database-independent domain code. Selected because those ideas map directly to SurgeGate. Stop before repository and application-layer patterns; they are out of scope. |
| [Package Oriented Design](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html) — William Kennedy / Ardan Labs | 2017-02-24; updated 2017-02-28. FOUNDATIONAL WITH OUTDATED SETUP DETAILS. | Explains package purpose and design validation. Selected as written support for the talk. Its vendor-centric directory prescription is historical; apply principles, not its exact tree. |
| [How We Implemented Domain-Driven Development in Golang](https://engineering.grab.com/domain-driven-development-in-golang) — Grab engineering | Publication date unavailable on the retrieved page. FOUNDATIONAL CASE STUDY. | Shows how a real team identified domains, relationships, entities, and business terminology. Selected as a case study to critique. Its repositories, layers, events, and interfaces exceed M01 and should not be copied. |

## Suggested study sequence

1. Read this plan and write a one-paragraph description of a flash sale using only domain language.
2. Read the official module-layout guide and relevant Go Code Review Comments.
3. Watch Kat Zień's overview, then read the DDD Lite sections through valid in-memory state.
4. Review package-oriented design and Rob Pike's simplicity talk. Compare the advice rather than treating either as a template.
5. Sketch rules and state transitions in plain text. Do not sketch database tables or HTTP endpoints.
6. Answer the checkpoint questions below. We will discuss unclear points before choosing the smallest code structure.

## Checkpoint questions

1. What is a domain invariant, and which object should prevent available inventory from becoming negative?
2. Why might `Available` become an unexported field with a constructor and query method? What invalid state does that prevent?
3. In M01, how are a `Product` and `FlashSale` different from a quantity or sale state? Why does identity matter?
4. Why is `sale.Start()` clearer and safer than allowing callers to assign `sale.State = Started` directly?
5. What should happen when a purchase is attempted before the sale starts, after stock reaches zero, and while stock remains?
6. Why should we begin with concrete structs and behavior tests instead of repositories, storage interfaces, or a full architecture template?

After studying, bring answers or questions. Implementation waits until readiness and will proceed in small learner-authored steps.
