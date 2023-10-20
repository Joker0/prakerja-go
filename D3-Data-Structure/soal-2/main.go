package main

import "fmt"

func mapping(slice []string) map[string]int {
	counts := make(map[string]int)
	for _, s := range slice {
		counts[s]++
	}
	return counts
}

func main() {
	// Contoh slice
	data := []string{"apel", "jeruk", "anggur", "apel", "anggur", "apel"}
	data2 := []string{}
	data3 := []string{"merah", "biru", "hijau", "merah", "merah", "hijau"}

	fmt.Println(mapping(data))
	fmt.Println(mapping(data2))
	fmt.Println(mapping(data3))
}
