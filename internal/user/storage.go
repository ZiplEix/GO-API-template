package user

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserDB struct {
	gorm.Model
	Email    string `json:"email" gorm:"text;not null;default:null;unique"`
	Password string `json:"password" gorm:"text;not null;default:null"`
}

type UserStorage struct {
	db *gorm.DB
}

// Constructor for the UserStorage struct
func NewUserStorage(db *gorm.DB) *UserStorage {
	return &UserStorage{db: db}
}

// CreateUser creates a user in the database.
func (s *UserStorage) CreateUser(email, password string) (string, error) {
	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := UserDB{
		Email:    email,
		Password: string(hash),
	}
	result := s.db.Create(&user)
	if result.Error != nil {
		return "", result.Error
	}

	// convert user.ID to string
	id := strconv.Itoa(int(user.ID))

	return id, nil
}

// GetUserByEmail retrieves a user from the database by email.
func (s *UserStorage) GetUserByEmail(email string) (*UserDB, error) {
	var user UserDB
	result := s.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if user.ID == 0 {
		return nil, nil
	}

	return &user, nil
}

// GetUserByID retrieves a user from the database by ID.
func (s *UserStorage) GetUserByID(id string) (*UserDB, error) {
	var user UserDB
	result := s.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
