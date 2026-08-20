# Bug reproduction

- Bug: concurrent event-log access can lose entries or reuse sequence numbers.
- Trigger: append and read changes concurrently during a publish burst.
- Error: the race-enabled event-log test reports a data race and inconsistent event state.
