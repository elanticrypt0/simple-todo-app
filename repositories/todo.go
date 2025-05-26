package repositories

import (
	"context"
	"database/sql"
	"time"
	"todo-app/models"
)

type TodoRepository struct {
	queries *models.Queries
}

func NewTodoRepository(dbPath string) (*TodoRepository, error) {
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	if err = database.Ping(); err != nil {
		return nil, err
	}

	queries := models.New(database)
	return &TodoRepository{queries}, nil
}

func (s *TodoRepository) GetAll() []models.Todo {
	dbTodos, err := s.queries.GetAllTodos(context.Background())
	if err != nil {
		return []models.Todo{}
	}

	todos := make([]models.Todo, len(dbTodos))
	for i, dbTodo := range dbTodos {
		todos[i] = models.Todo{
			ID:        dbTodo.ID,
			Title:     dbTodo.Title,
			Completed: dbTodo.Completed,
			CreatedAt: dbTodo.CreatedAt,
		}
	}
	return todos
}

func (d TodoRepository) GetByID(id string) (models.Todo, bool) {
	dbTodo, err := d.queries.GetTodoByID(context.Background(), id)
	if err != nil {
		return models.Todo{}, false
	}
	return models.Todo{
		ID:        dbTodo.ID,
		Title:     dbTodo.Title,
		Completed: dbTodo.Completed,
		CreatedAt: dbTodo.CreatedAt,
	}, true
}

func (s *TodoRepository) Create(todo models.Todo) {
	s.queries.CreateTodo(context.Background(), models.CreateTodoParams{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
}

func (d TodoRepository) Update(todo models.Todo) {
	d.queries.UpdateTodo(context.Background(), models.UpdateTodoParams{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
	})
}

func (d TodoRepository) Delete(id string) bool {
	err := d.queries.DeleteTodo(context.Background(), id)
	return err == nil
}
