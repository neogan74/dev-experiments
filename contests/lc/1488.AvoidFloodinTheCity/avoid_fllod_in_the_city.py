
# lc_1488_avoid_flood.py
from bisect import bisect_left
from typing import List

def avoidFlood(rains: List[int]) -> List[int]:
    n = len(rains)
    ans = [1] * n
    last = {}        # lake -> last rainy day index
    dry = []         # sorted indices of dry days (rains[i]==0)

    for i, x in enumerate(rains):
        if x == 0:
            dry.append(i)   # просто запишем индекс сухого дня; список отсортирован по построению
        else:
            ans[i] = -1
            if x in last:
                prev = last[x]
                # найти первый сухой день > prev
                jpos = bisect_left(dry, prev + 1)
                if jpos == len(dry):
                    return []  # нет подходящего сухого дня -> невозможно
                j = dry[jpos]
                ans[j] = x     # в этот сухой день осушаем озеро x
                dry.pop(jpos)  # этот сухой день использован
            last[x] = i
    # оставшиеся сухие дни уже стоят как 1 — это валидно
    return ans
