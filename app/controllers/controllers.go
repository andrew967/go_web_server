package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"go_web_server/app/auth"
	"go_web_server/models"
	"html/template"
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

func ShowAllUsers2(c *fiber.Ctx, db *gorm.DB) error {
	var users []models.User
	response := db.Find(&users)
	if response.Error != nil {
		return c.Status(500).SendString("Something goes wrong")
	}
	tmp1 := template.Must(template.ParseFiles("templates/users.html"))
	c.Response().Header.Set("Content-Type", "text/html")
	return tmp1.Execute(c.Response().BodyWriter(), users)
}

func UpdateUserInformation(c *fiber.Ctx, db *gorm.DB) error {
	var user models.User
	response := db.Find(&user, "id = ?", c.FormValue("id"))

	if response.Error != nil {
		return c.Status(500).SendString("Something goes wrong")
	}

	user.Name = c.FormValue("new_name")
	user.Surname = c.FormValue("new_surname")
	user.Sex = c.FormValue("new_sex")
	user.Email = c.FormValue("new_email")
	user.PhoneNumber = c.FormValue("new_phonenumber")
	age, err := strconv.Atoi(c.FormValue("new_age"))
	if err != nil {
		return c.Status(400).SendString("Age is not int")
	}
	user.Age = age

	db.Save(&user)
	return c.Status(http.StatusOK).SendString("Information has been updated successfully")
}

func DeleteUser(c *fiber.Ctx, db *gorm.DB) error {
	var user models.User
	response := db.Find(&user, "id = ?", c.FormValue("id"))

	if response.Error != nil {
		return c.Status(500).SendString("We can not find the user with id " + c.FormValue("id"))
	}

	response = db.Delete(&user)
	if response.Error != nil {
		return c.Status(500).SendString("We can not delete the user with id " + c.FormValue("id"))
	}

	return c.Status(http.StatusOK).SendString("User has been deleted successfully")
}

func ShowAllUsers3(c *fiber.Ctx, db *gorm.DB) error {
	var users []models.User
	response := db.Find(&users)

	if response.Error != nil {
		return c.Status(http.StatusBadRequest).SendString("Something goes wrong")
	}
	tmp1 := template.Must(template.ParseFiles("templates/userwb.html"))
	c.Response().Header.Set("Content-Type", "text/html")
	return tmp1.Execute(c.Response().BodyWriter(), users)
}

func Authorization(c *fiber.Ctx, db *gorm.DB) error {
	login := c.FormValue("login")
	password := c.FormValue("password")
	var user models.Authorization
	response := db.Where("login = ?", login).Find(&user)
	if response.Error != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error!")
	}

	if user.Password != password {
		return c.Status(fiber.StatusBadRequest).SendString("Login or password is incorrect.")
	}

	token, err := auth.GenerateToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate token")
	}
	return c.JSON(fiber.Map{"token": token})
}

func SignUp(c *fiber.Ctx) error {
	return c.SendFile("templates/signup.html")
}
