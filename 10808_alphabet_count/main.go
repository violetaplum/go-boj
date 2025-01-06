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
	text := scanner.Text()
	textArray := strings.Split(text, "")
	counts := make([]int, 26)
	for _, v := range textArray {
		indexOfAlphabet := strings.Index("abcdefghijklmnopqrstuvwxyz", v)
		counts[indexOfAlphabet] += 1
	}

	for i, v := range counts {
		fmt.Print(v)
		if i == len(counts)-1 {
			break
		}
		fmt.Print(" ")
	}

}
