package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// URL API tempat Anda ingin menyimpan data postingan
	apiURL := "https://jsonplaceholder.typicode.com/posts"

	// Data postingan baru yang akan Anda kirim
	newPost := Post{
		UserID: 1,
		ID:     101, // ID postingan harus unik
		Title:  "New Post Title",
		Body:   "This is the content of the new post.",
	}

	// Mengubah data postingan menjadi format JSON
	postData, err := json.Marshal(newPost)
	if err != nil {
		fmt.Println("Gagal mengubah data menjadi JSON:", err)
		return
	}

	// Mengirim permintaan HTTP POST ke API
	response, err := http.Post(apiURL, "application/json", bytes.NewBuffer(postData))
	if err != nil {
		fmt.Println("Gagal mengirim permintaan POST:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusCreated {
		fmt.Println("Data postingan berhasil disimpan di server.")
	} else {
		fmt.Println("Gagal menyimpan data postingan di server. Kode status:", response.Status)
	}
}
