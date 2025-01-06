package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	stack := []string{}
	top := -1
	for i := 0; i < count; i++ {
		scanner.Scan()
		command := strings.Split(scanner.Text(), " ")
		if len(command) > 0 {
			switch command[0] {
			case "push":
				stack = append(stack, command[1])
				top += 1
			case "pop":
				if top != -1 {
					fmt.Println(stack[top])
					stack = stack[:top]
					top -= 1
				} else {
					fmt.Println(top)
				}
			case "size":
				if top == -1 {
					fmt.Println(0)
				} else {
					fmt.Println(len(stack))
				}
			case "empty":
				if top == -1 {
					fmt.Println(1)
				} else {
					fmt.Println(0)
				}
			case "top":
				if top != -1 {
					fmt.Println(stack[top])
				} else {
					fmt.Println(top)
				}

			}
		}
	}
}
