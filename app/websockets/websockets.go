package websockets

import (
	"go_web_server/models"
	"time"

	"github.com/gofiber/websocket/v2"
	"gorm.io/gorm"
)

func ShowUsersWebsocket(c *websocket.Conn, db *gorm.DB) {
	var users []models.User

	for {
		db.Find(&users)

		err := c.WriteJSON(users)
		if err != nil {
			break
		}

		time.Sleep(5 * time.Second)
	}

}
