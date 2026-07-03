# Concurrent KV Store

A learning project: a **thread-safe in-memory key-value store in Go**, built to understand
practical concurrency hands-on. This is **Phase 2** of a backend learning roadmap.

The point is **not** to ship a product. It's to understand concurrency mechanisms by building
the *same* store under progressively better synchronization strategies — and **measuring the
difference** between them.

## How this project works (read this first)

The **build order matters more than the end result.** Each step must run and pass its tests
under the race detector before starting the next:

```
go test -race ./...
```

Don't jump ahead to a finished, fully-sharded, networked implementation. The whole point is to
*feel* why each layer of synchronization is needed by hitting the problem it solves first.

## Roadmap (do in order)

| # | Step | Issue |
|---|------|-------|
| 0 | Project skeleton + first test | [#1](../../issues/1) |
| 1 | Single-threaded store | [#2](../../issues/2) |
| 2 | Break it on purpose — see the data race | [#3](../../issues/3) |
| 3 | Fix with one global mutex | [#4](../../issues/4) |
| 4 | Benchmark harness (baseline) | [#5](../../issues/5) |
| 5 | Read-write lock (RWMutex) | [#6](../../issues/6) |
| 6 | Sharded locking | [#7](../../issues/7) |
| 7 | TCP server | [#8](../../issues/8) |

### Atomic-operation features (after Step 7)
| Feature | Issue |
|---------|-------|
| `INCR` / `DECR` — atomic read-modify-write | [#9](../../issues/9) |
| `CAS key old new` — compare-and-swap | [#10](../../issues/10) |
| TTL / expiry (`SET key value EX 30`) + background sweeper | [#11](../../issues/11) |

### CV-worthy deliverables
| Deliverable | Issue |
|-------------|-------|
| Benchmark comparison of the three locking strategies | [#12](../../issues/12) |
| Race-detector-clean test suite in CI | [#13](../../issues/13) |
| Design writeup (why sharded won, when to pick each, 100× scale) | [#14](../../issues/14) |

### Stretch goals (only after the core is solid)
| Goal | Issue |
|------|-------|
| RESP (Redis) wire protocol + `redis-benchmark` | [#15](../../issues/15) |
| Lock-free reads (atomics / copy-on-write) | [#16](../../issues/16) |

## Learning the Go first

This repo is Phase 2 of a backend roadmap, tackled by a Python dev learning Go. The
[`learn-go-basics/`](learn-go-basics/) directory holds the day-by-day Go warm-ups that
front-load the vocabulary each step here needs (slices, maps, `sync.Mutex`, benchmarks,
TCP, `sync/atomic`). Its roadmap is mapped issue-by-issue onto the steps below — see
[`learn-go-basics/README.md`](learn-go-basics/README.md).

## Tech

- **Language:** Go — chosen for cheap goroutines and the built-in race detector.
- **No external dependencies** for the core store; standard library only.
- Always run tests under the race detector: `go test -race ./...`

## Out of scope for Phase 2 (resist scope creep)

- Persistence to disk → Phase 3 (write-ahead log)
- Transactions / isolation levels → Phase 3
- Replication / distribution → Phase 5
- Pub/sub, clustering, query language → never (not the point)
