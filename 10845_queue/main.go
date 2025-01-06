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
	queue := []int{}
	tail := -1
	for i := 0; i < count; i++ {
		scanner.Scan()
		command := scanner.Text()
		commands := strings.Split(command, " ")
		if len(commands) > 0 {
			firstCommand := commands[0]
			secondCommand := 0
			if len(commands) > 1 {
				secondCommand, _ = strconv.Atoi(commands[1])
			}
			switch firstCommand {
			case "push":
				queue = append(queue, secondCommand)
				tail += 1
			case "pop":
				if tail != -1 && len(queue) > 0 { // 무언가 있을때
					fmt.Println(queue[0])
					if len(queue) > 1 {
						queue = queue[1:]
					} else {
						queue = []int{}
					}
					tail -= 1
				} else {
					fmt.Println(tail)
				}
			case "size":
				fmt.Println(len(queue))
			case "empty":
				if tail == -1 {
					fmt.Println(1)
				} else {
					fmt.Println(0)
				}
			case "front":
				if tail == -1 {
					fmt.Println(tail)
				} else if len(queue) > 0 {
					fmt.Println(queue[0])
				} else {
					fmt.Println("Error!!")
				}
			case "back":
				if tail == -1 {
					fmt.Println(tail)
				} else if len(queue) > 0 {
					fmt.Println(queue[tail])
				} else {
					fmt.Println("Error!?")
				}

			}

		}
	}
}
