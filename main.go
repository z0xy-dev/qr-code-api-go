package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/qr" {
		qrHandler(w, r)
	} else {
		// JSON response
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{"error": "Not found"}`))
		if err != nil {
			return
		}
	}
}

func qrHandler(w http.ResponseWriter, r *http.Request) {
	png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	_, err = w.Write(png)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", requestHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	} else {
		fmt.Println("Server is running on port 8080")
	}
}
