package storage

import "todo-app/repositories"

type StorageTodo interface {
	GetAll() []repositories.Todo
	GetByID(id string) (repositories.Todo, bool)
	Create(todo repositories.Todo)
	Update(todo repositories.Todo)
	Delete(id string) bool
}
