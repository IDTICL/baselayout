package user

type User struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
	Age      int    `json:"age"`
}

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

type House struct {
	City        string `json:"City"`
	Area        string `json:"Area"`
	ReadingRoom string `json:"Reading_Room"`
	Address     string `json:"Address"`
}

type Login struct {
	Username *string `json:"username" validate:"required"`
	Password *string `json:"password" validate:"required"`
}
