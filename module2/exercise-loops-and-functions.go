package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	var prevGuess float64

	for i:=0; i<10; i++ {
		z -= (z*z - x) / (2*z)

		if prevGuess == z {
			fmt.Println("matching previous guess", prevGuess, z)
		} else {
			fmt.Print("Getting closer")
		}

		prevGuess = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(49))
	fmt.Println(math.Sqrt(49))
}

