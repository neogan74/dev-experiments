package _494_FindtheMinimumAmountofTimetoBrewPotions

import "testing"

func minTime2D(skill, mana []int) int64 {
	n, m := len(skill), len(mana)
	dp := make([][]int64, n)
	for i := range dp {
		dp[i] = make([]int64, m)
		for j := 0; j < m; j++ {
			t := int64(skill[i]) * int64(mana[j])
			var up, left int64
			if i > 0 {
				up = dp[i-1][j]
			}
			if j > 0 {
				left = dp[i][j-1]
			}
			if up > left {
				dp[i][j] = up + t
			} else {
				dp[i][j] = left + t
			}
		}
	}
	return dp[n-1][m-1]
}

func TestBasic(t *testing.T) {
	if got := minTime([]int{1}, []int{5}); got != 5 {
		t.Fatalf("got %d want 5", got)
	}
	if got := minTime([]int{1, 2}, []int{3}); got != 9 { // 1*3 + 2*3
		t.Fatalf("got %d want 9", got)
	}
	if got := minTime([]int{1, 2}, []int{3, 4}); got != minTime2D([]int{1, 2}, []int{3, 4}) {
		t.Fatal("mismatch 2x2")
	}
}

func TestRandomSmall(t *testing.T) {
	tests := []struct{ skill, mana []int }{
		{[]int{1, 2, 3}, []int{1, 2}},
		{[]int{2, 2}, []int{3, 1, 4}},
		{[]int{5}, []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		got := minTime(tt.skill, tt.mana)
		want := minTime2D(tt.skill, tt.mana)
		if got != want {
			t.Fatalf("skill=%v mana=%v got=%d want=%d", tt.skill, tt.mana, got, want)
		}
	}
}
