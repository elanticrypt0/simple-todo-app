package storage

import (
	"context"
	"database/sql"
	"time"
	"todo-app/repositories"
)

type DatabaseStorage struct {
	queries *repositories.Queries
}

func NewDatabaseStorage(dbPath string) (*DatabaseStorage, error) {
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	if err = database.Ping(); err != nil {
		return nil, err
	}

	queries := repositories.New(database)
	return &DatabaseStorage{queries}, nil
}

func (s *DatabaseStorage) GetAll() []repositories.Todo {
	dbTodos, err := s.queries.GetAllTodos(context.Background())
	if err != nil {
		return []repositories.Todo{}
	}

	todos := make([]repositories.Todo, len(dbTodos))
	for i, dbTodo := range dbTodos {
		todos[i] = repositories.Todo{
			ID:        dbTodo.ID,
			Title:     dbTodo.Title,
			Completed: dbTodo.Completed,
			CreatedAt: dbTodo.CreatedAt,
		}
	}
	return todos
}

func (d DatabaseStorage) GetByID(id string) (repositories.Todo, bool) {
	dbTodo, err := d.queries.GetTodoByID(context.Background(), id)
	if err != nil {
		return repositories.Todo{}, false
	}
	return repositories.Todo{
		ID:        dbTodo.ID,
		Title:     dbTodo.Title,
		Completed: dbTodo.Completed,
		CreatedAt: dbTodo.CreatedAt,
	}, true
}

func (s *DatabaseStorage) Create(todo repositories.Todo) {
	s.queries.CreateTodo(context.Background(), repositories.CreateTodoParams{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
}

func (d DatabaseStorage) Update(todo repositories.Todo) {
	d.queries.UpdateTodo(context.Background(), repositories.UpdateTodoParams{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
	})
}

func (d DatabaseStorage) Delete(id string) bool {
	err := d.queries.DeleteTodo(context.Background(), id)
	return err == nil
}
