package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	command := scanner.Text()
	commands := strings.Split(command, "")
	st := []string{}
	count := 0
	top := -1
	for i, v := range commands {

	}
	fmt.Println(count)
}
