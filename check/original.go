package main

import "fmt"

func solution2(interview string) string {
	// delete the pattern!
	if len(interview) <= 1 {
		return interview
	}

	maxCount := 0
	maxPattern := ""

	for i := 0; i < len(interview); i++ {
		for j := i + 1; j <= len(interview); j++ {
			pattern := interview[i:j]
			count := 0

			tempStr := interview
			for {
				if idx := findIndex2(tempStr, pattern); idx != -1 {
					count++
					tempStr = tempStr[:idx] + tempStr[idx+len(pattern):]
				} else {
					break
				}
			}

			if count > maxCount {
				maxCount = count
				maxPattern = pattern
			}
		}
	}

	result := interview
	for i := 0; i < maxCount; i++ {
		idx := findIndex2(result, maxPattern)
		if idx != -1 {
			result = result[:idx] + result[idx+len(maxPattern):]
		}
	}

	return result
}

func findIndex2(s, pattern string) int {
	for i := 0; i <= len(s)-len(pattern); i++ {
		if s[i:i+len(pattern)] == pattern {
			return i
		}
	}
	return -1
}

func main() {
	st := "abcabcdefabc"
	fmt.Println(solution2(st))
}
