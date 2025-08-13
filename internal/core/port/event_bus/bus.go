package event_bus

type Event interface {
	ID() string
}

type EventBus interface {
	Publish(Event) error
}
