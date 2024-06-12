package daily

// https://leetcode.com/problems/merge-sorted-array
// 88. Merge Sorted Array
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	for k := m + n - 1; k >= 0; k-- {
		if j < 0 {
			break
		}
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
