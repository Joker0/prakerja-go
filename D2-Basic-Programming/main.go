package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/joker0/prakerja-go/D2-Basic-Programming/perhitungan"
)

// Function to generate an array of random numbers
func generateRandomArray(size, min, max int) []int {
	rand.Seed(time.Now().UnixNano())
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(max-min+1) + min
	}
	return array
}

// Function to loop through an array and perform some operation
func loopArray(arr []int, code int) {
	if code == 1 {
		for _, value := range arr {
			fmt.Println("Number", value, "is prime: ", perhitungan.IsPrime(value))
		}
	} else {
		for _, value := range arr {
			fmt.Println("Number", value, "is multiple of 7: ", perhitungan.IsMutipleOfSeven(value))
		}
	}

}
func main() {
	var a, b, t float64 = 4.00, 5.00, 3.00
	fmt.Println("Area of Trapezoid: ", perhitungan.AreaOfTrapezoid(a, b, t))

	primeNumbers := generateRandomArray(5, 1, 100)
	loopArray(primeNumbers, 1)

	multipleOfSeven := generateRandomArray(5, 1, 50)
	loopArray(multipleOfSeven, 2)

}
