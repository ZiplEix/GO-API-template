package todo

import (
	"context"
	"strconv"

	"gorm.io/gorm"
)

// Usage:
// Ce package fournit une interface pour interagir avec une base de données des tâches à faire ("todo").
// Il définit les structures de données, les fonctions de création, et la récupération de toutes les tâches à faire.
//
// Utilisez les méthodes `CreateTodo` et `GetAllTodos` pour créer de nouvelles tâches à faire et récupérer toutes les tâches existantes, respectivement.
//
// Exemple:
//    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//    if err != nil {
//      log.Fatal("failed to connect database")
//    }
//    todoStorage := NewTodoStorage(db)
//    id, err := todoStorage.CreateTodo("My title", "My description", false, context.Background())
//    todos, err := todoStorage.GetAllTodos(context.Background())
//

// TodoDB is the database model for a todo.
type TodoDB struct {
	gorm.Model
	Title       string `json:"title" gorm:"text;not null;default:null"`
	Description string `json:"description" gorm:"text;not null;default:null"`
	Completed   bool   `json:"completed" gorm:"not null;default:false"`
}

type TodoStorage struct {
	db *gorm.DB
}

// Constructor for the TodoStorage struct
func NewTodoStorage(db *gorm.DB) *TodoStorage {
	return &TodoStorage{db: db}
}

// CreateTodo creates a todo in the database.
func (s *TodoStorage) CreateTodo(title, description string, completed bool, ctx context.Context) (string, error) {
	todo := TodoDB{
		Title:       title,
		Description: description,
		Completed:   completed,
	}
	result := s.db.Create(&todo)
	if result.Error != nil {
		return "", result.Error
	}

	// convert todo.ID to string
	id := strconv.Itoa(int(todo.ID))

	return id, nil
}

// GetAllTodos gets all todos from the database.
func (s *TodoStorage) GetAllTodos(ctx context.Context) ([]TodoDB, error) {
	var todos []TodoDB
	result := s.db.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}