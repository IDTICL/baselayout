package routers

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"idticl.app/internal/user/presenters"
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
		authorized.POST("/users", presenters.Create)
		authorized.GET("/users", presenters.AllUser)
	}

	return router
}
