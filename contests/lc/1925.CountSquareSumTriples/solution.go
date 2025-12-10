package countsquaresumtriples

import "math"

// countTriples returns the number of ordered square triples (a, b, c)
// such that 1 <= a, b, c <= n and a^2 + b^2 == c^2.
func countTriples(n int) int {
	squares := make(map[int]struct{}, n)
	for i := 1; i <= n; i++ {
		squares[i*i] = struct{}{}
	}

	limit := n * n
	count := 0

	for a := 1; a <= n; a++ {
		aa := a * a
		for b := 1; b <= n; b++ {
			sum := aa + b*b
			if sum > limit {
				break // c would exceed n since squares grow with b
			}
			if _, ok := squares[sum]; ok {
				count++
			}
		}
	}

	return count
}

const mx = 23 // floor(sqrt(250 * 2)) + 1
var mu = [mx]int{1: 1}
var divisors [mx][]int

func init() {
	for i := 1; i < mx; i++ {
		for j := i * 2; j < mx; j += i {
			mu[j] -= mu[i]
		}
	}

	for i := 1; i < mx; i++ {
		if mu[i] == 0 {
			continue
		}
		for j := i; j < mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}
}

func countCoprime(n, x int) (res int) {
	for _, d := range divisors[x] {
		res += mu[d] * (n / d)
	}
	return
}

func countCoprimeOdd(n, x int) (res int) {
	return countCoprime(n, x) - countCoprime(n/2, x)
}

func countTriples2(n int) (ans int) {
	B := int(math.Cbrt(float64(n)))

	for u := 3; u*u < n*2; u += 2 {
		if u <= B {
			for v := 1; v < u && u*u+v*v <= n*2; v += 2 {
				if gcd(u, v) == 1 {
					ans += n * 2 / (u*u + v*v)
				}
			}
			continue
		}

		for l, r := 1, 0; l < u && u*u+l*l <= n*2; l = r + 1 {
			num := n * 2 / (u*u + l*l)
			r = min(int(math.Sqrt(float64(n*2/num-u*u))), u-1)
			validV := countCoprimeOdd(r, u) - countCoprimeOdd(l-1, u)
			ans += num * validV
		}
	}

	return ans * 2
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
