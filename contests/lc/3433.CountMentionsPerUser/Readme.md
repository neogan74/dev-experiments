# [3433. Count Mentions Per User](https://leetcode.com/problems/count-mentions-per-user/description)

## Problem
We have `numberOfUsers` users (`0..numberOfUsers-1`) and an `events` list where each event is:
- `["MESSAGE", timestamp, mentionsString]`
- `["OFFLINE", timestamp, userId]`

Rules:
- `id<number>` mentions that specific user (duplicates allowed).
- `ALL` mentions every user (even offline).
- `HERE` mentions only users currently online.
- `OFFLINE` keeps a user offline for `60` time units; they return online at `timestamp + 60`.
- Status changes (going offline or coming back online) happen before any `MESSAGE` at the same timestamp.
- All users start online.

Return `mentions[i]` = total mentions for user `i`.

## Approach
1. Parse and sort events by timestamp; when timestamps tie, process `OFFLINE` events before `MESSAGE` so status changes are applied first.
2. Track per-user `online` state and `offlineUntil` timestamp (when they automatically come back).
3. Before handling any timestamp, refresh all users whose `offlineUntil <= currentTime` to online.
4. For an `OFFLINE` event, mark the user offline and set `offlineUntil = timestamp + 60`.
5. For a `MESSAGE`, split `mentionsString` by spaces and:
   - `ALL`: add `+1` to every user.
   - `HERE`: add `+1` to each currently online user.
   - `idX`: increment that specific user (duplicates naturally counted).

## Complexity
- Time: `O(U * E + M)` where `U` is users (`<=100`), `E` timestamps processed, and `M` total tokens across messages (`<=10^4` here). With small constraints, direct loops are fine.
- Space: `O(U)` for states and counts.
