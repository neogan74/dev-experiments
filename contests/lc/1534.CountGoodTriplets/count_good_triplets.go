package main

func countGoodTriplets(arr []int, a int, b int, c int) int {
	count := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if abs(arr[i]-arr[j]) > a {
				continue
			}
			for k := j + 1; k < n; k++ {
				if abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					count++
				}
			}
		}
	}
	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countGoodTripletsOptimized(arr []int, a, b, c int) int {
	n := len(arr)
	count := 0

	for j := 1; j < n-1; j++ {
		for i := 0; i < j; i++ {
			if abs(arr[i]-arr[j]) <= a {
				for k := j + 1; k < n; k++ {
					if abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
						count++
					}
				}
			}
		}
	}
	return count
}
