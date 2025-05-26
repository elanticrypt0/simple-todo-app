package storage

import (
	"context"
	"database/sql"
	"time"
	"todo-app/db"
)

type DatabaseStorage struct {
	queries *db.Queries
}

func NewDatabaseStorage(dbPath string) (*DatabaseStorage, error) {
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	if err = database.Ping(); err != nil {
		return nil, err
	}

	queries := db.New(database)
	return &DatabaseStorage{queries}, nil
}

func (s *DatabaseStorage) GetAll() []db.Todo {
	dbTodos, err := s.queries.GetAllTodos(context.Background())
	if err != nil {
		return []db.Todo{}
	}

	todos := make([]db.Todo, len(dbTodos))
	for i, dbTodo := range dbTodos {
		todos[i] = db.Todo{
			ID:        dbTodo.ID,
			Title:     dbTodo.Title,
			Completed: dbTodo.Completed,
			CreatedAt: dbTodo.CreatedAt,
		}
	}
	return todos
}

func (d DatabaseStorage) GetByID(id string) (db.Todo, bool) {
	dbTodo, err := d.queries.GetTodoByID(context.Background(), id)
	if err != nil {
		return db.Todo{}, false
	}
	return db.Todo{
		ID:        dbTodo.ID,
		Title:     dbTodo.Title,
		Completed: dbTodo.Completed,
		CreatedAt: dbTodo.CreatedAt,
	}, true
}

func (s *DatabaseStorage) Create(todo db.Todo) {
	s.queries.CreateTodo(context.Background(), db.CreateTodoParams{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
}

func (d DatabaseStorage) Update(todo db.Todo) {
	d.queries.UpdateTodo(context.Background(), db.UpdateTodoParams{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
	})
}

func (d DatabaseStorage) Delete(id string) bool {
	err := d.queries.DeleteTodo(context.Background(), id)
	return err == nil
}
