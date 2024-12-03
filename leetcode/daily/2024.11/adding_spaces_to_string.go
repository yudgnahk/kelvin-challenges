package daily

// 2109. Adding Spaces to a String
// https://leetcode.com/problems/adding-spaces-to-a-string/
func addSpaces(s string, spaces []int) string {
	if len(spaces) == 0 {
		return s
	}

	resp := ""
	curSpace := 0
	if spaces[curSpace] == 0 {
		resp += " "
		curSpace++
	}

	isFinished := false
	finishedPos := 0

	for i := 0; i < len(s); i++ {
		if i < spaces[curSpace] {
			resp += string(s[i])
		} else {
			resp += " "
			if curSpace+1 < len(spaces) {
				curSpace++
			} else {
				isFinished = true
				finishedPos = i
				break
			}
		}
	}

	if isFinished {
		resp += s[finishedPos:]
	}

	return resp
}
