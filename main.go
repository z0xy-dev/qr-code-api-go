package main

import (
	"QRCodeAPI/routes"
	"QRCodeAPI/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", routes.RouteRequest)
	port := 3690

	fmt.Printf("Server is running on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), utils.AllowCORS(nil)))
}
