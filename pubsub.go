package pubsub

import (
	"sync"
)

// PubSub The guts of this
type PubSub struct {
	Name   string
	subs   []chan string
	mu     sync.RWMutex
	closed bool
}

// NewPubSub create a new Pub Sub
func NewPubSub(name string) *PubSub {

	return &PubSub{
		Name: name,
	}
}

// Subscribe Creates a channel for subscriber
func (ps *PubSub) Subscribe() chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if ps.closed {
		return nil
	}

	listenChannel := make(chan string, 10)
	ps.subs = append(ps.subs, listenChannel)
	return listenChannel
}

// Close the service
func (ps *PubSub) Close() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if ps.closed {
		return
	}

	ps.closed = true

	listLength := len(ps.subs)
	for i := 0; i < listLength; i++ {
		if ps.subs[i] != nil {
			close(ps.subs[i])
			ps.subs[i] = nil
		}
	}
}

// Publish Sends a message to all subscribers
func (ps *PubSub) Publish(message string) {
	if ps.closed {
		return
	}

	listLength := len(ps.subs)
	for i := 0; i < listLength; i++ {
		if ps.subs[i] != nil {
			select {
			case ps.subs[i] <- message:
			default:
			}
		}
	}
}
