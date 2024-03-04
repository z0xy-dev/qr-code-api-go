package main

import (
	"QRCodeAPI/routes"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", routes.RouteRequest)
	port := 3690
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
	fmt.Println("Server is running on port", port)
}
