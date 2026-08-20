# Bug reproduction

- Bug: retry delivery ignores cancellation while waiting and can invoke a sender after cancellation.
- Trigger: cancel a delivery context during retry backoff or before the first attempt.
- Error: the retry tests fail with a delivery error after extra attempts instead of the context error.
