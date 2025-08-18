package commands

import "context"

type Command interface {
	LogBody() ([]byte, error)
}

type CommandHandler[C Command] interface {
	Handle(context.Context, C) error
}

type CommandBus interface {
	Dispatch(context.Context, Command) error
}
