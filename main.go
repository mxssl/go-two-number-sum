package TwoNumberSum

import "sort"

func twoNumberSum(array []int, target int) []int {
	nums := make(map[int]bool)
	for _, num := range array {
		match := target - num
		if nums[match] {
			result := []int{match, num}
			sort.Ints(result)
			return result
		}
		nums[num] = true
	}
	return []int{}
}
