import re
from typing import List


class Solution:
    def validCouponCodes(
        self, code: List[str], businessLine: List[str], isActive: List[bool]
    ) -> List[str]:
        order = {
            "electronics": 0,
            "grocery": 1,
            "pharmacy": 2,
            "restaurant": 3,
        }
        valid_code = re.compile(r"^[A-Za-z0-9_]+$")

        coupons = []
        for c, b, active in zip(code, businessLine, isActive):
            if not active:
                continue
            if b not in order:
                continue
            if not c or not valid_code.fullmatch(c):
                continue
            coupons.append((order[b], c))

        coupons.sort(key=lambda x: (x[0], x[1]))
        return [c for _, c in coupons]
