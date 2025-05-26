package repositories

import "todo-app/models"

type StorageTodo interface {
	GetAll() []models.Todo
	GetByID(id string) (models.Todo, bool)
	Create(todo models.Todo)
	Update(todo models.Todo)
	Delete(id string) bool
}
