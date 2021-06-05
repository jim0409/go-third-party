package memorystore

import (
	"fmt"

	"go-third-party/redis_related/redigo/session_advance/store"
)

// session/memory.go
type MemoryStore struct {
	sessions map[string]store.Session
}

func NewMemoryStore() store.Store {
	return &MemoryStore{
		sessions: make(map[string]store.Session),
	}
}

func (m *MemoryStore) Get(id string) (store.Session, error) {
	session, ok := m.sessions[id]
	if !ok {
		return store.Session{}, fmt.Errorf("session not found.. create an new session")
	}

	return session, nil
}

func (m *MemoryStore) Set(id string, session store.Session) error {
	m.sessions[id] = session
	return nil
}
