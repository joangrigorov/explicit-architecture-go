package forms

type SignUp struct {
	Email     string `form:"email" json:"email" binding:"required,email"`
	Username  string `form:"username" json:"username" binding:"required"`
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
}
