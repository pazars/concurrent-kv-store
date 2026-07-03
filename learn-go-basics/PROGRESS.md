## TODO
- Read https://go.dev/tour/list

## What I learned

**001**:
  - Basic notation (main pkg, main func, fmt)
  - Basic file run (go run *.go)
  - For loops

**002 — Vars, types, zero values**:
  - `var x T` gives the **zero value** (int→0, string→"", bool→false, slice→nil).
    No null-pointer surprises: every declared var is a valid value.
  - A **nil slice is usable**: len=0, rangeable, appendable. NOT like Python's None.
  - **Constants**: use `const`, never `:=`. Value must be known at compile time.
  - **iota**: auto-incrementing counter (0,1,2…) inside a `const(...)` block;
    blank lines repeat the previous expression → e.g. `1 << (10*iota)` = KB/MB/GB.
  - Gotchas that bit me:
    - `%b` on a bool → `%!b(bool=false)` (wrong verb). Use `%t` for bools.
    - Unused local vars / imports = **compile error**, not a warning.
  - Python contrast:
    - Zero values vs NameError/None; nil slice ≠ None.
    - **Capitalization = export**, compiler-enforced. `Foo`=public, `foo`=package-private.
      No export keyword, no ALL_CAPS constants — all names are camelCase.
    - fmt verbs: `%v` (any, default) / `%+v` (struct w/ field names) / `%q` (quoted str)
      / `%t` (bool). No `Printfln` — add `\n` yourself. `%%` = literal %.
    - Escape `\` via `"\\"` or backtick raw strings (like Python r"...").
    - `go vet` catches format-verb mismatches before running.

**003 — Functions & multiple returns**:
  - `(T, error)` return pattern: error is always last, `nil` means success;
    on failure, other return values are conventionally the zero value.
  - `nil` is predeclared, no import needed. Constructing a real error
    (`errors.New`, `fmt.Errorf`) needs `import "errors"` / `"fmt"`.
  - Named returns pre-declare vars in the signature (zero-valued at start);
    a naked `return` sends back their current values. Most useful for
    documenting multi-value returns like `(result T, err error)`.
  - **Divide by zero differs by type**: `int` division by zero **panics**
    at runtime (no way to get a value back — must guard before dividing).
    `float64` division by zero does **not** panic — IEEE-754 gives back
    `+Inf`/`-Inf`/`NaN` as an ordinary value. My zero-check is required
    for `int` (or the program crashes) but a *design choice* for `float64`
    (otherwise `Inf`/`NaN` silently propagates).
  - **panic**: runtime-triggered (nil deref, index out of range, int div/0,
    etc.) or explicit via `panic(v)`. Unwinds the stack, running `defer`s
    along the way, until `recover()` catches it (only callable inside a
    `defer`) or the program crashes with a stack trace. Reserved for bugs/
    broken invariants, not routine failure — that's what `error` is for.
  - **errcheck**: a linter (often via `golangci-lint`) that flags calls
    returning `error` where the return value is silently dropped.
  - Go only has `float32`/`float64` (no bare `float`); untyped float
    literals default to `float64`. No implicit conversion between them.
  - Style nit: error strings should be lowercase, no trailing punctuation
    (they often get wrapped into larger messages).
  - Python contrast:
    - No exceptions for routine errors — no invisible non-local jumps;
      failure is a value sitting right next to the result, checked inline.
    - Named returns have no Python equivalent (implicit pre-declared locals
      that a bare `return` sends back).
    - A Go panic is like an uncaught Python exception reaching the top and
      crashing the interpreter — except Go treats that as the *expected*
      default for bugs, not the mechanism for everyday error handling.
    - Python's single `float` is `float64` under the hood — Go makes you
      choose the width explicitly.

**004 — Slices**:
  - A slice is always a 3-field header: `{pointer, len, cap}` — even `nil`.
    `var s []int` → pointer nil, len 0, cap 0 (no backing array exists yet).
  - `make([]T, len)` allocates a backing array and zero-fills it, len==cap.
    `make([]T, len, cap)` allocates `cap` slots but only exposes `len` of them.
  - `append` returns a **new slice header** — must reassign (`s = append(s, x)`).
    Reuses the backing array if there's spare `cap`; otherwise allocates a new,
    bigger one, copies old elements over, zero-fills the rest.
  - Slicing `s[low:high]` aliases the **same backing array**, no copy — high
    is exclusive. The result's `cap` extends to the *original* array's end,
    not just `high`, so `s[low:high]` can have `cap` >> `len`. Slicing within
    `cap` but past `len` exposes real (zero-valued, if freshly grown) elements
    of the backing array; slicing past `cap` panics.
  - **Growth capacity is an implementation detail, not a language guarantee.**
    Tested this directly: the exact same operation (`append` one element onto
    a nil `[]int`) produced `cap==1` in one program and `cap==4` in another,
    same Go version, differing only in surrounding code. The only real
    contract is `cap(s) >= len(s)` after append — never hardcode assumptions
    about the growth factor or exact resulting capacity.
  - `range` gives `(index, value copy)` per iteration — `v` is a copy, so
    mutating `v` doesn't mutate `s[i]`. This copy is `O(1)` extra space
    regardless of slice size (one element reused each iteration, not an
    accumulating buffer) — unrelated to `append`'s growth/doubling, which is
    a completely separate mechanism that only triggers on `append` itself.
    For large struct elements, the per-iteration copy is a real CPU cost
    (not a memory-safety issue) — mitigate with `for i := range s` + `s[i]`,
    or a slice of pointers.
  - Python contrast:
    - Python's `list[a:b]` **copies**; Go's `s[low:high]` **aliases** the
      same backing array — mutations through one slice can be visible
      through another that shares the same backing storage.
    - CPython lists also over-allocate on append (amortized growth), but
      like Go, the exact factor is an implementation detail, not part of
      the language spec — same lesson, different language.
