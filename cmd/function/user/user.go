package main

import (
	"github.com/apex/gateway"
	"idticl.app/internal/user/routers"
	"log"
	"net/http"
	"os"
)

func main() {

	mode := os.Getenv("GIN_MODE")

	if mode == "release" {
		log.Fatal(gateway.ListenAndServe(":8080", routers.GetRouters()))
	} else {
		log.Fatal(http.ListenAndServe(":8080", routers.GetRouters()))
	}
}
