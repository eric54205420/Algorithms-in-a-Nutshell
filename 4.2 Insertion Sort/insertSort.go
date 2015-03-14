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
	fmt.Println(insertSort(l))
}

func randInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func insertSort(l []int) []int {
	var nl []int
	for i := 0; i < len(l); i++ {
		inserted := false
		nl = []int{}
		for j := 0; j < i; j++ {
			if l[i] < l[j] {
				nl = append(nl, l[i])
				nl = append(nl, l[j:i]...)
				inserted = true
				break
			} else {
				nl = append(nl, l[j])
			}
		}
		if !inserted {
			nl = append(nl, l[i])
		}
		nl = append(nl, l[i+1:len(l)]...)
		l = nl
	}
	return nl
}
