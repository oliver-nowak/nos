package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	arry := make([]int, 250001)

	// resource := "/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/small_best.txt" // 2015/10/04 11:43:04 selectionsort took 3.604193ms
	// resource := "/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/small_worst.txt" // 2015/10/04 11:43:27 selectionsort took 4.71628ms
	// resource := "/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/small_avg.txt" // 2015/10/04 11:43:44 selectionsort took 4.182817ms

	// resource := "/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/med_best.txt" // 2015/10/04 11:44:17 selectionsort took 342.872365ms
	// resource := "/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/med_worst.txt" // 2015/10/04 11:44:35 selectionsort took 341.910855ms
	// resource := "/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/med_avg.txt" // 2015/10/04 11:44:56 selectionsort took 344.568104ms

	// resource := "/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/large_best.txt" // 2015/10/04 11:45:58 selectionsort took 34.381607187s
	// resource := "/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/large_worst.txt" // 2015/10/04 11:46:47 selectionsort took 34.397779091s
	resource := "/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/large_avg.txt" // 2015/10/04 11:47:42 selectionsort took 34.407796662s

	f, err := os.Open(resource)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)

	scanner := bufio.NewScanner(r)
	idx := 0

	// read each line
	for scanner.Scan() {
		txt := scanner.Text()
		i, _ := strconv.Atoi(txt)
		arry[idx] = i
		idx += 1
	}

	sorted := SelectionSort(arry)

	fmt.Println(sorted[:5])
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func SelectionSort(A []int) []int {
	defer timeTrack(time.Now(), "selectionsort")

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

	return A
}
