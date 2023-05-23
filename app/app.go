package app

import (
	"go_web_server/app/auth"
	controllers "go_web_server/app/controllers"
	handlers "go_web_server/app/handlers"
	websockets "go_web_server/app/websockets"
	"go_web_server/models"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var connections models.Connections

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

	api.Get("/show_all_users", func(c *fiber.Ctx) error {
		return controllers.ShowAllUsers3(c, db)
	})

	//websocket protocol
	api.Get("/ws", websocket.New(func(c *websocket.Conn) {
		websockets.ShowUsersWebsocket(c, db)
	}))

	api.Post("/login", func(c *fiber.Ctx) error {
		return controllers.Authorization(c, db)
	})

	api.Get("/protected", auth.Protected)

	api.Get("/signup", controllers.SignUp)
}
