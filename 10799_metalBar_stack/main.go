package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	str, _ := reader.ReadString('\n')
	str = str[:len(str)-1] // 개행문자 제거

	st := make([]byte, 0, len(str))

	count := 0
	top := -1
	laser := false
	for _, v := range str {
		if v == '(' {
			st = append(st, '(')
			top++
			laser = true
		} else if v == ')' {
			if laser {
				st = st[:len(st)-1]
				top--
				count += len(st)
			} else {
				st = st[:top]
				count += 1
				top--
			}
			laser = false
		}
	}
	fmt.Fprint(writer, count)
}
