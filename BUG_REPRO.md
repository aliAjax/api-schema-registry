# Bug reproduction

- Bug: import rollback runs on the wrong transaction path and manifests never become ready.
- Trigger: make commit fail, cancel before commit, or build a complete manifest.
- Error: rollback tests report residual data or an incorrectly undone successful import; readiness is false.
