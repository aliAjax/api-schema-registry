# Bug reproduction

- Bug: parser sentinel errors are formatted away across import layers, and invalid documents can be accepted.
- Trigger: validate an oversized, wrong-dialect, or parser-error document through the importer.
- Error: the targeted tests report a missing parser sentinel or that the invalid document was accepted.
