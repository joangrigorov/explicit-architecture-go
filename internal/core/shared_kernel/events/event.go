package events

import "reflect"

type EventID string

func (i EventID) String() string {
	return string(i)
}

type Event interface {
	ID() EventID
}

func makeEventID(e Event) EventID {
	t := reflect.TypeOf(e)

	// If e is a pointer, get the element type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// return package path + struct name
	return EventID(t.PkgPath() + "." + t.Name())
}
