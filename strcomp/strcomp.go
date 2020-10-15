package strcomp

import (
	"fmt"
	"strings"
)

// Compress ...
func Compress(str string) string {
	result := make([][]int32, 0)
	for _, char := range str {
		if len(result) > 0 {
			prevIdx := len(result) - 1
			prev := result[prevIdx]
			if prev[0] == char {
				result[prevIdx][1]++
				continue
			}
		}
		result = append(result, []int32{
			char, 1,
		})
	}
	return output(result)
}

func output(result [][]int32) string {
	var strBuilder strings.Builder
	defer strBuilder.Reset()

	for _, r := range result {
		if r[1] > 1 {
			fmt.Fprintf(&strBuilder, "%s%d", string(r[0]), int(r[1]))
			continue
		}
		fmt.Fprintf(&strBuilder, "%s", string(r[0]))
	}

	// so that we have a new line afterwards
	return strBuilder.String()
}
