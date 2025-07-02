package findtheoriginaltypedstringii

const MOD = 1_000_000_007

type ModInt uint32

func NewModInt(x int) ModInt {
	return ModInt(x % MOD)
}

func (a ModInt) Add(b ModInt) ModInt {
	res := int(a) + int(b)
	if res >= MOD {
		res -= MOD
	}
	return ModInt(res)
}

func (a ModInt) Sub(b ModInt) ModInt {
	res := int(a) - int(b)
	if res < 0 {
		res += MOD
	}
	return ModInt(res)
}

func (a ModInt) Mul(b ModInt) ModInt {
	return ModInt((uint64(a) * uint64(b)) % MOD)
}

func (a ModInt) Pow(exp int) ModInt {
	result := ModInt(1)
	base := a
	for exp > 0 {
		if exp%2 == 1 {
			result = result.Mul(base)
		}
		base = base.Mul(base)
		exp /= 2
	}
	return result
}

func (a ModInt) Inv() ModInt {
	return a.Pow(MOD - 2)
}

func (a ModInt) Div(b ModInt) ModInt {
	return a.Mul(b.Inv())
}

func possibleStringCount(word string, k int) int {
	s := []byte(word)
	n := len(s)

	var cnt []int
	for i := 0; i < n; {
		j := i
		for j < n && s[i] == s[j] {
			j++
		}
		cnt = append(cnt, j-i)
		i = j
	}

	ans := NewModInt(1)
	for _, c := range cnt {
		ans = ans.Mul(NewModInt(c))
	}

	if len(cnt) < k {
		sum := k - len(cnt) - 1
		dp := make([]ModInt, sum+1)
		pref := make([]ModInt, sum+2)
		dp[0] = NewModInt(1)

		for _, c := range cnt {
			c -= 1
			pref[0] = NewModInt(0)
			for i := 0; i <= sum; i++ {
				pref[i+1] = pref[i].Add(dp[i])
			}
			for s := sum; s >= 0; s-- {
				l := 0
				if s >= c {
					l = s - c
				}
				dp[s] = pref[s+1].Sub(pref[l])
			}
		}

		bad := NewModInt(0)
		for _, v := range dp {
			bad = bad.Add(v)
		}
		ans = ans.Sub(bad)
	}

	return int(ans)
}
