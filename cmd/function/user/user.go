package main

import (
	"github.com/apex/gateway"
	pgx "idticl.app/internal/pkg/dao"
	"idticl.app/internal/user/routers"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := pgx.New(); err != nil {
		panic(err)
	}

	defer pgx.Pool.Close()

	mode := os.Getenv("GIN_MODE")

	if mode == "release" {
		log.Fatal(gateway.ListenAndServe(":8080", routers.GetRouters()))
	} else {
		log.Fatal(http.ListenAndServe(":8080", routers.GetRouters()))
	}
}
