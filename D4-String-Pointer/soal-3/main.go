package main

import "fmt"

func findMinMax(numbers [6]int) (int, int) {
	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

func main() {
	var inputNumbers [6]int

	fmt.Println("Masukkan 6 angka:")

	for i := 0; i < 6; i++ {
		fmt.Printf("Angka ke-%d: ", i+1)
		fmt.Scan(&inputNumbers[i])
	}

	min, max := findMinMax(inputNumbers)

	fmt.Printf("Nilai Minimum: %d\n", min)
	fmt.Printf("Nilai Maksimum: %d\n", max)
}
