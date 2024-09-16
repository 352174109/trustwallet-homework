package dal

import (
	"context"
	"sync"
)

type SubscribeDal struct {
	data map[string]struct{}

	lock sync.RWMutex
}

func NewSubscribeDal() (*SubscribeDal, error) {
	return &SubscribeDal{data: make(map[string]struct{})}, nil
}

func (m *SubscribeDal) Subscribe(ctx context.Context, key string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.data[key] = struct{}{}
	return nil
}

func (m *SubscribeDal) Subscribed(ctx context.Context, key string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	_, ok := m.data[key]

	return ok
}
