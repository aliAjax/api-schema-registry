# Bug reproduction

- Bug: resolver reads ignore cancellation, return aliased cached bytes, and accept traversal references.
- Trigger: cancel a resolver operation, mutate returned cached bytes, or load a parent-directory reference.
- Error: resolver tests report a nil cancellation error, changed stored bytes, or accepted traversal.
