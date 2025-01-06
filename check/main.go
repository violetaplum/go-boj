package main

//
//import "fmt"
//
//func findMostFrequentPattern(interview string) string {
//	// 길이가 1인 경우는 처리할 필요 없음
//	if len(interview) <= 1 {
//		return interview
//	}
//
//	// 가능한 모든 패턴을 찾아서 카운트
//	maxCount := 0
//	maxPattern := ""
//
//	// 문자열의 각 위치에서 시작하는 모든 가능한 패턴 탐색
//	for i := 0; i < len(interview); i++ {
//		for j := i + 1; j <= len(interview); j++ {
//			pattern := interview[i:j]
//			count := 0
//
//			// 현재 패턴이 문자열에서 몇 번 등장하는지 카운트
//			tempStr := interview
//			for {
//				if idx := findIndex(tempStr, pattern); idx != -1 {
//					count++
//					// 패턴이 발견된 부분을 제거
//					tempStr = tempStr[:idx] + tempStr[idx+len(pattern):]
//				} else {
//					break
//				}
//			}
//
//			// 더 많이 등장하는 패턴을 찾았거나,
//			// 같은 횟수로 등장하는 더 긴 패턴을 찾은 경우 업데이트
//			if count > maxCount {
//				maxCount = count
//				maxPattern = pattern
//			}
//		}
//	}
//
//	// 패턴을 모두 제거하고 남은 문자열 반환
//	result := interview
//	for i := 0; i < maxCount; i++ {
//		if idx := findIndex(result, maxPattern); idx != -1 {
//			result = result[:idx] + result[idx+len(maxPattern):]
//		}
//	}
//
//	return result
//}
//
//// 문자열에서 패턴의 첫 등장 위치를 찾는 헬퍼 함수
//func findIndex(s, pattern string) int {
//	for i := 0; i <= len(s)-len(pattern); i++ {
//		if s[i:i+len(pattern)] == pattern {
//			return i
//		}
//	}
//	return -1
//}
//
//func solution(interview string) string {
//	return findMostFrequentPattern(interview)
//}
//
//func main() {
//	str := "abcabcdefabc"
//	fmt.Println(solution(str))
//}
