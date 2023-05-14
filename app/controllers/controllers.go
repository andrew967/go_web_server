package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"go_web_server/models"
	"net/http"

	"gorm.io/gorm"
)

func AddNewUser(c *fiber.Ctx, db *gorm.DB) error {
	var user models.User
	age, err := strconv.Atoi(c.FormValue("age"))

	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Age is not string")
	}

	user.FillData(c.FormValue("name"), c.FormValue("surname"), c.FormValue("phonenumber"), c.FormValue("email"), c.FormValue("sex"), age)

	err = user.AddToDatabase(db)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Something goes wrong")
	}

	return c.Status(http.StatusOK).SendString("User has been successfully added.")
}

func ShowAllUsers(c *fiber.Ctx, db *gorm.DB) error {
	users := []models.User{}
	result := db.Find(&users)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).SendString("Something goes wrong.")
	}

	jsonString, err := json.Marshal(users)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Something goes wrong.")
	}

	return c.Status(http.StatusAccepted).SendString(string(jsonString))
}
