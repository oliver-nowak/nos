package main

import (
	"fmt"
)

var A = []int{
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1}

func main() {
	fmt.Println(A)

	for j := 1; j < len(A); j++ {

		key := A[j]

		i := j - 1

		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = key
	}

	fmt.Println(A)
}
