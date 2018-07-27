package headFirstPatterns

import (
	"sync"

	"github.com/pkg/errors" // add context to the error stack
	uuid "github.com/satori/go.uuid"
)

type Subject interface {
	RegisterObserver(os Observer) error
	RemoveObserver(os Observer) error
	NumOfObservers() int
	ObserverExists(os Observer) bool

	/*
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
	mutex             *sync.Mutex
	observers         []chan Event
	maxNumOfObservers int
}

type observer struct {
	ch chan Event
}

// implement Subject
func NewSubject() Subject {
	aSubject := &subject{
		mutex: &sync.Mutex{},
	}
	return aSubject
}

// RegisterObserver supports concurrent observer adds. It also
// ensures that no duplicated observers are added.
func (s *subject) RegisterObserver(os Observer) error {
	if os == nil {
		return errors.New("nil obServer")
	}
	if !s.ObserverExists(os) {
		s.addObserver(os)
	}
	return nil
}

func (s *subject) RemoveObserver(os Observer) error {
	if os == nil {
		return errors.New("nil observer")
	}
	for i, ch := range s.observers {
		if os.Channel() == ch {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
	return nil
}

func (s *subject) NumOfObservers() int {
	var count int
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, ch := range s.observers {
		if ch != nil {
			count++
		}
	}
	return count
}

func (s *subject) ObserverExists(os Observer) bool {
	if os == nil {
		return false
	}
	if s.NumOfObservers() == 0 {
		return false
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, ch := range s.observers {
		if os.Channel() == ch {
			return true
		}
	}
	return false
}

func (s *subject) addObserver(os Observer) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.observers = append(s.observers, os.Channel())
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
