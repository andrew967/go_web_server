package auth

import (
	"go_web_server/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var jwtKey = []byte("my_secret_key")

func GenerateToken(user models.Authorization) (string, error) {
	claims := jwt.MapClaims{
		"login": user.Login,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func Protected(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	if len(tokenString) < 7 {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing token")
	}

	tokenString = tokenString[7:]

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil { // Проверяем наличие ошибок валидации.
		if err == jwt.ErrSignatureInvalid {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid token signature")
		}

		return c.Status(fiber.StatusBadRequest).SendString("Invalid token")
	}

	if !token.Valid { // Проверяем, является ли токен действительным.
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
	}

	return c.SendString("You are authorized")
}
