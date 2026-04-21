package main

import (
	"fmt"
	"sync"
	"time"
)

type Mutex struct {
	mu  sync.Mutex
	key map[string]int
}

func (m *Mutex) Increment(value string) {
	m.mu.Lock()
	m.key[value]++
	m.mu.Unlock()
}

func (m *Mutex) PrintValue(value string) int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.key[value]
}

func main() {
	v := Mutex{key: make(map[string]int)}
	for i := 1; i < 11; i++ {
		go v.Increment("Sanjana")
	}
	time.Sleep(2 * time.Second)
	fmt.Println(v.PrintValue("Sanjana"))
}
