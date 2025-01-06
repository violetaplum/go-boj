package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateSlice(size int) []int {
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Int()%1000 - 500
	}

	return slice
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	middle := len(items) / 2
	left := mergeSort(items[:middle])
	right := mergeSort(items[middle:])

	// 정렬은 병합할때 일어난다
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	// left, right 가 둘다 배열 크기가 0이 됐을때 반복분을 종료한다
	for len(left) > 0 && len(right) > 0 {
		// 하나씩 비교해서 넣는다..!!
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	// result 에 left, right 삽입
	result = append(result, left...)
	result = append(result, right...)

	return result
}

func main() {
	slice := generateSlice(20)
	fmt.Println("---Unsorted---")
	fmt.Println(slice)
	fmt.Println("---Sorted---")
	fmt.Println(mergeSort(slice))
}
