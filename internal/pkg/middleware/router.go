package middleware

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetDefaultRouter() *gin.Engine {
	var env string

	if env = os.Getenv("APP_ENV"); len(env) == 0 {
		env = "dev"
	}

	if env != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "*"},
	}))

	return r
}
