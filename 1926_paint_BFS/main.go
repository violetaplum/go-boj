package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Point는 그림에서의 좌표를 나타냅니다
type Point struct {
	y, x int
}

// 방향 벡터 정의 (상, 우, 하, 좌)
var dy = []int{1, 0, -1, 0}
var dx = []int{0, 1, 0, -1}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// N, M 입력 받기
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	// 그림 정보 입력 받기
	picture := make([][]int, n)
	visited := make([][]bool, n)
	for i := range picture {
		picture[i] = make([]int, m)
		visited[i] = make([]bool, m)
		for j := range picture[i] {
			scanner.Scan()
			picture[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}

	pictureCnt := 0
	maxPicture := 0

	// 모든 좌표에 대해 BFS 수행
	for row := 0; row < n; row++ {
		for col := 0; col < m; col++ {
			if picture[row][col] == 1 && !visited[row][col] {
				cnt := BFS(Point{row, col}, n, m, picture, visited)
				pictureCnt++
				if cnt > maxPicture {
					maxPicture = cnt
				}
			}
		}
	}

	fmt.Println(pictureCnt)
	fmt.Println(maxPicture)
}

// BFS 함수 구현
func BFS(start Point, n, m int, picture [][]int, visited [][]bool) int {
	cnt := 0
	queue := []Point{start}
	visited[start.y][start.x] = true

	for len(queue) > 0 {
		// 큐에서 원소 추출 (파이썬의 pop()은 Go에서 슬라이스로 구현)
		current := queue[0]
		queue = queue[1:]
		cnt++

		// 4방향 탐색
		for i := 0; i < 4; i++ {
			ny := current.y + dy[i]
			nx := current.x + dx[i]

			// 범위 체크 및 방문 조건 확인
			if ny >= 0 && ny < n && nx >= 0 && nx < m {
				if !visited[ny][nx] && picture[ny][nx] == 1 {
					visited[ny][nx] = true
					queue = append(queue, Point{ny, nx})
				}
			}
		}
	}

	return cnt
}
