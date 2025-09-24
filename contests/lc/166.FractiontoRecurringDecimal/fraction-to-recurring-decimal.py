class Solution:
    def fractionToDecimal(self, numerator: int, denominator: int) -> str:
        if numerator == 0:
            return "0"
        ans = []
        neg = (numerator > 0) ^ (denominator > 0)
        if neg:
            ans.append("-")
        a, b = abs(numerator), abs(denominator)
        ans.append(str(a // b))
        a %= b
        if a == 0:
            return "".join(ans)
        ans.append(".")
        d = {}
        while a:
            d[a] = len(ans)
            a *= 10
            ans.append(str(a // b))
            a %= b
            if a in d:
                ans.insert(d[a], "(")
                ans.append(")")
                break
        return "".join(ans)
    
def main():
    examples = (
        (1, 2, "0.5"),
        (2, 1, "2"),
        (4, 333, "0.(012)"),
        (1, 6, "0.1(6)"),
        (-50, 8,"-6.25"),
        (7, -12, "-0.58(3)"),
        (1, 333, "0.(003)"),
        (1, 5, "0.2"),
        (0, -5, "0"),
        (1, -2147483648, "-0.0000000004656612873077392578125"),
    )
    for numerator, denominator, expected in examples:
        res = (f"{numerator}/{denominator} = {Solution().fractionToDecimal(numerator, denominator)}")
        # print(res)
        res2 = res.split(" = ")
        if res2[1] == expected:
            print(f"{numerator}/{denominator} = {res2[1]} PASS")
        else:
            print(f"  FAIL: expected {expected}")
if __name__ == "__main__":
    main()