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
	OwnerID     uint   `json:"owner_id" gorm:"not null;default:null"`
}

type TodoStorage struct {
	db *gorm.DB
}

// Constructor for the TodoStorage struct
func NewTodoStorage(db *gorm.DB) *TodoStorage {
	return &TodoStorage{db: db}
}

// CreateTodo creates a todo in the database.
func (s *TodoStorage) CreateTodo(title, description string, completed bool, ownerID uint, ctx context.Context) (string, error) {
	todo := TodoDB{
		Title:       title,
		Description: description,
		Completed:   completed,
		OwnerID:     ownerID,
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
func (s *TodoStorage) GetAllTodos(ownerID uint, ctx context.Context) ([]TodoDB, error) {
	var todos []TodoDB
	result := s.db.Find(&todos, "owner_id = ?", ownerID)
	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}

// GetTodoByID gets a todo from the database by its ID.
func (s *TodoStorage) GetTodoByID(ownerID uint, id string, ctx context.Context) (*TodoDB, error) {
	var todo TodoDB
	// result := s.db.First(&todo, id)
	result := s.db.Where("owner_id = ? AND id = ?", ownerID, id).First(&todo)
	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

// UpdateTodo updates a todo in the database.
func (s *TodoStorage) UpdateTodo(
	updatedTodo TodoDB,
	ctx context.Context,
) (*TodoDB, error) {
	result := s.db.Save(&updatedTodo)
	if result.Error != nil {
		return nil, result.Error
	}

	return &updatedTodo, nil
}

// DeleteTodo deletes a todo from the database.
func (s *TodoStorage) DeleteTodo(
	toDelete TodoDB,
	ctx context.Context,
) error {
	result := s.db.Delete(&toDelete)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
