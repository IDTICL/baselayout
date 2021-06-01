package byd

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetODataByte(host_number, username, password string) []byte {

	url := BYD_URL_PRE + host_number + BYD_URL_MID + BYD_URL_DEV_API + BYD_FORMAT_JSON
	log.Println(url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
