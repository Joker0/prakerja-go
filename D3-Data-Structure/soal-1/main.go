package main

import "fmt"

func mergeArrays(array1, array2 []string) []string {
	// Menggabungkan kedua array
	mergedArray := append(array1, array2...)

	// Membuat map untuk menghitung kemunculan setiap nama
	nameCounts := make(map[string]int)

	// Menyaring nama yang hanya muncul satu kali
	var uniqueNames []string
	for _, name := range mergedArray {
		if nameCounts[name] == 0 {
			uniqueNames = append(uniqueNames, name)
		}
		nameCounts[name]++
	}

	return uniqueNames
}

func main() {
	array1 := []string{"Ryu", "Ken", "Dalsim", "Blanka"}
	array2 := []string{"Vega", "Chunli", "Ken", "Guile"}

	fmt.Println(mergeArrays(array1, array2))
}
