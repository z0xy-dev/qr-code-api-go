package routes

import (
	"QRCodeAPI/utils"
	"net/http"
)

func RouteRequest(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/qr":
		QRHandler(w, r)
	default:
		utils.JSONResponse(w, http.StatusNotFound, []byte(`{"message": "Not found"}`))
	}
}
