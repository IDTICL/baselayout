package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type Base struct {
	gorm.Model
}

type DeletedAt sql.NullTime

type User struct {
	Base
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Age      int    `json:"age"`
}

type UpdateUser struct {
	Base
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Age      int    `json:"age"`
}
