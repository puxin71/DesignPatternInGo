package headFirstPatterns

import (
	"github.com/pkg/errors" // add context to the error stack
	uuid "github.com/satori/go.uuid"
)

const (
	MaxNumOfObservers = int(10)
)

type Subject interface {
	RegisterObserver(os Observer) error
	/*
		RemoveObserver()
		NotifyObserver()
	*/
}

type Observer interface {
	Channel() chan Event
	/*
		Update()
	*/
}

type Event struct {
	id      uuid.UUID
	payload interface{}
}

type subject struct {
	observers         []chan Event
	maxNumOfObservers int
}

type observer struct {
	ch chan Event
}

// implement Subject
func NewSubject(maxNumOfObservers int) (Subject, error) {
	if maxNumOfObservers <= 0 {
		return nil, errors.Errorf("invalid max number of observers. number: %d", maxNumOfObservers)
	}
	aSubject := &subject{
		observers: make([]chan Event, maxNumOfObservers),
	}
	return aSubject, nil
}

func (s *subject) RegisterObserver(os Observer) error {
	if os == nil {
		return errors.New("nil obServer")
	}
	s.observers = append(s.observers, os.Channel())
	return nil
}

// implement Observer
func NewObserver() Observer {
	return &observer{
		ch: make(chan Event),
	}
}

func (os *observer) Channel() chan Event {
	return os.ch
}
