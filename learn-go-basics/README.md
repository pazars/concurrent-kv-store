# learn-go-basics

Go warm-up exercises for a Python dev, building the fluency needed to work through
[`concurrent-kv-store`](../README.md). One small, standalone exercise per day
(`NNN_topic.go`), ~20–40 min each, focused on what's *different* in Go — static types,
pointers, explicit errors, interfaces, and the concurrency model — not programming
fundamentals.

## Layout

| Path | What it is |
|------|-----------|
| `ROADMAP.md` | The day-by-day curriculum (days 001–023), with **REPO** markers showing which parent-repo issue each cluster of days unlocks. |
| `PROGRESS.md` | Learning log: what was learned per day, gotchas, and Python contrasts. Grows as exercises are done. |
| `CLAUDE.md` | Ground rules for using Claude here — guide and review, don't write the exercise solutions. |
| `code/` | The exercises, one `NNN_topic.go` file per day. |

## Running an exercise

Each file is a self-contained `package main`, run directly — no module yet (a `go.mod`
arrives on day 010):

```
go run code/001_hello_world.go
go vet code/002_vars_types.go
```

From day 010 on, the exercises become `_test.go` files run under the project's core loop:

```
go test -race ./...
```

## Progress

Done: **001** hello world · **002** vars/types/zero values · **003** functions & errors ·
**004** slices. Next up: **005** maps (see `ROADMAP.md`). Each day's notes land in
`PROGRESS.md`.

## Why a separate directory

The parent repo's build order is the real curriculum; these exercises exist so each step
there feels trivial by the time it's reached. The roadmap ties them together — e.g. days
005/006/008/011 land before Issue #2, and days 012–014 (goroutines, the race detector,
`sync.Mutex`) land before the store is broken and refixed in Issues #3–#4.
