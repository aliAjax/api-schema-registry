# Bug reproduction

- Bug: zero-value asset storage can write through nil maps and service validation misses safety checks.
- Trigger: create and save a version on an unconstructed in-memory repository.
- Error: panic: assignment to entry in nil map.
