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

	sorted := MergeSort(arrayToSort)

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

func MergeSort(A []int) []int {
	var arrayLength = len(A)

	// Base case
	if arrayLength == 1 {
		return A
	}

	// Split the arrays in half
	var left = make([]int, arrayLength/2)
	var right = make([]int, arrayLength/2+1)

	// Recurse to split in half
	left = MergeSort(left)
	right = MergeSort(right)

	// After the recursive split has completed and returned, merge the two arrays back up the stack
	return Merge(left, right)
}

func Merge(left []int, right []int) []int {

	var sorted []int

	for len(left) > 0 && len(right) > 0 {

	}

	return sorted
}
