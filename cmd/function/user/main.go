package main

import (
	"log"
	"net/http"
	"os"

	pgx "fish/internal/pkg/dao"
	"fish/internal/user/router"

	"github.com/apex/gateway"
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
