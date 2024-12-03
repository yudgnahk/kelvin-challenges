package daily

import "strings"

// 2109. Adding Spaces to a String
// https://leetcode.com/problems/adding-spaces-to-a-string/
func addSpaces(s string, spaces []int) string {
	if len(spaces) == 0 {
		return s
	}

	var builder strings.Builder
	builder.Grow(len(s) + len(spaces))

	lastPos := 0
	for _, space := range spaces {
		builder.WriteString(s[lastPos:+space])
		builder.WriteByte(' ')
		lastPos = space
	}
	builder.WriteString(s[lastPos:])

	return builder.String()
}
