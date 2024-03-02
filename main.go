package main

import (
	"fmt"
	"net/http"
)

func qrHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/api/qr", qrHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	} else {
		fmt.Println("Server is running on port 8080")
	}
}
