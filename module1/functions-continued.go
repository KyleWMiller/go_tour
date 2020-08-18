package main

import "fmt"

func add(x, y int) int {
	return x + y 
}

func main() {
	sum := add(700, 40)
	fmt.Println(sum)
}
