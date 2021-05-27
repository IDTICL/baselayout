package router

import (
	"os"

	"fish/internal/user/presenter"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouters() *gin.Engine {
	var env string

	if env = os.Getenv("APP_ENV"); len(env) == 0 {
		env = "dev"
	}

	if env != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "*"},
	}))

	authorized := router.Group("/v1")
	{
		authorized.POST("/user", presenter.Create)
		authorized.GET("/user", presenter.AllUser)
	}

	return router
}
