package main

import (
	"fmt"

	"github.com/joker0/prakerja-go/D2-Basic-Programming/perhitungan"
)

func main() {
	a, b, t := 4, 5, 3
	fmt.Println("Area of Trapezoid: ", perhitungan.AreaOfTrapezoid(a, b, t))
}
