package uuid

type Generator interface {
	Generate() string
}
