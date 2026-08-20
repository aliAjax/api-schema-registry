# Bug reproduction

- Bug: publishing does not complete the asset state transition, history, timestamp, or published query classification.
- Trigger: apply a ready publication and then read the stored version and published list.
- Error: targeted tests report a candidate state, missing history or timestamp, or an empty published result.
