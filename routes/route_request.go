package routes

import (
	"QRCodeAPI/utils"
	"net/http"
)

func RouteRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/qr" {
		QRHandler(w, r)
	} else {
		utils.JSONResponse(w, http.StatusNotFound, []byte(`{"message": "Not found"}`))
		return
	}
}
