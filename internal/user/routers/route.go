package routers

import (
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	UserRoute(r)

	r.Run()
}
