# Algorithm Explanation

## Overview
We must tally how many times each user is mentioned across a sequence of timestamped events. Users can be temporarily offline for exactly 60 time units; coming back online happens automatically. A `MESSAGE` event can mention all users (`ALL`), only currently online users (`HERE`), or explicit user ids (`idX`). Status changes are applied before any message at the same timestamp, and all users start online.

## Event Ordering
1. Parse all events as `(type, timestamp, payload)`.
2. Sort by `(timestamp, typePriority)` where `OFFLINE` has higher priority (0) than `MESSAGE` (1). This enforces "status changes first" when timestamps tie.

## State Tracking
- `online[i]`: whether user `i` is currently online (start `True`).
- `offlineUntil[i]`: timestamp when an offline user returns online (meaning online again at `time >= offlineUntil[i]`).
- `mentions[i]`: count of mentions for user `i`.

## Processing Loop
For each sorted event:
1. **Refresh Online Status**: if the timestamp changed, bring back any users with `now >= offlineUntil[i]` (they auto-return).
2. **OFFLINE**: mark the user offline and set `offlineUntil = timestamp + 60`.
3. **MESSAGE**: split the payload by spaces and handle each token:
   - `ALL`: increment every `mentions[i]`.
   - `HERE`: increment `mentions[i]` for users currently online.
   - `idX`: parse `X`; if valid, increment `mentions[X]` (duplicates count naturally).

## Correctness Proof (Sketch)
We show that the algorithm produces correct mention counts.
1. **Ordering**: Sorting by timestamp and processing `OFFLINE` before `MESSAGE` at equal timestamps matches the required rule, so status changes are applied before any message at that time.
2. **Status Refresh**: Before handling events at a new timestamp `t`, all users with `offlineUntil <= t` are set online, exactly matching the auto-return rule.
3. **OFFLINE Handling**: Setting `online=false` and `offlineUntil = t+60` replicates the 60-unit offline duration starting at the event time.
4. **MESSAGE Handling**:
   - `ALL` increments every user, including offline ones, matching the spec.
   - `HERE` increments only users with `online=true`, matching "online users only".
   - `idX` increments only the specified user; duplicates in the payload produce multiple increments, matching "each mention counts".
5. By 1–4, each message contributes exactly the specified mentions to the correct users under the correct online/offline statuses. Summing these increments yields the true mention counts. Therefore, the algorithm is correct.

## Complexity Analysis
- Let `U` be `numberOfUsers (<=100)`, `E` be number of events (`<=100`), and `T` be total tokens across messages (`<=10^4` here).
- Sorting: `O(E log E)` (small, negligible).
- Per event work: `O(U)` for refresh and `O(U)` for `ALL`/`HERE`, plus `O(T)` total for explicit ids.
- Overall: `O(E log E + U*E + T)` time, `O(U)` space.

## Reference Implementations
- Go: `solution.go`
- Python: `solution.py`
