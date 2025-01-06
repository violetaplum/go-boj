package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// N -> 밴치의 길이
	// K -> 사람이 뻗을 수 있는 최대 거리
	// P -> 사람의 위치
	// H -> 햄버거의 위치
	// https://www.acmicpc.net/status?user_id=wkdgmltn77&problem_id=19941&from_mine=1

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	firstLine := strings.Split(scanner.Text(), " ")

	var N int
	var K int
	var gogo bool
	count := 0

	if len(firstLine) > 1 {
		N, _ = strconv.Atoi(firstLine[0])
		K, _ = strconv.Atoi(firstLine[1])
	}

	scanner.Scan()
	secondLine := strings.Split(scanner.Text(), "")

	for i := 0; i < N; i++ {
		if secondLine[i] == "P" {
			gogo = true

			for j := K; j >= 1; j-- {
				if !(i-j < 0) {
					if secondLine[i-j] == "H" {
						secondLine[i-j] = ""
						count++
						gogo = false // 밑에를 타지 말게끔
						break
					}
				}
			}
			if gogo {
				for j := 1; j <= K; j++ {
					if i+j >= N {
						break
					} else {
						if secondLine[i+j] == "H" {
							secondLine[i+j] = ""
							count++
							break
						}
					}
				}
			}
		}
	}

	// 햄버거와 사람을 false 로 천천히 바꾸고
	// 나중에 false 로 바뀐 사람들을 세면 몇명이 먹었는지 알수있다
	fmt.Println(count)
}
