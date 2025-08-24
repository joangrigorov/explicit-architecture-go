package hmac

type Verifier interface {
	Verify(message, secret, signature string) bool
}
