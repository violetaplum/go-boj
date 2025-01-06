package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// L : 커서를 왼쪽으로 한 칸 옮김 (커서가 문장의 맨 앞이면 무시)
	// D : 커서를 오른쪽으로 한 칸 옮김 (커서가 문장의 맨 뒤이면 무시)
	// B : 커서 왼쪽에 있는 문자를 삭제함 (커서가 문장의 맨 앞이면 무시)
	// 삭제로 인해서 커서는 한 칸 왼쪽으로 이동한 것처럼 나타나고 실제로 커서의 오른쪽에 있던 문자는 그대로임
	// P $ : $ 라는 문자를 커서 왼쪽에 추가함

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	originalStr := scanner.Text()
	scanner.Scan()
	insertCount, _ := strconv.Atoi(scanner.Text())
	originalStrArray := strings.Split(originalStr, "")
	cursorIndex := len(originalStrArray) - 1

	for i := 0; i < insertCount; i++ {
		scanner.Scan()
		insertText := scanner.Text()
		switch insertText {
		case "L":
			if cursorIndex != 0 {
				cursorIndex--
			}

		case "D":
			if cursorIndex != len(originalStrArray)-1 {
				cursorIndex++
			}

		case "B":
			if cursorIndex != 0 {
				tempArr := originalStrArray[:cursorIndex]
				tempArr = append(tempArr, originalStrArray[cursorIndex:]...)
				originalStrArray = tempArr
				cursorIndex--
			}
		default: // P 가 들어왔을 경우를 생각한다
			tempArr := originalStrArray[:cursorIndex+1]
			tempArr = append(tempArr, string(insertText[2]))
			tempArr = append(tempArr, originalStrArray[cursorIndex:]...)
			originalStrArray = tempArr
		}
	}

	fmt.Print(strings.Join(originalStrArray, ""))

}
