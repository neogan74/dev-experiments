package _75_sort_colors

func sortColors(nums []int) {
	i, j, k := -1, len(nums), 0
	for k < j {
		switch nums[k] {
		case 0:
			i++
			nums[i], nums[k] = nums[k], nums[i]
			k++
		case 2:
			j--
			nums[j], nums[k] = nums[k], nums[j]
		default:
			k++
		}
	}
}
