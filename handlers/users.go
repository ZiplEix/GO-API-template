package handlers

import (
	"fmt"
	"os"
	"time"

	"github.com/ZiplEix/API_template/database"
	"github.com/ZiplEix/API_template/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	// get email and password
	var body struct {
		Email    string
		Password string
	}

	if err := c.BodyParser(&body); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
		return err
	}
	// check if email is empty
	if body.Email == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email cannot be empty",
		})
		return nil
	}
	// check if password is empty
	if body.Password == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Password cannot be empty",
		})
		return nil
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot hash password",
		})
		return err
	}

	// create the user
	user := models.User{
		Email:    body.Email,
		Password: string(hash),
	}
	result := database.DB.Db.Create(&user)
	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create user",
		})
		return err
	}

	// generate JWT
	err = generateAndStoreJWT(user, c)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not login",
		})
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func Login(c *fiber.Ctx) error {
	// get credentials
	var body struct {
		Email    string
		Password string
	}

	if err := c.BodyParser(&body); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
		return err
	}
	// check if email is empty
	if body.Email == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email cannot be empty",
		})
		return nil
	}
	// check if password is empty
	if body.Password == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Password cannot be empty",
		})
		return nil
	}

	// find user
	var user models.User
	database.DB.Db.Where("email = ?", body.Email).First(&user)
	if user.ID == 0 {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
		return nil
	}

	// check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Incorrect password",
		})
		return err
	}

	// generate JWT
	err = generateAndStoreJWT(user, c)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not login",
		})
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
	})
}

func generateAndStoreJWT(user models.User, c *fiber.Ctx) error {
	fmt.Println("start generateAndStoreJWT")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println("err generateAndStoreJWT")
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

	fmt.Println("end generateAndStoreJWT")
	return nil
}

func Private(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	fmt.Println("oui oui")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": user,
	})
}
