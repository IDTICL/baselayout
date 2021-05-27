package user

type User struct {
	Username *string `json:"username" validate:"required"`
	Password *string `json:"password" validate:"required"`
	Role     *string `json:"role"`
	Age      *int    `json:"age"`
}
