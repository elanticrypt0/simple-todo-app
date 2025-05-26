package main

import (
	"database/sql"
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"todo-app/handlers"
	"todo-app/repositories"
)

//go:embed database/migrations/*.sql
var migrationFS embed.FS

func main() {

	dbpath := "mytodos.db"

	err := runMigrations(dbpath)
	if err != nil {
		log.Fatal(err)
	}

	store, err := repositories.NewTodoRepository(dbpath)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	todoHandler := handlers.NewTodoHandler(store)

	// routes
	r.GET("/todos/:id", todoHandler.GetTodoByID)
	r.GET("/", todoHandler.IndexPage)
	r.GET("/about", todoHandler.AboutPage)
	r.POST("/todos", todoHandler.CreateTodo)
	r.PATCH("/todos", todoHandler.ToggleTodo)
	r.DELETE("/todos", todoHandler.DeleteTodo)

	log.Println("Servidor iniciado en http://localhost:8080")

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func runMigrations(dbPath string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return err
	}

	d, err := iofs.New(migrationFS, "database/migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", d, "sqlite3", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
