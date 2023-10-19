package main

import (
	"fmt"

	"github.com/joker0/prakerja-go/D2-Basic-Programming/perhitungan"
)

func main() {
	var a, b, t float64 = 4.00, 5.00, 3.00
	fmt.Println("Area of Trapezoid: ", perhitungan.AreaOfTrapezoid(a, b, t))

	fmt.Println("Is Mutiple of Seven: ", perhitungan.IsMutipleOfSeven(14))
	fmt.Println("Is Mutiple of Seven: ", perhitungan.IsMutipleOfSeven(8))

	fmt.Println("Is Prime: ", perhitungan.IsPrime(5))
	fmt.Println("Is Prime: ", perhitungan.IsPrime(9))
	fmt.Println("Is Prime: ", perhitungan.IsPrime(31))
	fmt.Println("Is Prime: ", perhitungan.IsPrime(33))

}
