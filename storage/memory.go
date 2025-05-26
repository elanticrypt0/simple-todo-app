package storage

import (
	"sync"
	"todo-app/db"
	"todo-app/models"
)

type Storage interface {
	GetAll() []db.Todo
	GetByID(id string) (db.Todo, bool)
	Create(todo db.Todo)
	Update(todo db.Todo)
	Delete(id string) bool
}

type MemoryStorage struct {
	todos map[string]models.Todo
	mutex sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		todos: make(map[string]models.Todo),
		mutex: sync.RWMutex{},
	}
}

func (m *MemoryStorage) GetAll() []models.Todo {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	todos := make([]models.Todo, 0, len(m.todos))
	for _, t := range m.todos {
		todos = append(todos, t)
	}

	return todos
}

func (m *MemoryStorage) GetByID(id string) (models.Todo, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	todo, exists := m.todos[id]
	return todo, exists
}

func (m *MemoryStorage) Create(todo models.Todo) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.todos[todo.ID] = todo
}

func (m *MemoryStorage) Update(todo models.Todo) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.todos[todo.ID] = todo
}

func (m *MemoryStorage) Delete(id string) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.todos[id]; exists {
		delete(m.todos, id)
		return true
	}
	return false
}
