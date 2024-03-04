package routes

import (
	"QRCodeAPI/utils"
	"fmt"
	"github.com/skip2/go-qrcode"
	"net/http"
	"strconv"
	"strings"
)

func QRHandler(w http.ResponseWriter, r *http.Request) {
	// Get data from query string
	data := r.URL.Query().Get("data")
	// If data is empty, return error message
	if data == "" {
		data = "Missing data parameter in query string"
	}
	// If data is too long, return error message
	// The maximum capacity is 2,953 bytes, 4,296 alphanumeric characters, 7,089 numeric digits, or a combination of these.
	if len(data) > 2953 {
		utils.JSONResponse(
			w,
			http.StatusBadRequest,
			[]byte("{\"message\": \"Data parameter is too long\"}"),
		)
		return
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
		utils.JSONResponse(
			w,
			http.StatusInternalServerError,
			[]byte(fmt.Sprintf("{\"message\": \"%s\"}", err.Error())),
		)
		return
	}

	// Write response to client with image/png content type header and png data
	w.Header().Set("Content-Type", "image/png")
	if _, err = w.Write(png); err != nil {
		utils.JSONResponse(
			w,
			http.StatusInternalServerError,
			[]byte(fmt.Sprintf("{\"message\": \"%s\"}", err.Error())),
		)
		return
	}
}
