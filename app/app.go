package app

import (
	controllers "go_web_server/app/controllers"
	handlers "go_web_server/app/handlers"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

func SetupRouters(api fiber.Router, db *gorm.DB) {
	api.Get("/user/:name", handlers.GetHelloUser)
	api.Get("/request_information", handlers.GetRequestInformation)

	api.Post("/new_user", func(c *fiber.Ctx) error {
		return controllers.AddNewUser(c, db)
	})
	api.Get("/show_users", func(c *fiber.Ctx) error {
		return controllers.ShowAllUsers(c, db)
	})
	api.Get("/show_users_table", func(c *fiber.Ctx) error {
		return controllers.ShowAllUsers2(c, db)
	})
	api.Patch("/update_user_by_id", func(c *fiber.Ctx) error {
		return controllers.UpdateUserInformation(c, db)
	})
	api.Delete("/delete_user_by_id", func(c *fiber.Ctx) error {
		return controllers.DeleteUser(c, db)
	})
}
