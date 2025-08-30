package forms

type PasswordSetup struct {
	Password        string `form:"password" json:"password" binding:"required,min=8"`
	ConfirmPassword string `form:"confirm_password" json:"-" binding:"required,eqfield=Password"`
	VerificationID  string `form:"id" json:"-" binding:"required"`
	Token           string `form:"token" json:"token" binding:"required"`
}
