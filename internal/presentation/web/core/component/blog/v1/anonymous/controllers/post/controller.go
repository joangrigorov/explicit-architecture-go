package post

import (
	. "app/internal/core/component/activity/application/repositories"
	. "github.com/go-playground/universal-translator"
)

type Controller struct {
	translator     Translator
	PostRepository ActivityRepository
}

func NewController(postRepository ActivityRepository, translator Translator) *Controller {
	return &Controller{PostRepository: postRepository, translator: translator}
}
