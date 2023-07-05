package main

import (
	"math"
	"sync"
)

type Message struct {
	data      []byte
	timestamp int
}

type Process struct {
	clock int
	mtx   sync.Mutex
}

func New() *Process {
	return &Process{
		clock: 0,
		mtx:   sync.Mutex{},
	}
}

func (p *Process) tick() {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	p.clock++
}

func (p *Process) Timestamp() int {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	return p.clock
}

// func (p *Process) () {}

func (p *Process) LocalEvent() {
	p.tick()
}

func (p *Process) SendMessage() error {
	p.tick()
	// t := p.Timestamp()
	// send message with t

	return nil
}

func (p *Process) RecieveMessage(m *Message) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	p.clock = int(math.Max(float64(p.clock), float64(m.timestamp)) + 1)
}
