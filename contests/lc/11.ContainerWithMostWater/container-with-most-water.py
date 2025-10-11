from typing import List

def MaxArea(height: List[int]) -> int:
    l, r = 0, len(height) - 1
    best = 0
    while l < r:
        h = height[l] if height[l] < height[r] else height[r]
        area = h * (r -l)
        if area > best:
            best = area
        if height[l] < height[r]:
            l += 1
        else:
            r -= 1
    return best

if __name__ == '__main__':
    heigh_list = [1,8,6,2,5,4,8,3,7]
    heigh_list2 = [1,1]
    print(f'Result for height list {heigh_list} is {MaxArea(heigh_list)}')
    print(f'Result for height list {heigh_list2} is {MaxArea(heigh_list2)}')