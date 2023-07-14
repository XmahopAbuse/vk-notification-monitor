package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"log"
	"strings"
	"vk-notification-monitor/store"
)

func AddKeyword(repo *store.Repository) gin.HandlerFunc {

	r := struct {
		Keyword string `json:"keyword"`
	}{}

	return func(c *gin.Context) {
		err := c.BindJSON(&r)
		if err != nil {
			c.JSON(400, "Error")
			return
		}

		err = repo.Keyword.Add(strings.ToLower(r.Keyword))

		if err != nil {
			pqErr, ok := err.(*pq.Error)
			if ok && pqErr.Code == "23505" {
				c.JSON(400, "Ключевое слово уже существует")
				return
			}
			c.JSON(400, "Error")
			return
		}

		c.JSON(200, "OK")
	}

}

func GetAllKeywords(repo *store.Repository) gin.HandlerFunc {

	return func(c *gin.Context) {
		keywords, err := repo.Keyword.GetAll()

		if err != nil {
			c.JSON(500, "Error")
			return
		}

		c.JSON(200, keywords)

	}
}

func DeleteKeyword(repo *store.Repository) gin.HandlerFunc {

	return func(c *gin.Context) {
		name := c.Param("name")
		err := repo.Keyword.Delete(name)
		if err != nil {
			c.JSON(500, "Error")
			log.Println(err)
			return
		}
		c.JSON(200, "Ok")
	}
}
