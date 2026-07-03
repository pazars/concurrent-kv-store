# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

A learning project: a thread-safe in-memory key-value store in Go, built to understand
practical concurrency hands-on (Phase 2 of a backend learning roadmap). The goal is **not**
a shippable product — it's to build the *same* store under progressively better
synchronization strategies (single global mutex → RWMutex → sharded locking) and measure the
difference between them.

The build order is the point. Each step must run and pass its tests under the race detector
before the next one begins. Do not jump ahead to a finished, fully-sharded, networked
implementation — each layer of synchronization exists to solve a problem you're meant to hit
first. See `README.md` for the ordered roadmap and its mapping to GitHub issues (#1–#16).

## Commands

The store itself is standard-library-only Go, developed and verified under the race detector:

```
go test -race ./...        # the core loop — must stay clean at every step
go test -race -run TestName ./...   # a single test
go test -bench . -benchmem ./...    # benchmarks (locking-strategy comparison)
go vet ./...
```

The `learn-go-basics/` warm-up exercises are currently standalone `package main` files with
no module, run individually:

```
go run learn-go-basics/code/001_hello_world.go
```

## Structure

- Root: the KV store (built incrementally per the README roadmap; stdlib only, no external
  dependencies for the core).
- `learn-go-basics/`: companion Go warm-up exercises that front-load the vocabulary each build
  step needs. **This directory has its own `CLAUDE.md` with strict rules — read it before
  touching anything there. In short: do not write or complete the exercise solutions; only
  guide, explain, review already-written code, and help with the progress log.** That
  restriction applies only inside `learn-go-basics/`, not to the store itself.

## Scope discipline

Resist scope creep. Out of scope for this phase: persistence, transactions, replication,
pub/sub, clustering, query languages. If a change reaches for one of these, stop and confirm.

## Writing style

Avoid emojis in documentation, comments, commit messages, and code. The existing docs are
deliberately terse and plain-text; keep that voice — no emojis, no marketing language, no
filler.
