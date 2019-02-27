package main

import (
	"fmt"
	"time"
)

func main() {
	mArray := []int{1, 4, 3, 4, 5}

	histogram(mArray)
}

func histogram(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < arr[i]; j++ {
			fmt.Printf("*")
			time.Sleep(time.Second)
		}
		fmt.Printf("\n")
	}
}
