package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Value of z in %v loop is: %v \n", i+1, z)
	}
	return z
}

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello my name is Renan")
	fmt.Println("The pi number is", math.Pi)
	fmt.Println("The five plus five is", add(5, 5))
	fmt.Println("The squared root of 2 is ", Sqrt(2))

}
