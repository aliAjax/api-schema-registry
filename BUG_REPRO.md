# Bug reproduction

- Bug: invalid schema regular expressions panic during validation instead of becoming a rule violation.
- Trigger: validate a schema containing a malformed pattern such as `[`. 
- Error: panic: regexp: Compile(`[`) : error parsing regexp: missing closing ]: `[`. 
