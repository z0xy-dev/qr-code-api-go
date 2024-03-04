package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"net/http"
	"strconv"
	"strings"
)

func jsonResponse(w http.ResponseWriter, code int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(data); err != nil {
		return
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/qr" {
		qrHandler(w, r)
	} else {
		jsonResponse(w, http.StatusNotFound, []byte(`{"message": "Not found"}`))
		return
	}
}

func qrHandler(w http.ResponseWriter, r *http.Request) {
	// Get data from query string
	data := r.URL.Query().Get("data")
	if data == "" {
		data = "Missing data parameter in query string"
	}

	// Get recovery level from query string
	recoveryLevelStr := r.URL.Query().Get("level")
	recoveryLevel := qrcode.Medium
	switch strings.ToLower(recoveryLevelStr) {
	case "low":
		recoveryLevel = qrcode.Low
	case "medium":
		recoveryLevel = qrcode.Medium
	case "high":
		recoveryLevel = qrcode.High
	case "highest":
		recoveryLevel = qrcode.Highest
	}

	// Get size from query string
	sizeStr := r.URL.Query().Get("size")
	size, sizeErr := strconv.Atoi(sizeStr)
	if sizeErr != nil {
		size = 256
	} else {
		if size < 256 {
			size = 256
		}
		if size > 1024 {
			size = 1024
		}
	}

	// Generate QR code
	png, err := qrcode.Encode(data, recoveryLevel, size)
	if err != nil {
		jsonResponse(
			w,
			http.StatusInternalServerError,
			[]byte(fmt.Sprintf("{\"message\": \"%s\"}", err.Error())),
		)
		return
	}

	// Write response to client with image/png content type header and png data
	w.Header().Set("Content-Type", "image/png")
	if _, err = w.Write(png); err != nil {
		jsonResponse(
			w,
			http.StatusInternalServerError,
			[]byte(fmt.Sprintf("{\"message\": \"%s\"}", err.Error())),
		)
		return
	}
}

func main() {
	http.HandleFunc("/", requestHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	} else {
		fmt.Print("Server is running on port 8080")
	}
}
