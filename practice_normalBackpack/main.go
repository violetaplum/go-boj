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
	firstLine := strings.Split(scanner.Text(), " ")
	var N int
	var K int

	if len(firstLine) > 1 {
		N, _ = strconv.Atoi(firstLine[0])
		K, _ = strconv.Atoi(firstLine[1])
	}

	weights := make([]int, N)
	values := make([]int, N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		lines := strings.Split(scanner.Text(), " ")
		if len(lines) > 1 {
			weights[i], _ = strconv.Atoi(lines[0])
			values[i], _ = strconv.Atoi(lines[1])
		} else {
			continue
		}
	}
	fmt.Println(knapsack(K, weights, values, N))

}

func knapsack(K int, weights []int, values []int, N int) int {
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, K+1)
	}

	for i := 1; i <= N; i++ {
		for w := 0; w <= K; w++ {
			if weights[i-1] <= w {
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]]+values[i-1])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}
	return dp[N][K]
}
