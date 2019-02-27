package main

import (
	"fmt"
	"sort"
)

func main() {
	mArray := []int{1, 2, 3, 4, 5}

	fmt.Printf("Original array: %v \n", mArray)
	//reverse with func
	reverse(mArray)
	//reverse with
	sort.Sort(sort.Reverse(sort.IntSlice(mArray)))
	fmt.Printf("Reverse by sort: %v \n", mArray)
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	fmt.Printf("Reverse by func: %v \n", arr)
}
