package iterations

import "strings"

func Iterate(a string, count int) string {
	var result strings.Builder
	for i := 0; i < count; i++ {
		result.WriteString(a)
	}
	return result.String()
}
