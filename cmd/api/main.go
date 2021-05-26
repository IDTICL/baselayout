package main

import (
	"github.com/gin-gonic/gin"
	pgx "idticl.app/internal/pkg/dao"
	"idticl.app/internal/user/presenters"
	"log"
	"net/http"
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
		authorized.POST("/users", presenters.Create)
		authorized.GET("/users", presenters.AllUser)
	}

	log.Fatal(http.ListenAndServe(":8080", router))

}
