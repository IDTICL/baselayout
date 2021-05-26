package internal

import (
	"github.com/gin-gonic/gin"
	"idticl.app/internal/user/routers"
)

func Routes() {

	r := gin.Default()

	routers.GetRouters(r)

	r.Run()
}
