package get_verification_preflight

type Query struct {
	verificationID string
	Token          string
}

func NewQuery(
	verificationID string,
	csrfToken string,
) *Query {
	return &Query{verificationID: verificationID, Token: csrfToken}
}
