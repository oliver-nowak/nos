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

	// resource := "/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/small_best.txt" // 2015/10/04 11:18:01 insertionsort took 10.261µs
	// resource := "/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/small_worst.txt" // 2015/10/04 11:18:19 insertionsort took 4.423505ms
	// resource := "/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/small_avg.txt" // 2015/10/04 11:18:36 insertionsort took 2.295572ms

	// resource := "/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/med_best.txt" // 2015/10/04 11:18:58 insertionsort took 85.881µs
	// resource := "/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/med_worst.txt" // 2015/10/04 11:19:14 insertionsort took 435.427908ms
	// resource := "/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/med_avg.txt" // 2015/10/04 11:19:29 insertionsort took 214.910065ms

	// resource := "/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/large_best.txt" // 2015/10/04 11:19:47 insertionsort took 838.455µs
	// resource := "/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/large_worst.txt" // 2015/10/04 11:20:44 insertionsort took 42.937227964s
	resource := "/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/large_avg.txt" // 2015/10/04 11:21:37 insertionsort took 21.552095589s

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

	sorted := InsertionSort(arry)

	fmt.Println(sorted[:5])
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func InsertionSort(A []int) []int {
	defer timeTrack(time.Now(), "insertionsort")

	for j := 1; j < len(A); j++ {

		key := A[j]

		i := j - 1

		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = key
	}

	return A
}
