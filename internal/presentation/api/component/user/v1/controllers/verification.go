package controllers

import (
	"app/internal/core/port/cqrs"
	"app/internal/infrastructure/framework/http"
)

type Verification struct {
	commandBus cqrs.CommandBus
	queryBus   cqrs.QueryBus
}

func NewVerification(commandBus cqrs.CommandBus, queryBus cqrs.QueryBus) *Verification {
	return &Verification{commandBus: commandBus, queryBus: queryBus}
}

func (c *Verification) PreflightValidate(ctx http.Context) {
	verificationID := ctx.ParamString("id")
	token := ctx.Query("token")
}
