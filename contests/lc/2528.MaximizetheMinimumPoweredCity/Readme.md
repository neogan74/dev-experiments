# [2528. Maximize the Minimum Powered City](https://leetcode.com/problems/maximize-the-minimum-powered-city/description)

You are given a 0-indexed integer array `stations` of length `n`, where `stations[i]` represents the number of power stations in the `i`-th city.

Each power station can provide power to every city in a fixed range. If the range is denoted by `r`, then a power station at city `i` can provide power to all cities `j` such that `|i - j| <= r` and `0 <= i, j <= n - 1`. The power of a city is the total number of power stations that can reach it.

The government will build `k` additional power stations, each with the same range `r`. You may distribute these new stations across any cities. Return the maximum possible value of the minimum city power after optimally placing the `k` stations.

---

### Example 1

```
Input:  stations = [1, 2, 4, 5, 0], r = 1, k = 2
Output: 5
```

One optimal placement adds both new stations to city `1`, resulting in `stations = [1, 4, 4, 5, 0]`.

- City 0 power: `1 + 4 = 5`
- City 1 power: `1 + 4 + 4 = 9`
- City 2 power: `4 + 4 + 5 = 13`
- City 3 power: `5 + 4 = 9`
- City 4 power: `5 + 0 = 5`

The minimum city power is `5`, which is optimal.

### Example 2

```
Input:  stations = [4, 4, 4, 4], r = 0, k = 3
Output: 4
```

With `r = 0`, each station serves only its own city. No placement of three new stations can raise the minimum power beyond `4`.

---

### Constraints

- `n == stations.length`
- `1 <= n <= 10^5`
- `0 <= stations[i] <= 10^5`
- `0 <= r <= n - 1`
- `0 <= k <= 10^9`
