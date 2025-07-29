package main

import (
	"app/internal/infrastructure/bootstrap"
)

func main() {
	bootstrap.
		NewApp().
		Run()
}
