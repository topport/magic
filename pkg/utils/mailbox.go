package utils

import (
	"sync"
)

type Mailbox struct {
	chNotify chan struct{}
	mu       sync.Mutex
	queue    []interface{}
	capacity uint64
}

func NewMailbox(capacity uint64) *Mailbox {
	queueCap := capacity
	if queueCap == 0 {
		queueCap = 100
	}
	return &Mailbox{
		chNotify: make(chan struct{}, 1),
		queue:    make([]interface{}, 0, queueCap),
		capacity: capacity,
	}
}

func (m *Mailbox) Notify() chan struct{} {
	return m.chNotify
}

func (m *Mailbox) Deliver(x interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.queue = append([]interface{}{x}, m.queue...)
	if uint64(len(m.queue)) > m.capacity && m.capacity > 0 {
		m.queue = m.queue[:len(m.queue)-1]
	}

	select {
	case m.chNotify <- struct{}{}:
	default:
	}
}

func (m *Mailbox) Retrieve() interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.queue) == 0 {
		return nil
	}
	x := m.queue[len(m.queue)-1]
	m.queue = m.queue[:len(m.queue)-1]
	return x
}

func (m *Mailbox) RetrieveAll() []interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	queue := m.queue
	m.queue = nil
	for i, j := 0, len(queue)-1; i < j; i, j = i+1, j-1 {
		queue[i], queue[j] = queue[j], queue[i]
	}
	return queue
}

func (m *Mailbox) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()

	queueCap := m.capacity
	if queueCap == 0 {
		queueCap = 100
	}
	m.queue = make([]interface{}, 0, queueCap)
}
