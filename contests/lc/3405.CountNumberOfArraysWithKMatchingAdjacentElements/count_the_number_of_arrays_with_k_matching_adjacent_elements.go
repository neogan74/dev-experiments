package _3405_CountNumberOfArraysWithKMatchingAdjacentElements

// mod is special for this problem
const mod = 1e9 + 7

// modPow is func for pow with mod
func modPow(a, e int64) int64 {
	res := int64(1)
	for e > 0 {
		if e&1 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		e >>= 1
	}
	return res
}

func modInv(a int64) int64 {
	return modPow(a, mod-2)
}

func countGoodArrays(n, m, k int) int {
	N := int64(n - 1)
	K := int64(k)
	M := int64(m)

	// C(n-1, k) = fact[N] / (fact[K] * fact[N-K])
	// Предварительных вычислений нет, используем формулу на ходу:
	num := int64(1)
	for i := int64(0); i < K; i++ {
		num = num * (N - i) % mod
		num = num * modInv(i+1) % mod
	}

	pow := modPow(M-1, int64(n-k-1))
	ans := num * M % mod * pow % mod
	return int(ans)
}
