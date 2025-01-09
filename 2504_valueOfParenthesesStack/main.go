package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkStr(input []byte) bool {
	check := true
	top := -1
	st := []byte{}
	for _, v := range input {
		switch v {
		case '(':
			st = append(st, v)
			top++
		case ')':
			if top != -1 {
				if st[top] == '(' {
					st = st[:top]
					top--
				} else {
					check = false
					break
				}
			} else {
				check = false
				break
			}
		case '[':
			st = append(st, v)
			top++
		case ']':
			if top != -1 {
				if st[top] == '[' {
					st = st[:top]
					top--
				} else {
					check = false
					break
				}
			} else { // 이 부분이 빠져있었습니다
				check = false
				break
			}
		}
	}
	if len(st) > 0 {
		check = false
	}
	return check
}

func calculateValue(str string) int {
	stack := make([]int, 0, len(str))

	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '(':
			stack = append(stack, -1)
		case '[':
			stack = append(stack, -2)
		case ')':
			temp := 0
			for len(stack) > 0 {
				val := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if val == -1 {
					if temp == 0 {
						temp = 2
					} else {
						temp *= 2
					}
					stack = append(stack, temp)
					break
				} else if val == -2 {
					return 0
				} else {
					temp += val
				}
			}
			if len(stack) == 0 { // 스택이 비어있는 경우 처리
				return 0
			}
		case ']':
			temp := 0
			for len(stack) > 0 {
				val := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if val == -2 {
					if temp == 0 {
						temp = 3
					} else {
						temp *= 3
					}
					stack = append(stack, temp)
					break
				} else if val == -1 {
					return 0
				} else {
					temp += val
				}
			}
			if len(stack) == 0 { // 스택이 비어있는 경우 처리
				return 0
			}
		}
	}

	result := 0
	for _, v := range stack {
		if v == -1 || v == -2 { // 스택에 여는 괄호가 남아있는 경우
			return 0
		}
		result += v
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	str, _ := reader.ReadBytes('\n')
	for len(str) > 0 && (str[len(str)-1] == '\n' || str[len(str)-1] == '\r') {
		str = str[:len(str)-1]
	}

	if !checkStr(str) {
		fmt.Fprint(writer, 0)
	} else {
		result := calculateValue(string(str))
		fmt.Fprint(writer, result)
	}
}
