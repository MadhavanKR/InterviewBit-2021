package main

import "fmt"

func main() {
	set := make(map[int]bool)
	set[2] = true
	set[1] = true
	set[5] = true
	fmt.Println(set)
}
