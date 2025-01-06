package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://www.acmicpc.net/problem/20125

func main() {
	// 첫 별이 찍혔을때 머리이다
	// 심장의 위치, 팔, 다리, 허리의 길이를 구해라

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	//pan := make([][]string, N)
	//pan2 := make([]string, N)
	headFlag := true
	armsFlag := true
	bodyFlag := true

	heartLocation := ""
	heartt := 0

	leftArmLength := 0
	rightArmLength := 0
	bodyLength := 0

	leftLegLength := 0
	rightLegLenth := 0
	leftLeg := 0
	rightLeg := 0

	for i := 0; i < N; i++ {
		scanner.Scan()
		input := scanner.Text()
		if strings.Contains(input, "*") && headFlag {
			heartt = strings.Index(input, "*")
			heartLocation = strconv.Itoa(i+2) + " " + strconv.Itoa(heartt+1)
			headFlag = false
			continue
		}
		if strings.Contains(input, "***") && !headFlag && armsFlag {
			for j := 0; j < len(input); j++ {
				if input[j] == '*' && j < heartt {
					leftArmLength++
				}
				if input[j] == '*' && j > heartt {
					rightArmLength++
				}
			}
			armsFlag = false
		}

		if strings.Contains(input, "_*_") && !headFlag && !armsFlag && bodyFlag {
			bodyLength++
		}
		if strings.Contains(input, "*_*") && !headFlag && !armsFlag && bodyFlag {
			bodyFlag = false
			leftLegLength++
			rightLegLenth++
			left := true
			for j := 0; j < len(input); j++ {
				if left && input[j] == '*' {
					left = false
					leftLeg = j
				}
				if !left && input[j] == '*' {
					rightLeg = j
					break
				}
			}
		}
		if !bodyFlag {
			if input[leftLeg] == '*' {
				leftLegLength++
			}
			if input[rightLeg] == '*' {
				rightLegLenth++
			}
		}
	}
	fmt.Println(heartLocation)
	fmt.Println(leftArmLength, rightArmLength, bodyLength, leftLegLength, rightLegLenth)
}
