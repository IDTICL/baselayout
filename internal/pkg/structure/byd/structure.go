package byd

const (
	BYD_URL_PRE      = "https://my"
	BYD_URL_DEV_SITE = "356727"
	BYD_URL_MID      = ".sapbydesign.com/sap/byd/odata/cust/v1/"
	BYD_URL_DEV_API  = "test_get/CustomerInvoiceCustomerInvoiceCollection"
	BYD_FORMAT_JSON  = "?$format=json"
)

type SapDemo struct {
	Metadata struct {
		URI  string `json:"uri"`
		Type string `json:"type"`
	} `json:"__metadata"`
	Objectid         string `json:"ObjectID"`
	ID               string `json:"ID"`
	Date             string `json:"Date"`
	Totalgrossamount string `json:"TotalGrossAmount"`
	Currencycode     string `json:"currencyCode"`
	Currencycodetext string `json:"currencyCodeText"`
}
