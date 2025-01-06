package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deque struct {
	data []int
	head int
	tail int
}

func NewDeque(size int) *Deque {
	return &Deque{
		data: make([]int, size),
		head: size / 2,
		tail: size/2 + 1,
	}
}

func (d *Deque) PushFront(x int) {
	d.data[d.head] = x
	d.head--
}

func (d *Deque) PushBack(x int) {
	d.data[d.tail] = x
	d.tail++
}

func (d *Deque) PopBack() int {
	if d.IsEmpty() {
		return -1
	}
	d.tail--
	res := d.data[d.tail]
	d.data[d.tail] = 0
	return res
}

func (d *Deque) PopFront() int {
	if d.IsEmpty() {
		return -1
	}
	d.head++
	res := d.data[d.head]
	d.data[d.head] = 0
	return res
}

func (d *Deque) Size() int {
	return d.tail - d.head - 1
}

func (d *Deque) IsEmpty() bool {
	return d.tail-d.head <= 1
}

func (d *Deque) Front() int {
	if d.IsEmpty() {
		return -1
	}
	return d.data[d.head+1]
}

func (d *Deque) Back() int {
	if d.IsEmpty() {
		return -1
	}

	return d.data[d.tail-1]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	count, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	deque := NewDeque(20002)

	for i := 0; i < count; i++ {
		scanner.Scan()
		command := scanner.Text()
		commands := strings.Split(command, " ")
		switch commands[0] {
		case "push_front":
			x, _ := strconv.Atoi(commands[1])
			deque.PushFront(x)
		case "push_back":
			x, _ := strconv.Atoi(commands[1])
			deque.PushBack(x)
		case "pop_front":
			fmt.Println(deque.PopFront())
		case "pop_back":
			fmt.Println(deque.PopBack())
		case "size":
			fmt.Println(deque.Size())
		case "empty":
			if deque.IsEmpty() {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		case "front":
			fmt.Println(deque.Front())
		case "back":
			fmt.Println(deque.Back())
		}
	}
}
