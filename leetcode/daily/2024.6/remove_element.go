package daily

import "math"

// https://leetcode.com/problems/remove-element
// 27. Remove Element
func removeElement(nums []int, val int) int {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		j++
	}

	k := i
	for k < len(nums) {
		nums[k] = math.MaxInt
		k++
	}

	return i
}
