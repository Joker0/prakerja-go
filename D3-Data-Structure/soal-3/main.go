package main

import (
	"fmt"
	"strconv"
)

func munculSekali(input string) []int {
	countMap := make(map[int]int)

	for _, char := range input {
		if char >= '0' && char <= '9' {
			num, _ := strconv.Atoi(string(char))
			countMap[num]++
		}
	}

	var uniqueNumbers []int

	for num, count := range countMap {
		if count == 1 {
			uniqueNumbers = append(uniqueNumbers, num)
		}
	}

	return uniqueNumbers
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
