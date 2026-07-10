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

**005 — Maps**:
  - Comma-ok `v, ok := m[k]`: a plain `v := m[k]` on a missing key doesn't
    panic or error, it returns the value type's **zero value** — which is
    often a legitimate stored value too (e.g. `0` for `int`), so `ok` is
    the only way to tell "present with zero value" from "absent".
  - `delete(m, k)` removes a key; it's a no-op (not an error) if the key
    isn't there, and it's also safe to call on a `nil` map.
  - **nil-map asymmetry**: reading from a `nil` map is fine (`len` 0, comma-ok
    reports `false`, `range` does nothing) — but *writing* (`m[k] = v`)
    panics: `assignment to entry in nil map`. Reads can be defined as
    "always miss" against nothing; a write needs a hash table to insert
    into, and unlike slices there's no auto-allocate-on-write for maps.
  - Why nil maps are allowed at all rather than banned: it keeps the
    "every zero value is valid" guarantee (readable without `make`), lets
    nil vs `make(...)`-but-empty be a meaningful distinction (unconfigured
    vs configured-but-empty), and avoids forcing an allocation for map
    variables that never get used.
  - `range` over a map yields `k, v` pairs; drop one side with `_` if you
    only need keys or values. Iteration order is **deliberately randomized**
    per run — need sorted output, collect + sort the keys yourself.
  - `:=` requires at least one new variable on the left side, so reusing
    `v, ok` for a second lookup in the same scope is a compile error —
    that's why repeated lookups need distinct names (`v2, ok2`, etc.),
    not a style choice.
  - **New tool: `gofmt`**. Formats a file in place (`gofmt -w file.go`) or
    just lists files that need it (`gofmt -l .`). Catches things like
    stray trailing whitespace and inconsistent indentation automatically —
    worth running before considering any file done, same reflex as `go vet`.
  - Python contrast:
    - Go's comma-ok is the `m[k]` (raises `KeyError`) vs `m.get(k)`
      (returns `None`) split, but collapsed into one lookup expression —
      no exception, no `None`, just a second boolean return.
    - Python dicts (3.7+) guarantee insertion order; Go map order is
      explicitly *not* guaranteed and is randomized to stop people from
      relying on it.

**006 — Structs & methods**:
  - Struct bodies hold **fields only** (`name type` pairs) — methods and
    constructors are never nested inside `type T struct {...}`. They're
    separate top-level `func` declarations, distinguished only by the
    receiver clause: `func (c *Counter) Inc()`.
  - **Pointer receiver vs value receiver** is the whole point: a value
    receiver gets a *copy* of the struct, so mutations inside the method
    are invisible to the caller — only a pointer receiver can actually
    change the caller's struct. First-hand proof: `Inc()` needed
    `*Counter` to make the increment stick across calls.
  - Convention once a type has any pointer-receiver method: give **all**
    its methods a pointer receiver, to avoid mixing receiver kinds on the
    same type.
  - Constructors are just an ordinary func by convention (`NewT`), not a
    language feature — no `__init__`. Capitalized like any other exported
    name; caught myself defaulting to lowercase `newCounter` out of Python
    habit before fixing it to `NewCounter`.
  - **Auto-`&`/auto-`*` at call sites, and the asymmetry between them**:
    - value → pointer-receiver method: Go rewrites `c.Inc()` to
      `(&c).Inc()` automatically, but only if `c` is *addressable* (a
      plain variable is; a literal like `Counter{}.Inc()` is not and
      won't compile).
    - pointer → value-receiver method: Go rewrites `counter.Value()` to
      `(*counter).Value()` automatically, and this one always compiles —
      no addressability check — which pushes the risk to runtime: if
      `counter` were `nil`, it'd panic on the dereference instead of
      failing to compile.
  - Python contrast:
    - No `__init__`/`self` magic — receivers are just a syntactic marker
      on an otherwise ordinary function, and "constructor" is a naming
      convention (`NewT`), not a reserved method name.
    - Python has no copy-vs-reference choice per method — `self` is
      always a reference. Go makes you pick, explicitly, per method.

**007 — Pointers**:
  - `&x` takes the address of `x` (→ `*T`); `*p` dereferences — reads through
    the pointer, or writes through it on the left of `=`. The whole mechanism
    is just address-of / go-to-that-address, in either direction.
  - Nil pointers: zero value of any pointer type is `nil`; dereferencing one
    panics at runtime (`invalid memory address or nil pointer dereference`),
    no compile-time check.
  - **Methods can only be defined on named types declared in the same
    package** — `int`/`*int` aren't eligible receivers, even though `*Counter`
    (day 006) was fine. First draft tried `func (s *int) pointInc()` and hit
    this directly; the fix wasn't a syntax tweak, it was switching to plain
    functions (`func PointInc(val *int)`) since the exercise wanted a
    function taking a pointer param, not a method.
  - **`++`/`--` are statements, not expressions** — they can never appear
    where a value is expected (`return s++`, `&s++`, `f(s++)` are all
    invalid). Go deliberately kept them statement-only, unlike C where
    `x++` evaluates to a value and can be composed/abused. General test for
    statement vs expression: can it go on the right of `=`, as a `return`
    value, or as an operand to `&`/`*`? If not, it's a statement.
  - `*val++` parses as `(*val)++` (dereference, then increment the pointee),
    not `*(val++)`. Confirmed by running it. Reasoning: `++` is always a
    suffix on top of a complete preceding expression, never an inner piece —
    so `*val` is built first as an ordinary unary-`*` expression, then `++`
    applies to that whole thing. Also `val++` alone (incrementing the
    pointer itself) wouldn't compile regardless — Go has no pointer
    arithmetic outside `unsafe`, so that parse wasn't even possible here.
  - Smaller bugs hit along the way: method name called with different
    casing than declared (`PointInc` vs `pointInc` — Go is case-sensitive,
    no fuzzy matching); assigning to a not-yet-declared variable with `=`
    instead of `:=` (same "at least one new var" rule from day 005's
    comma-ok, just via plain assignment this time instead of multi-assign).
  - Python contrast:
    - Python has no pointers in the C/Go sense — every name is already a
      reference to its object, so there's no separate `&`/`*` step and no
      copy-vs-pointer choice to make per call; Go's pointer receiver vs
      value receiver (day 006) only exists *because* Go defaults to copying.
    - Python has no `++`/`--` at all (`x += 1` only) — so the
      "statement, not expression" trap doesn't come up the same way, but
      the general lesson (assignment is a statement, not usable as a value)
      is the same one Python enforces even more strictly (`y = (x = 5)` is
      also illegal there).
