# M00 — Go fundamentals: study plan

Status: study preparation, awaiting readiness. No exercises, tests, or implementation have been completed. Resources researched on 2026-09-17; availability checked using source pages/search results, not full video playback. Dates below are publication dates when available; a search crawl date is not an update date.

## Why this milestone exists

Before modeling stock, we need to reason about Go values, mutation, visibility, failure and tests. A copied inventory struct, a shared slice backing array, or a mishandled error can change behavior even in a sequential program. Understanding these mechanics makes later concurrency experiments explainable.

Go packages organize code and exported identifiers define their public surface. Modules group packages and declare dependency/toolchain requirements. Structs describe data, methods attach behavior, pointers enable access to an existing value, and small interfaces describe behavior without requiring class inheritance. Errors are explicit values returned by functions. Tests specify the behavior we intend to preserve. Use the resources below to explore these ideas before repository implementation.

For an experienced TypeScript/Node engineer, focus on value copying and pointer receivers, slice aliasing, zero values and nil, explicit multiple returns/errors, implicit interface satisfaction, composition, package visibility and the built-in toolchain. Do not recreate NestJS layers or dependency injection machinery.

## Videos

These are five choices, not a requirement to watch multiple complete courses. Start with one main course, then use the others for targeted explanations. FOUNDATIONAL means enduring language concepts, not current installation/dependency advice.

| Resource / source | Date and classification | What to study and why selected |
| --- | --- | --- |
| [Go Programming – Golang Course with Bonus Projects](https://www.youtube.com/watch?v=un6ZyFkqFKo) — Boot.dev / Lane Wagner; published by freeCodeCamp | 2023; exact upload date not verified from the retrieved primary page. FOUNDATIONAL. | Main course option for functions, structs, interfaces, errors, slices, maps and pointers. Selected for a structured sequence with practice. Study fundamentals only; defer concurrency and backend bonus projects. |
| [Go Programming — Full Course](https://www.youtube.com/watch?v=V-lI7AmusGs) — Tech With Tim | 2026 according to indexed source metadata; exact day unavailable. CURRENT IMPLEMENTATION overview plus foundations. | Alternative recent overview of types, collections, functions, structs, interfaces, errors, generics and pointers. Selected for chapter-based review; skip generic programming basics you already know and defer concurrency. Official docs remain the authority for setup. |
| [Learn Go Programming — Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU) — Michael Van Sickle / freeCodeCamp | 2019-06-20. FOUNDATIONAL. | Targeted explanations: slices at 1:47:53, maps/structs at 2:17:20, pointers at 4:03:57, functions at 4:21:30 and interfaces at 4:57:59. Selected for detailed treatment of the Go value model. Skip old setup guidance and defer goroutines/channels. |
| [GopherCon 2016: Understanding nil](https://www.youtube.com/watch?v=ynoY2xz-F8s) — Francesc Campoy / Gopher Academy | 2016-08-18. FOUNDATIONAL. | Study nil, zero values and interface pitfalls after the basics. Selected to challenge JavaScript null/undefined intuition. Defer channel-specific sections. Source listing verified; direct page extraction returned an error during research. |
| [Go in 100 Seconds](https://www.youtube.com/watch?v=446E-r0rXHI) — Fireship | 2021-10-07. FOUNDATIONAL orientation. | A brief language/tooling overview. Selected as a quick map before deeper study, not as sufficient instruction for any checkpoint. Use official documentation for modern details. |

## Written resources

Official living documentation often has no per-page publication/update date. Such dates are marked unavailable rather than invented.

| Resource / source | Date and classification | What it teaches and why selected |
| --- | --- | --- |
| [Download and install](https://go.dev/doc/install) — Go team | Living docs; date not stated. CURRENT IMPLEMENTATION. | Toolchain setup and version verification. Selected as the authoritative replacement for old videos' environment instructions. |
| [A Tour of Go](https://go.dev/tour/welcome/1) — Go team | Living tutorial; date not stated. FOUNDATIONAL, current hosted edition. | Packages, exported names, functions, values, pointers, structs, slices, maps, methods and interfaces. Selected for small interactive examples; complete Basics and Methods and interfaces, leaving concurrency for later. |
| [Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module) — Go team | Living docs; date not stated. CURRENT IMPLEMENTATION. | Module/package structure, calling functions, errors, slices, maps and tests across the linked sequence. Selected to explain go.mod and visibility without a framework. Read first; do exercises when ready. |
| [Add a test](https://go.dev/doc/tutorial/add-a-test) — Go team | Living docs; date not stated. CURRENT IMPLEMENTATION. | Test files, testing.T, failures and go test. Selected as a small official first-test example. |
| [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests) — Chris James and contributors | Living book; exact update date not verified. FOUNDATIONAL, maintained online edition. | Hello World, arrays/slices, structs/methods/interfaces, pointers/errors and maps. Selected to learn through behavior tests. Defer concurrency, mocking frameworks/large abstractions and application chapters. |
| [Effective Go](https://go.dev/doc/effective_go) — Go team | Originally 2009; the page warns it is not actively updated. FOUNDATIONAL. | Formatting, naming, composition, methods and interfaces. Selected for idiomatic reasoning; it does not cover modern modules or generics and is not a complete current guide. |
| [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors) — Damien Neil and Jonathan Amsterdam / Go team | 2019-10-17. FOUNDATIONAL. | Error wrapping with %w and inspection with errors.Is/As. Selected to connect contextual errors to preserved error identity; study after ordinary error returns. |
| [go command documentation](https://pkg.go.dev/cmd/go) — Go standard library | Versioned living reference; no single article date. CURRENT IMPLEMENTATION. | Look up mod, fmt, vet and test as needed. Selected to understand what commands actually do instead of memorizing scripts. |

## Study sequence

1. Read installation guidance and the module tutorial. Understand the difference between a module, package and executable. Local setup verified on 2026-09-17: Go 1.27.1 on darwin/amd64 detects this repository's go.mod and satisfies its Go 1.26.0 baseline.
2. Work through the Tour's basic language and data sections, supported by one course. Pay attention to copying, mutation and aliasing rather than spending time on familiar control flow.
3. Study methods, pointers, interfaces, nil and explicit errors. Read the selected Effective Go sections, then error wrapping.
4. Read the official test example and the early Learn Go with Tests chapters. Understand go fmt, go vet and go test.
5. Discuss the questions below and indicate readiness. The first repository exercise can be a tiny sequential inventory operation and its boundary test, built together in small steps.

Generics can be recognized without forcing them into the initial model. Concurrency, APIs, databases and queues are later learning packages.

## Checkpoint questions

1. What is the difference between a module, a package and package main? What does go.mod declare, and how does capitalization affect visibility?
2. If an Inventory struct is passed to a function by value, what is copied? When would a pointer receiver be appropriate, and how does this differ from passing a JavaScript object?
3. Can two slices share storage? How can append change that relationship? What differs between a nil slice and a nil map when adding elements?
4. How would an operation return an out-of-stock error alongside a result? Why wrap an error, and why might errors.Is be preferable to comparing its message?
5. How does a type satisfy an interface without an implements declaration? When would an interface improve this small program, and when would a concrete type be clearer?
6. How would you test purchases with zero, one and several units of stock? What do go fmt, go vet and go test each check—and what do they not prove?

These questions guide discussion, not grading. No implementation or claimed learning outcomes should be added until study and readiness. M00 exits after small exercises and actual checks, with the learner able to explain core Go code.
