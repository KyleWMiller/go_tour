package main

import (
	"golang.org/x/tour/pic"
	"fmt"
)

func Pic(dx, dy int) [][]uint8 {
	fmt.Println("dx Value: ", dx)
	fmt.Printf("dx Type: %t", dx)
	fmt.Println("dy Value: ", dy)
	fmt.Printf("dy Type: %t", dy)
//	y := make([]uint8, dy)

//	for i, v := range y {
//		y[i] = make([]uint8, dx)
//	}

	y := make([][]uint8, 9)
	return y
}

func main() {
	pic.Show(Pic)
}
