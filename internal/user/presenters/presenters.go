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

	model := models.UserModel{Read: db.ReadWrite.Read, Write: db.ReadWrite.Write}
	model.Insert()
}
