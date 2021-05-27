package main

import (
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	pgx "idticl.app/internal/pkg/dao"
	"idticl.app/internal/user/router"
)

func main() {
	if err := pgx.New(); err != nil {
		panic(err)
	}

	defer pgx.Pool.Close()

	mode := os.Getenv("GIN_MODE")

	if mode == "release" {
		log.Fatal(gateway.ListenAndServe(":8080", router.GetRouters()))
	} else {
		log.Fatal(http.ListenAndServe(":8080", router.GetRouters()))
	}
}
