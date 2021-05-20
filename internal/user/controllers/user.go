package controllers

import (
	"baselayout/internal/pkg/dao"
	"baselayout/internal/user/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func All(c *gin.Context) {

	db := dao.Open()
	var user model.User
	var users []model.User

	db.Model(&user).Scan(&users)

	c.JSON(http.StatusOK, users)
}

func Add(c *gin.Context) {
	db := dao.Open()

	var person model.User
	if err := c.BindJSON(&person); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if err := db.Create(&person).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, person)

}

func FetchOne(c *gin.Context) {
	id := c.Param("id")

	db := dao.Open()

	var person model.User
	if err := db.Model(&person).Where("id = ?", id).First(&person).Error; err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, person)
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	db := dao.Open()
	var person model.User

	if err := db.Model(&person).Where("id = ?", id).Error; err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
