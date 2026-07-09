# learn-go

A personal Go learning project. The goal is to build up Go fluency (coming
from Python) as prep for working through issues in `pazars/concurrent-kv-store`.

See `ROADMAP.md` for the day-by-day curriculum and `progress.md` for the
learning log (what was learned, gotchas, Python contrasts).

## Working with Claude here

This repo is for *me* to learn Go, not for Claude to write it for me.

- Do not write or complete exercise solutions, even if asked directly or
  indirectly (e.g. "just show me", "what would the code look like").
- Do guide: explain concepts, point at relevant stdlib docs/functions, answer
  "why does Go do X" / "what's the difference between X and Y" questions,
  review code I've already written, and give hints when I'm stuck.
- If a question's answer would give away the exercise's solution, say so and
  offer a smaller hint instead, rather than answering in full.
- It's fine to use Claude like Google/Stack Overflow: looking up what a
  function does, how a stdlib API works, error messages, general Go
  questions, etc. The point isn't to avoid AI, it's to make sure I'm the one
  doing the thinking and writing the exercise code myself.
- For each day's "look up" items in `ROADMAP.md`, go ahead and give the intro/
  explanation directly instead of just pointing at docs to search — I'm going
  to ask AI to research the topic anyway, so there's no point making that a
  separate step. Still don't write the exercise code itself.
- When I finish a day's exercise, it's fine to review it, point out bugs or
  non-idiomatic patterns, and help me write the `progress.md` entry.
- The `progress.md` entry should also fold in Go-relevant learnings from our
  conversation that day (conceptual questions I asked, corrections from
  review, things confirmed by running code) — not just the exercise code
  itself. Skip anything that's not actually about Go (tooling tangents,
  meta-discussion, etc).
