from typing import List
from collections import Counter
import bisect

def maximumTotalDamage(power: List[int]) -> int:
    cnt = Counter(power)
    vals = sorted(cnt)                             # уникальные v по возрастанию
    gains = [v * cnt[v] for v in vals]            # суммарная "выгода" на этом v

    n = len(vals)
    dp = [0] * (n + 1)                            # dp[i] — ответ на префиксе [0..i-1]
    for i in range(1, n + 1):
        v = vals[i - 1]
        # найдём последний индекс j (< i-1), где vals[j] <= v-3
        j = bisect.bisect_right(vals, v - 3)      # даёт первый > v-3
        j_dp = dp[j]                               # уже готов на префиксе [0..j-1]
        dp[i] = max(dp[i - 1], j_dp + gains[i - 1])
    return dp[n]