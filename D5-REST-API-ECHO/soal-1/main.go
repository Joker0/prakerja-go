package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct untuk merepresentasikan data post dari API
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// URL API yang akan diambil data dari
	apiURL := "https://jsonplaceholder.typicode.com/posts"

	// Membuat permintaan HTTP GET ke API
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Gagal mengambil data dari API:", err)
		return
	}
	defer response.Body.Close()

	// Mengurai data JSON yang diterima dari API
	var posts []Post
	err = json.NewDecoder(response.Body).Decode(&posts)
	if err != nil {
		fmt.Println("Gagal mengurai data JSON:", err)
		return
	}

	// Menampilkan data yang diperoleh dari API
	fmt.Println("Data dari API:")
	for _, post := range posts {
		fmt.Printf("User ID: %d\n", post.UserID)
		fmt.Printf("ID: %d\n", post.ID)
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Body: %s\n", post.Body)
		fmt.Println()
	}
}
