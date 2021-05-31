package presenter

import (
	byd "fish/internal/pkg/structure/byd"
	"fish/internal/pkg/structure/user"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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

func GetOData(c *gin.Context) {
	//var s []byd.SapDemo
	var username string = "ADMINISTRATION01"
	var passwd string = "Welcome1"
	url := byd.BYD_URL_PRE + "356727" + byd.BYD_URL_MID + byd.BYD_URL_DEV_API + byd.BYD_FORMAT_JSON
	log.Println(url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	//req, err := http.NewRequest("GET", "https://my356727.sapbydesign.com/sap/byd/odata/cust/v1/test_get/CustomerInvoiceCustomerInvoiceCollection?format=json", nil)
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	//log.Println(bodyText)
	b := string(bodyText)
	//json.Unmarshal(bodyText, &s)
	log.Println(b)
	//c.JSON(http.StatusOK, s)
	//s := string(bodyText)
	//fmt.Println(bodyText)
	//return s
}
