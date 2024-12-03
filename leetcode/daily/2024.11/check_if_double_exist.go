package daily

// 1346. Check If N and Its Double Exist
// https://leetcode.com/problems/check-if-n-and-its-double-exist/
func checkIfExist(arr []int) bool {
	m := make(map[int]bool)
	for _, v := range arr {
		if m[v*2] || (v%2 == 0 && m[v/2]) {
			return true
		}
		m[v] = true
	}
	return false
}
