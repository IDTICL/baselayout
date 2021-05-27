package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pgx "idticl.app/internal/pkg/dao"
	"idticl.app/internal/user/presenter"
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
	}

	log.Fatal(http.ListenAndServe(":8080", router))

}
