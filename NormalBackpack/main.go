package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 준서가 가지고갈 물건 갯수 - N
	// 물건의 무게 - W
	// 물건의 가치 - V (준서가 생각하는 중요도)
	// 준서가 들 수 있는 무게 - K

	// 첫입력으로는 N, K 가 주어진다
	// (1 <= N <= 100)   (1 <= K <= 100,000)
	// 두번째줄부터는 추가로 N번의 입력을 받는다
	// 즉, W  V 를 N 번 입력받는다고 보면 된다
	// K 한도내에서 V 를 최대로 가질 수 있는 방법을 찾아보자

	scanner := bufio.NewScanner(os.Stdin)
	// N W 입력받기
	scanner.Scan()
	firstLine := strings.Split(scanner.Text(), " ")
	var N int
	var K int
	if len(firstLine) > 1 {
		N, _ = strconv.Atoi(firstLine[0])
		K, _ = strconv.Atoi(firstLine[1])
	}

	weights := make([]int, N)
	values := make([]int, N)

	// N번 만큼 추가 입력 받기
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

	// 입력받은 값 확인
	//fmt.Println("N, K : ", N, K)
	//for i, w := range weights {
	//	fmt.Println("W, V : ", w, values[i])
	//}

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
