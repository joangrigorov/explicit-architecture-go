package get_verification_preflight

import "encoding/json"

type Query struct {
	verificationID string
	Token          string
}

func (q Query) LogBody() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"verificationID": q.verificationID,
		"Token":          "-",
	})
}

func NewQuery(
	verificationID string,
	csrfToken string,
) Query {
	return Query{verificationID: verificationID, Token: csrfToken}
}
