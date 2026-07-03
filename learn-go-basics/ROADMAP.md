# Go Roadmap → `concurrent-kv-store`

Goal: build the Go vocabulary needed to work through the issues in
`pazars/concurrent-kv-store`, one coffee-sized warm-up per day.

**How to use this:**
- One warm-up per morning. Each is a small standalone `NNN_topic.go` (or `_test.go`)
  like `001_hello_world.go`. Target ~20–40 min. Double up if you're caffeinated.
- "Look up" = find the API yourself and try it; don't paste a full solution.
- Repo work is interleaved: a **REPO** line means "you now know enough — go do that issue."
- Log each day in `progress.md` (template at the bottom).

You're a pro Python dev, so this skips programming fundamentals and focuses on
what's *different* in Go: static types, pointers, explicit errors, interfaces,
and the concurrency model.

---

## Phase 1 — Go syntax & data structures (foundation)

These have no concurrency. They make Step 0/1 of the repo feel trivial.

| Day | Topic | Look up | Tiny exercise | Done when |
|-----|-------|---------|---------------|-----------|
| 002 | Vars, types, zero values | `var` vs `:=`, typed constants, zero values, `const`/`iota` | Declare an int/string/bool/slice without initialising; print their zero values. Add one typed const. | You can predict each zero value before running. |
| 003 | Functions & multiple returns | named returns, `(T, error)` return pattern | `divide(a, b int) (int, error)` returning an error on divide-by-zero; call it both ways. | Caller handles the error branch. |
| 004 | Slices | `make`, `append`, `len`/`cap`, slicing `s[1:3]`, `range` | Build a slice with `append` in a loop; print it; iterate with `range` (index + value). | You understand why `cap` can exceed `len`. |
| 005 | Maps | `make(map[K]V)`, comma-ok `v, ok := m[k]`, `delete`, nil-map panic | Make a `map[string]int`, do a hit + a miss with comma-ok, `delete` a key, range over it. | You can explain why writing to a `nil` map panics. |
| 006 | Structs & methods | struct literal, `func (s *T) M()` vs `func (s T) M()`, constructor `NewT()` | A `Counter` struct with `Inc()` (pointer receiver) and `Value()` (value receiver); a `NewCounter()`. | `Inc` actually mutates — proves you grasp pointer receivers. |
| 007 | Pointers | `&x`, `*p`, when Go auto-derefs, nil pointers | Write a func that takes `*int` and mutates it; call it; show value-receiver vs pointer-receiver difference. | You know why 006's `Inc` needed `*Counter`. |
| 008 | Idiomatic errors | `errors.New`, `fmt.Errorf("...: %w", err)`, sentinel `var ErrX = errors.New`, `errors.Is` | Define `ErrNotFound`; a func returns it; caller checks with `errors.Is`. | You stop reaching for exceptions-style flow. |
| 009 | Interfaces | implicit satisfaction, small interfaces, `fmt.Stringer`, `any` | Define a `Stringer`-style interface; make a struct satisfy it; pass it to a func taking the interface. | No `implements` keyword needed — it just works. |

## Phase 2 — Testing (the repo's core loop)

The whole repo runs on `go test -race`. Own the loop before Step 0.

| Day | Topic | Look up | Tiny exercise | Done when |
|-----|-------|---------|---------------|-----------|
| 010 | `go mod` + first test | `go mod init`, `_test.go`, `func TestXxx(t *testing.T)`, `t.Errorf` vs `t.Fatalf`, `go vet` | New module; one test asserting `1+1==2`; run `go test ./...` and `go vet ./...`. | Both pass. You know Error (continue) vs Fatal (stop). |
| 011 | Table-driven tests | slice of case structs, `t.Run(name, ...)` subtests | Test 003's `divide` with a table of cases incl. the error case. | Subtests show up named in `-v` output. |

**REPO — Issue #1 (Step 0):** project skeleton + first test. You're fully equipped.
**REPO — Issue #2 (Step 1):** single-threaded `Store` over `map[string]string`
   (uses days 005, 006, 008, 011 directly). Table-driven tests for every method.

## Phase 3 — Concurrency primitives (the heart of the project)

| Day | Topic | Look up | Tiny exercise | Done when |
|-----|-------|---------|---------------|-----------|
| 012 | Goroutines + WaitGroup | `go func(){}()`, `sync.WaitGroup` Add/Done/Wait, **loop-variable capture** | Launch 5 goroutines that print their index via WaitGroup. Then reproduce the capture bug and fix it. | You know why `Add` goes *before* the `go`, and the capture gotcha. |
| 013 | The race detector | what `-race` instruments, non-determinism | Shared `int` counter incremented by 100 goroutines (no lock); run `go test -race`; capture the report. | You've seen a `DATA RACE` report on your own code. |

**REPO — Issue #3 (Step 2):** break the store on purpose, document the race. **Do not fix it.**

| Day | Topic | Look up | Tiny exercise | Done when |
|-----|-------|---------|---------------|-----------|
| 014 | `sync.Mutex` | `Lock`/`Unlock`, `defer mu.Unlock()`, embedding a mutex in a struct | Fix day 013's counter with a `sync.Mutex`; `-race` now clean. | Final count is exactly N×M, race-free. |

**REPO — Issue #4 (Step 3):** fix the store with one global mutex.

| Day | Topic | Look up | Tiny exercise | Done when |
|-----|-------|---------|---------------|-----------|
| 015 | Benchmarks | `func BenchmarkXxx(b *testing.B)`, `b.N`, `b.RunParallel`+`pb.Next()`, `b.ResetTimer`, `-bench -benchmem`, per-goroutine `rand` | Benchmark a trivial mutex-counter under `RunParallel`; read ns/op. | You can convert ns/op → ops/sec. |

**REPO — Issue #5 (Step 4):** benchmark harness + record the global-mutex **baseline**.

| Day | Topic | Look up | Tiny exercise | Done when |
|-----|-------|---------|---------------|-----------|
| 016 | `sync.RWMutex` | `RLock`/`RUnlock` vs `Lock`, when readers contend | Swap day 014's mutex for RWMutex; benchmark read-heavy vs write-heavy. | You see RWMutex win on read-heavy, not write-heavy. |

**REPO — Issue #6 (Step 5):** RWMutex store; re-run benchmark.

| Day | Topic | Look up | Tiny exercise | Done when |
|-----|-------|---------|---------------|-----------|
| 017 | Hashing & bit masking | `hash/fnv` (fnv-1a), `% n` vs `& (n-1)`, power-of-two, `runtime.NumCPU` | Hash 1000 keys into N=16 buckets with fnv-1a + masking; print the per-bucket counts to eyeball distribution. | You can argue why power-of-two lets you mask instead of mod. |

**REPO — Issue #7 (Step 6):** sharded locking. Now all three strategies are comparable — write the comparison.

## Phase 4 — Networking (turn the store into a server)

| Day | Topic | Look up | Tiny exercise | Done when |
|-----|-------|---------|---------------|-----------|
| 018 | strings + strconv | `strings.Fields`, `strings.SplitN`, `strconv.Atoi`/`Itoa` | Parse `"SET foo bar baz"` into cmd + key + value-with-spaces using `SplitN`. | You know why `SplitN` beats `Fields` for `SET key value with spaces`. |
| 019 | bufio line reading | `bufio.Scanner` vs `bufio.Reader.ReadString('\n')`, `io.EOF`, Scanner line limits | Read a multi-line `strings.Reader` line by line; stop cleanly at EOF. | You handle EOF as a normal end, not an error. |
| 020 | TCP server | `net.Listen`, `ln.Accept()`, `go handleConn(conn)`, `defer conn.Close()` | A TCP echo server; test with `nc localhost PORT`. | Two `nc` sessions echo concurrently. |

**REPO — Issue #8 (Step 7):** TCP server in front of the store (`GET`/`SET`/`DELETE`, newline-delimited).

## Phase 5 — Atomic features & shutdown

| Day | Topic | Look up | Tiny exercise | Done when |
|-----|-------|---------|---------------|-----------|
| 021 | Channels + select | unbuffered vs buffered, `select`, `close`, done-channel pattern | A goroutine that loops on `select` over a work channel + a `done` channel; close `done` to stop it. | Goroutine exits cleanly; no leak. |
| 022 | context + time.Ticker | `context.WithCancel`, `ctx.Done()`, `time.Ticker`, `time.Now`/`Duration`, comparing deadlines | A ticker-driven loop that prints every 100ms and stops on `ctx` cancel. | The loop stops within one tick of cancel. |
| 023 | sync/atomic | `atomic.Int64`, `Add`, `Load`, `CompareAndSwap` | Re-do the counter with `atomic.Int64` (no mutex); benchmark vs the mutex version. | You see lock-free beat mutex on the hot path. |

**REPO — Issue #9 (INCR/DECR):** demonstrate the lost-update bug, then fix with one critical section.
   (Day 023's `CompareAndSwap` + the read-modify-write lesson.)
**REPO — Issue #10 (CAS):** compare-and-swap command.
**REPO — Issue #11 (TTL/expiry):** value-struct + lazy expiry + a sweeper goroutine
   (days 021–022: ticker + context shutdown, no goroutine leak).

## Then: deliverables & stretch

- **Issues #12–#14** (benchmark table, race-clean CI, design writeup) are *writing/wiring*, not new Go.
  For CI: look up GitHub Actions + `go test -race ./...` in a workflow file.
- **Issue #15** (RESP protocol) — extends day 019/020 parsing.
- **Issue #16** (lock-free reads) — extends day 023 atomics + copy-on-write.

---

## Daily log template (append to `progress.md`)

```
**NNN — <topic>**:
  - <thing 1 you learned>
  - <thing 2 / a gotcha that bit you>
  - Python contrast: <how it differs from how you'd do it in Python>
```

The "Python contrast" line is worth keeping — naming the difference is what makes
the Go idiom stick.
