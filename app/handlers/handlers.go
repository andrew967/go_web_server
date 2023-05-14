package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetHelloUser(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("Hello " + c.Params("name"))
}

func GetRequestInformation(c *fiber.Ctx) error {
	information := struct {
		Body    []byte
		Method  string
		IP      string
		BaseUrl string
	}{
		Body:    c.Body(),
		Method:  c.Method(),
		IP:      c.IP(),
		BaseUrl: c.BaseURL(),
	}
	return c.Status(http.StatusOK).JSON(information)
}
