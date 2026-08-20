# Bug reproduction

- Bug: the resolver graph is accessed concurrently without a synchronized snapshot and remote host policy allows unlisted hosts.
- Trigger: add and enumerate graph edges concurrently or evaluate a remote reference to an unlisted host.
- Error: panic: concurrent map read and map write, or the policy test reports that an unlisted host was allowed.
