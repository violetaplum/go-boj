package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkStr(input string) bool {
	check := true
	top := -1
	st := []byte{}

	for _, v := range input {
		switch v {
		case '(':
			st = append(st, '(')
			top++
		case '[':
			st = append(st, '[')
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
		case ']':
			if top != -1 {
				if st[top] == '[' {
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
		}
	}

	if len(st) > 0 { // 괄호가 남은경우!!
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
				} else if val == -2 { // 잘못된 괄호쌍
					return 0
				} else {
					temp += val
				}
			}
		case ']':
			temp := 0
			for len(stack) > 0 {
				val := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if val == -2 { // 제대로된 괄호가 합쳐졌을때!
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
		}
	}

	result := 0
	for _, v := range stack {
		if v == -1 || v == -2 { // 스택에 괄호가 남아있는경우
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

	// 줄바꿈 문자 \n 이 나올때까지 계속 byte 를 읽는다
	str, _ := reader.ReadBytes('\n')

	str = str[:len(str)-1]

	if !checkStr(string(str)) {
		fmt.Fprint(writer, 0)
	} else {
		result := calculateValue(string(str))
		fmt.Fprint(writer, result)
	}
}
