---
name: tests_for_spec_not_implementation
description: When writing tests, always test against the problem specification, not the current implementation
type: feedback
---

Always write tests that verify the problem specification / task requirements, not the current implementation.

**Why:** The user uses tests to validate correctness against the problem statement. Tests matching a buggy implementation defeat the purpose.

**How to apply:** When writing tests, derive expected values from the problem description (e.g., 1-based indices if the spec says so), even if the current code returns something different.
