package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	sum := add(42, 13)
	fmt.Println(sum)
}
