package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	l := []int{}
	for i := 0; i < 10; i++ {
		l = append(l, randInt(100))
	}
	fmt.Println(l)
	fmt.Println(medianSort(l))
}

func randInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func medianSort(l []int) []int {
	m := len(l) / 2

	left := []int{}
	right := []int{}
	for i, n := range l {
		if i == m {
			continue
		}
		if n > l[m] {
			right = append(right, n)
		} else {
			left = append(left, n)
		}
		if len(right) > 1 {
			right = medianSort(right)
		}
		if len(left) > 1 {
			left = medianSort(left)
		}
	}

	r := left
	r = append(left, l[m])
	r = append(r, right...)
	return r
}
