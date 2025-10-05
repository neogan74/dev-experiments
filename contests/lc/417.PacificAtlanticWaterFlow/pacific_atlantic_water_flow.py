from collections import deque
from typing import List

def pacificAtlantic(heights: List[List[int]]) -> List[List[int]]:
    if not heights or not heights[0]:
        return []
    m, n = len(heights), len(heights[0])

    def bfs(starts):
        vis = [[False]*n for _ in range(m)]
        q = deque(starts)
        for r, c in starts:
            vis[r][c] = True
        while q:
            r, c = q.popleft()
            h = heights[r][c]
            for dr, dc in ((1,0),(-1,0),(0,1),(0,-1)):
                nr, nc = r+dr, c+dc
                if 0<=nr<m and 0<=nc<n and not vis[nr][nc] and heights[nr][nc] >= h:
                    vis[nr][nc] = True
                    q.append((nr, nc))
        return vis

    pac_starts = [(0, c) for c in range(n)] + [(r, 0) for r in range(m)]
    atl_starts = [(m-1, c) for c in range(n)] + [(r, n-1) for r in range(m)]
    pac = bfs(pac_starts)
    atl = bfs(atl_starts)

    res = []
    for r in range(m):
        for c in range(n):
            if pac[r][c] and atl[r][c]:
                res.append([r, c])
    return res

if __name__ == "__main__":
    heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
    res = pacificAtlantic(heights)
    print(f'For array:\n {heights}\n we have a result {res}')