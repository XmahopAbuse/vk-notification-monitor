package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"vk-notification-monitor/store"
)

func AddGroup(repo *store.Repository) gin.HandlerFunc {
	r := struct {
		Address string `json:"address"`
	}{}
	return func(c *gin.Context) {
		err := c.BindJSON(&r)
		if err != nil {
			c.JSON(400, "error")
			return
		}

		if strings.Contains(r.Address, "vk.com/") == false {
			c.JSON(400, "Неверный формат")
			return
		}

		groupAddr := strings.Split(r.Address, "vk.com/")[1]

		group, err := repo.Group.Add(groupAddr)

		if err != nil {
			c.JSON(500, "Ошибка при добавлении")
			log.Println(err)
			return
		}

		c.JSON(200, group)
	}
}

func GetAllGroups(repo *store.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		groups, err := repo.Group.GetAllGroups()
		if err != nil {
			c.JSON(500, "Error")
			return
		}

		c.JSON(200, &groups)
		return
	}
}

func DeleteGroupByAddress(repo *store.Repository) gin.HandlerFunc {
	r := struct {
		Address string `json:"address"`
	}{}
	return func(c *gin.Context) {
		err := c.BindJSON(&r)
		if err != nil {
			log.Println(err)
			c.JSON(500, "Error")
			return
		}

		err = repo.Group.DeleteByAddress(r.Address)

		if err != nil {
			log.Println(err)
			c.JSON(500, "Error")
			return
		}

		c.JSON(200, "Ok")

	}
}

func GetGroupById(repo *store.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		group, err := repo.Group.GetById(id)
		if err != nil {
			log.Println(err)
			c.JSON(404, "Not found")
			return
		}

		c.JSON(200, group)
	}
}
