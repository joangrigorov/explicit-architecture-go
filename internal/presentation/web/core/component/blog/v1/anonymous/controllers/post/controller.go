package post

import (
	. "app/internal/core/component/blog/application/repositories"
	. "github.com/go-playground/universal-translator"
)

type Controller struct {
	translator     Translator
	PostRepository PostRepository
}

func NewController(postRepository PostRepository, translator Translator) *Controller {
	return &Controller{PostRepository: postRepository, translator: translator}
}
