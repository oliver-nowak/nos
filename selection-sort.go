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

	for j := 0; j < len(A)-1; j++ {
		iMin := j

		for i := j + 1; i < len(A); i++ {
			if A[i] < A[iMin] {
				iMin = i
			}
		}

		if iMin != j {
			tmp := A[j]

			A[j] = A[iMin]

			A[iMin] = tmp
		}
	}

	fmt.Println(A)
}
