# Bug reproduction

- Bug: redacted exports alias nested documents, empty exports succeed, and checksums ignore the document.
- Trigger: redact and mutate a nested metadata value, export an empty item set, or compare two documents.
- Error: the targeted tests report nested metadata aliasing, a successful empty export, or equal checksums.
