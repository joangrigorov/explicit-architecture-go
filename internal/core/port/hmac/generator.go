package hmac

type Generator interface {
	Generate(message string) (hmac string, secret string, err error)
}
