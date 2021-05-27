package presenter

import (
	"fmt"
	"net/http"

	"fish/internal/pkg/structure/user"
	"fish/internal/user/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Create(c *gin.Context) {
	user := user.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "JSON format not valid",
		})

		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Json Format Error",
		})

		return
	}

	if err := model.Insert(c, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Server internal error",
		})
		return
	}

	c.Status(http.StatusAccepted)
}

func AllUser(c *gin.Context) {

	if err := model.FindAll(c); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Server internal error",
		})
		return
	}

	c.Status(http.StatusAccepted)
}
