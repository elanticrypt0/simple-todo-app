package handlers

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"todo-app/repositories"
	"todo-app/storage"
	"todo-app/templates"
)

type TodoHandler struct {
	storage storage.StorageTodo
}

func NewTodoHandler(store storage.StorageTodo) *TodoHandler {
	return &TodoHandler{
		storage: store,
	}
}

func (h *TodoHandler) IndexPage(c *gin.Context) {
	todos := h.storage.GetAll()

	component := templates.TodosPage(todos)
	err := component.Render(c.Request.Context(), c.Writer)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "templates/errors/500.html", nil)
	}
}

func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		log.Println("ID is empty")
	} else {
		log.Println("ID is ", id)
	}

	todo, err := h.storage.GetByID(id)
	if err {
		c.HTML(http.StatusInternalServerError, "templates/errors/500.html", gin.H{
			"message": err,
		})
	}

	c.JSONP(http.StatusOK, todo)
}

func (h *TodoHandler) AboutPage(c *gin.Context) {
	component := templates.AboutPage()
	err := component.Render(c.Request.Context(), c.Writer)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "templates/errors/500.html", gin.H{
			"message": err.Error(),
		})
	}
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	title := c.PostForm("title")
	if title == "" {
		c.HTML(http.StatusBadRequest, "templates/errors/errors.html", gin.H{
			"message": "Title is required",
		})
	}

	todo := repositories.Todo{
		ID:        uuid.New().String(),
		Title:     title,
		Completed: false,
	}

	h.storage.Create(todo)

	component := templates.TodoItem(todo)
	err := component.Render(c.Request.Context(), c.Writer)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "templates/errors/500.html", gin.H{
			"message": err.Error(),
		})
	}

}

func (h *TodoHandler) ToggleTodo(c *gin.Context) {
	id := c.Param("id")

	todo, exists := h.storage.GetByID(id)
	if !exists {
		c.HTML(http.StatusInternalServerError, "templates/errors/500.html", gin.H{
			"message": "Element not found",
		})
	}
	todo.Completed = !todo.Completed
	h.storage.Update(todo)

	component := templates.TodoItem(todo)
	err := component.Render(c.Request.Context(), c.Writer)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "templates/errors/500.html", gin.H{
			"message": err.Error(),
		})
	}
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	if !h.storage.Delete(id) {
		c.HTML(http.StatusInternalServerError, "templates/errors/500.html", gin.H{
			"message": "Element not found",
		})
	}

	c.Status(http.StatusOK)
}
