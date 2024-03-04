package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"net/http"
)

func qrHandler(w http.ResponseWriter, r *http.Request) {
	png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(png)
}

func main() {
	http.HandleFunc("/", qrHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	} else {
		fmt.Println("Server is running on port 8080")
	}
}
