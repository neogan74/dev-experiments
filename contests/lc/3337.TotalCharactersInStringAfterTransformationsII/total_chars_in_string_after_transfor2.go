package _3337_total_characters_in_string_after_transfor2

const mod = 1_000_000_007
const size = 26

// Основная функция
func lengthAfterTransformations(s string, t int, nums []int) int {
	count := make([]int, size)
	for _, c := range s {
		count[c-'a']++
	}

	T := buildMatrix(nums)
	Tt := matPow(T, t)

	// result = T^t × count
	res := 0
	for i := 0; i < size; i++ {
		sum := 0
		for j := 0; j < size; j++ {
			sum = (sum + Tt[i][j]*count[j]) % mod
		}
		res = (res + sum) % mod
	}
	return res
}

// Построение матрицы переходов: T[to][from] = +1 для каждого правила
func buildMatrix(nums []int) [][]int {
	T := make([][]int, size)
	for i := 0; i < size; i++ {
		T[i] = make([]int, size)
	}
	for from := 0; from < size; from++ {
		for step := 1; step <= nums[from]; step++ {
			to := (from + step) % size
			T[to][from]++
		}
	}
	return T
}

// Умножение двух матриц (T × T)
func matMul(a, b [][]int) [][]int {
	res := make([][]int, size)
	for i := range res {
		res[i] = make([]int, size)
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				res[i][j] = (res[i][j] + a[i][k]*b[k][j]) % mod
			}
		}
	}
	return res
}

// Быстрое возведение матрицы в степень
func matPow(mat [][]int, exp int) [][]int {
	res := identity()
	for exp > 0 {
		if exp%2 == 1 {
			res = matMul(res, mat)
		}
		mat = matMul(mat, mat)
		exp /= 2
	}
	return res
}

// Единичная матрица 26×26
func identity() [][]int {
	id := make([][]int, size)
	for i := range id {
		id[i] = make([]int, size)
		id[i][i] = 1
	}
	return id
}
