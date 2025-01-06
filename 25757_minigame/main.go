package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 미니게임 Y, F, O
	// Y - 2명 -> 1명 조합
	// F - 3명 -> 2명 조합
	// O - 4명 -> 3명 조합

	// 신청횟수 N -> 입력횟수
	// 종류 주어진다
	// 최대 몇번이나 임스와 게임할 수 있는가?
	// 임스는 한번 같이 플레이한 사람과는 다시 플레이 하지 않는다

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	firstLine := strings.Split(scanner.Text(), " ")
	var N int
	var game string
	if len(firstLine) > 1 {
		N, _ = strconv.Atoi(firstLine[0])
		game = firstLine[1]
	} else {
		//fmt.Println("Not enough input")
	}

	gamers := make(map[string]bool, N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		line := scanner.Text()
		if _, ok := gamers[line]; !ok {
			gamers[line] = true
		}
	}

	var gameValue int

	switch game {
	case "Y":
		gameValue = len(gamers) / 1
	case "F":
		gameValue = len(gamers) / 2
	case "O":
		gameValue = len(gamers) / 3
	default:
		//fmt.Println("invalid game kind..")
	}

	fmt.Println(gameValue)

}
