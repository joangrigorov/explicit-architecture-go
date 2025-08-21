package controllers

import (
	"github.com/gin-gonic/gin"
)

type Home struct{}

func NewHome() *Home {
	return &Home{}
}

func (h *Home) Index(c *gin.Context) {
	c.HTML(200, "home/index", gin.H{
		"Title":   "Welcome Page",
		"Heading": "Hello from Gin!",
		"Message": "This page is rendered with Gin and html/template.",
	})
}
