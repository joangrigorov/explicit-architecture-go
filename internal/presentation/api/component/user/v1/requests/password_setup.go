package requests

type PasswordSetup struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required"`
}
