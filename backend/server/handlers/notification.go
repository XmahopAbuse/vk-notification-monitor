package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"vk-notification-monitor/entity"
	"vk-notification-monitor/store"
)

func AddNotification(repository *store.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		n := entity.Notification{}

		err := c.BindJSON(&n)
		if err != nil {
			c.JSON(400, "Error")
			log.Println(err)
			return
		}

		err = repository.Notification.AddNotification(n.Type, n.Value)
		if err != nil {
			c.JSON(500, "Error")
			log.Println(err)

			return
		}
		c.JSON(200, "Ok")

	}
}

func GetNotificationByType(repo *store.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := c.Query("type")
		fmt.Println(t)
		c.JSON(200, "Ok")
	}
}
