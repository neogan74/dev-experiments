package _494_FindtheMinimumAmountofTimetoBrewPotions

func minTime(skill []int, mana []int) int64 {
	n, m := len(skill), len(mana)
	if n == 0 || m == 0 {
		return 0
	}
	var sumSkill int64
	for _, s := range skill {
		sumSkill += int64(s)
	}

	// время окончания последнего мага для текущего зелья
	prevWizardDone := sumSkill * int64(mana[0])

	for j := 1; j < m; j++ {
		prevPotionDone := prevWizardDone
		for i := n - 2; i >= 0; i-- {
			prevPotionDone -= int64(skill[i+1]) * int64(mana[j-1])
			cand := prevWizardDone - int64(skill[i])*int64(mana[j])
			if prevPotionDone > cand {
				prevWizardDone = prevPotionDone
			} else {
				prevWizardDone = cand
			}
		}
		prevWizardDone += sumSkill * int64(mana[j])
	}
	return prevWizardDone
}
