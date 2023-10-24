package main

import (
	"fmt"
	"net/http"
)

func main() {
	// URL API tempat Anda ingin menghapus data
	apiURL := "https://jsonplaceholder.typicode.com/posts/1" // Ganti dengan URL yang sesuai

	// Membuat permintaan HTTP DELETE ke API
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", apiURL, nil)
	if err != nil {
		fmt.Println("Gagal membuat permintaan DELETE:", err)
		return
	}

	// Mengirim permintaan DELETE ke server
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Gagal mengirim permintaan DELETE:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNoContent {
		fmt.Println("Data berhasil dihapus dari server.")
	} else {
		fmt.Println("Gagal menghapus data dari server. Kode status:", response.Status)
	}
}
