package headFirstPatterns

import (
	uuid "github.com/satori/go.uuid"
)

type Subject interface {
	RegisterObserver()
	RemoveObserver()
	NotifyObserver()
}

type Observer interface {
	Update()
}

type Event struct {
	id      uuid.UUID
	payload interface{}
}

type subject struct {
	observers []chan Event
}
