package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	pgx "idticl.app/internal/pkg/dao"
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
	router.Use(cors.Default())

	authorized := router.Group("/v1")
	{
		authorized.GET("/users", user.All)
		authorized.POST("/users", user.Add)
		authorized.GET("/users/:id", user.FetchOne)
		authorized.DELETE("/users/:id", user.Delete)
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}
