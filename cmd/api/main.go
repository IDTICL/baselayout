package main

import (
	"log"
	"net/http"

	byd "fish/internal/byd/presenter"
	pgx "fish/internal/pkg/dao"
	user "fish/internal/user/presenter"

	"fish/internal/pkg/middleware/auth"

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
	authorized.Use(auth.AuthRequired)
	{
		authorized.POST("/user", user.Create)
		authorized.GET("/user", user.AllUser)

		authorized.GET("/sapget", user.SAPGet)

		authorized.GET("/byd/getodata", byd.GetOData)
	}
	router.POST("/login", user.Login)

	log.Fatal(http.ListenAndServe(":8080", router))

}
