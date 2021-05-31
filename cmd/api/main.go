package main

import (
	"log"
	"net/http"

	pgx "fish/internal/pkg/dao"
	"fish/internal/user/presenter"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := pgx.New(); err != nil {
		log.Println(err)

		return
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	authorized := router.Group("/v1")
	{
		authorized.POST("/user", presenter.Create)
		authorized.GET("/user", presenter.AllUser)

		authorized.GET("/sapget", presenter.SAPGet)
	}

	log.Fatal(http.ListenAndServe(":8080", router))

}
