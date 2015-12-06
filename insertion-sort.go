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

	arrayToSort := getArrayToSort()

	sorted := InsertionSort(arrayToSort)

	fmt.Println(sorted[:10])
}

func getArrayToSort() []int {
	arry := make([]int, 250001)

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

	return arry
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
