// Package events is a small in-process publish/subscribe bus.
package events

import "sync"

// Event is something that happened.
type Event struct {
	Topic   string
	Payload any
}

// Bus delivers events to the subscribers of their topic.
type Bus struct {
	mu   sync.RWMutex
	subs map[string][]func(Event)
}

func NewBus() *Bus { return &Bus{subs: map[string][]func(Event){}} }

// Subscribe registers fn for a topic.
func (b *Bus) Subscribe(topic string, fn func(Event)) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.subs[topic] = append(b.subs[topic], fn)
}

// Publish delivers e to every subscriber of its topic, in subscription order,
// and returns how many received it.
func (b *Bus) Publish(e Event) int {
	b.mu.RLock()
	fns := append([]func(Event){}, b.subs[e.Topic]...)
	b.mu.RUnlock()
	for _, fn := range fns {
		fn(e)
	}
	return len(fns)
}
