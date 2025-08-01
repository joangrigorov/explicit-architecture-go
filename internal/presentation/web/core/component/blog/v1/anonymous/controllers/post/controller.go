package post

import (
	. "app/internal/core/component/blog/application/repositories"
)

type Controller struct {
	PostRepository PostRepository
}

func NewController(postRepository PostRepository) *Controller {
	return &Controller{PostRepository: postRepository}
}
