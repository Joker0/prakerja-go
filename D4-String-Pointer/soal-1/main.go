package main

import (
	"fmt"
)

type Car struct {
	Type    string
	FuelInL float64
}

func (c *Car) EstimateDistance() float64 {
	const fuelEfficiency = 1.5 // Efisiensi bahan bakar dalam L/mil

	distance := c.FuelInL / fuelEfficiency
	return distance
}

func main() {
	myCar := Car{
		Type:    "Sedan",
		FuelInL: 10.0, // Misalnya, 10 liter bahan bakar terisi
	}

	estimatedDistance := myCar.EstimateDistance()
	fmt.Printf("Dengan %s, perkiraan jarak yang bisa ditempuh adalah %.2f mil.\n", myCar.Type, estimatedDistance)
}
