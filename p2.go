package main

import "fmt"

func main() {
	marray := []int{1, 2, 3, 4, 5}

	sum(marray)
	multiply(marray)
}

func sum(arr []int) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}

func multiply(arr []int) {
	mul := 1
	for _, j := range arr {
		mul *= j
	}
	fmt.Println(mul)
}
