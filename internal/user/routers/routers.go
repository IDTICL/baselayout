package routers

import (
	user "baselayout/internal/user/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"},
		AllowHeaders: []string{"Authorization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host"},
		MaxAge: 12 * time.Hour,
	},
	))

	authorized := r.Group("/v1")
	{
		authorized.GET("/users", user.All)
		authorized.POST("/users", user.Add)
		authorized.GET("/users/:id", user.FetchOne)
		authorized.DELETE("/users/:id", user.Delete)
	}
}
