package user

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	storage *UserStorage
}

// Constructor for the UserController struct
func NewUserController(storage *UserStorage) *UserController {
	return &UserController{storage: storage}
}

type createUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createUserResponse struct {
	Token string `json:"token"`
}

// @Summary Create one user and login him.
// @Description creates one user and loged in him.
// @Tags users
// @Accept */*
// @Produce json
// @Param user body createUserRequest true "User to create"
// @Success 200 {object} createUserResponse
// @Router /register [post]
func (u *UserController) Register(c *fiber.Ctx) error {
	// parse hte request body
	var req createUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// create the user
	id, err := u.storage.CreateUser(
		req.Email,
		req.Password,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Fail creating user",
		})
	}

	// get the user
	user, err := u.storage.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Fail fetching user",
		})
	}

	// generate the jwt
	err = generateAndStoreJWT(user, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Fail generating jwt",
		})
	}

	// return the ok message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
	})
}

type loginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginUserResponse struct {
	Token string `json:"token"`
}

// @Summary Login one user.
// @Description login one user.
// @Tags users
// @Accept */*
// @Produce json
// @Param user body loginUserRequest true "User to login"
// @Success 200 {object} loginUserResponse
// @Router /login [post]
func (u *UserController) Login(c *fiber.Ctx) error {
	// parse the request body
	var req loginUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// get the user with the given email
	user, err := u.storage.GetUserByEmail(req.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Fail fetching user",
		})
	}
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// generate the jwt
	err = generateAndStoreJWT(user, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Fail generating token",
		})
	}

	// return the ok message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
	})
}

func generateAndStoreJWT(user *UserDB, c *fiber.Ctx) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}

	// set the JWT as a cookie
	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})

	return nil
}
