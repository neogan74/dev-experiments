from pacific_atlantic_water_flow import pacificAtlantic

def toset(ans):
    return {tuple(p) for p in ans}

def test_examples():
    heights = [
        [1,2,2,3,5],
        [3,2,3,4,4],
        [2,4,5,3,1],
        [6,7,1,4,5],
        [5,1,1,2,4],
    ]
    exp = {(0,4),(1,3),(1,4),(2,2),(3,0),(3,1),(4,0)}
    assert toset(pacificAtlantic(heights)) == exp

def test_single_row_col_and_edge():
    assert toset(pacificAtlantic([[1,2,3]])) == {(0,0),(0,1),(0,2)}
    assert toset(pacificAtlantic([[3],[2],[1]])) == {(0,0),(1,0),(2,0)}
    assert pacificAtlantic([]) == []