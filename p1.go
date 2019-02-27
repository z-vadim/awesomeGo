package main

import "fmt"

func main() {
	name := "qwe"
	sayhey(name)
}

func sayhey(name string) {
	fmt.Printf("Hey %s", name)
}
