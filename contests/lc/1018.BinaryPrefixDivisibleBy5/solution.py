from typing import List


def prefixesDivBy5(nums: List[int]) -> List[bool]:
    """
    Return a list where each element indicates whether the corresponding
    binary prefix of nums is divisible by 5.
    """
    result: List[bool] = []
    mod = 0

    for bit in nums:
        mod = ((mod << 1) + bit) % 5
        result.append(mod == 0)

    return result


if __name__ == "__main__":
    sample = [0, 1, 1]
    print(prefixesDivBy5(sample))
