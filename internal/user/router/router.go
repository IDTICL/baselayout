package router

import (
	"fish/internal/pkg/middleware"
	"fish/internal/user/presenter"

	"github.com/gin-gonic/gin"
)

func GetRouters() *gin.Engine {
	r := middleware.GetDefaultRouter()

	authorized := r.Group("/v1")
	{
		authorized.POST("/user", presenter.Create)
		authorized.GET("/user", presenter.AllUser)
		authorized.GET("/sapget", presenter.SAPGet)
	}

	return r
}
