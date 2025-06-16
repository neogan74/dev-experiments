package _3335_total_characters_in_string_after_trans

const MOD int = 1000000007

func totalCharacters(s string, t int) int {
	cnt := make([]int, 26)

	for _, char := range s {
		cnt[char-'a']++
	}

	for j := 0; j < t; j++ {
		tmp := make([]int, 26)
		for i := 0; i < 26; i++ {
			if i == 25 {
				tmp[0] = (tmp[0] + cnt[i]) % MOD
				tmp[1] = (tmp[1] + cnt[i]) % MOD
			} else {
				tmp[i+1] = (tmp[i+1] + cnt[i]) % MOD
			}
		}
		cnt = tmp
	}

	len := 0
	for _, c := range cnt {
		len = (len + c) % MOD
	}

	return len
}
