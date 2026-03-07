# Binary Search Patterns

| Pattern | Function | When to Use |
|---------|----------|-------------|
| Exact match | `Base` | Find a specific value in sorted array |
| Lower bound | `LowerBound` | Find first element >= target (leftmost insertion point) |
| Upper bound | `UpperBound` | Find first element > target (rightmost insertion point) |
| First true | `MonotonicPredicateFirstTrue` | Find transition point false→true in boolean array |
| Last true | `MonotonicPredicateLastTrue` | Find last true before false→true boundary |
| Generic predicate | `FirstTrue` | Smallest value in [lo, hi] satisfying a custom predicate |
| Continuous search | `DoSqrt` | Binary search on floats (e.g. square root with precision) |
