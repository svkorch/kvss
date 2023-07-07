package kvss

import "sync"

type Storager interface {
	Add(k, v string)
	Get(k string) (v string, ok bool)
}

type KVStorage struct {
	mu sync.RWMutex
	mp map[string]string
}

func New() *KVStorage {
	x := &KVStorage{
		mp: make(map[string]string),
	}

	var _ Storager = x

	return x
}

func (s *KVStorage) Add(k, v string) {
	s.mu.RLock()
	s.mp[k] = v
	s.mu.RUnlock()
}

func (s *KVStorage) Get(k string) (v string, ok bool) {
	s.mu.Lock()
	v, ok = s.mp[k]
	s.mu.Unlock()

	return v, ok
}
