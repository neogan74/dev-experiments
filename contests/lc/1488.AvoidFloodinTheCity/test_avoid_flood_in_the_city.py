# test_lc_1488_avoid_flood.py
from avoid_fllod_in_the_city import avoidFlood

def valid_plan(rains, ans):
    full = set()
    for i, x in enumerate(rains):
        if x > 0:
            if x in full:
                return False
            full.add(x)
        else:
            lake = ans[i]
            if lake != -1 and lake in full:
                full.remove(lake)
            # lake==1 (или другое ненужное) допустимо, если этого озера нет в full
    return True

def test_examples():
    assert avoidFlood([1,2,3,4]) == [-1,-1,-1,-1]
    res = avoidFlood([1,2,0,0,2,1])
    assert res and valid_plan([1,2,0,0,2,1], res)
    res = avoidFlood([1,2,0,1,2])
    assert res == []

def test_edge_cases():
    assert avoidFlood([0,0]) in ([1,1],[1,42],[42,1])  # любые числа ок
    res = avoidFlood([1,0,2,0,1,2])
    assert res and valid_plan([1,0,2,0,1,2], res)
