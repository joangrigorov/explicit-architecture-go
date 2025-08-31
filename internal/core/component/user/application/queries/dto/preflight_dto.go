package dto

type PreflightDTO struct {
	ValidCSRF   bool   `json:"valid_csrf"`
	MaskedEmail string `json:"masked_email"`
	Expired     bool   `json:"expired"`
	UserID      string `json:"-"`
}
