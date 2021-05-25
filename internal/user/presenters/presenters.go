package presenters

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"idticl.app/internal/pkg/structure/users"
	"idticl.app/internal/user/models"
	"net/http"
)

func Create(c *gin.Context) {
	user := users.User{}

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

	if err := models.Insert(c, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Server internal error",
		})
		return
	}

	c.Status(http.StatusAccepted)
}
