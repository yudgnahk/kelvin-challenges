package daily

// https://leetcode.com/problems/rotate-array
// 189. Rotate Array
func rotate(nums []int, k int) {
	if k == 0 {
		return
	}

	numsLen := len(nums)
	copyNums := make([]int, numsLen)
	copy(copyNums, nums)

	if k > numsLen {
		k = k % numsLen
	}
	l := numsLen - k

	for i := 0; i < numsLen; i++ {
		nums[i] = copyNums[(i+l)%numsLen]
	}
}
