package presenter

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/goccy/go-json"

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

func SAPGet(c *gin.Context) {
	//var sap []user.SapDemo
	var h []user.House
	//var username string = "ADMINISTRATION01"
	//var passwd string = "Welcome1"
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://data.ntpc.gov.tw/api/datasets/4A03827A-588B-4058-AB21-EC02283E2BB7/json?page=0&size=100", nil)
	//req, err := http.NewRequest("GET", "https://my356727.sapbydesign.com/sap/byd/odata/cust/v1/test_get/CustomerInvoiceCustomerInvoiceCollection?format=json", nil)
	//req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyText, &h)
	c.JSON(http.StatusOK, h)
	//s := string(bodyText)
	//fmt.Println(bodyText)
	//return s
}
