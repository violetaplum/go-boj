package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	command := ""

	for {
		scanner.Scan()
		command = scanner.Text()
		if command == "." {
			break
		}
		st := []string{}
		top := -1
		check := 1
		letters := strings.Split(command, "")
		for i := 0; i < len(letters); i++ {
			if letters[i] == "(" {
				st = append(st, letters[i])
				top++
			}
			if letters[i] == "[" {
				st = append(st, letters[i])
				top++
			}
			if letters[i] == ")" {
				if top == -1 {
					check *= 0
					break
				}
				if st[top] == "(" {
					st = st[:top]
					top--
				} else {
					check *= 0
					break
				}
			}

			if letters[i] == "]" {
				if top == -1 {
					check *= 0
					break
				}
				if st[top] == "[" {
					st = st[:top]
					top--
				} else {
					check *= 0
					break
				}
			}
		}
		if len(st) > 0 {
			check *= 0
		}
		if check > 0 {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
