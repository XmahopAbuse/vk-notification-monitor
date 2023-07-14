package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"strconv"
	"vk-notification-monitor/entity"
	"vk-notification-monitor/store"
)

func SyncPostsByKeywords(repo *store.Repository) gin.HandlerFunc {

	return func(c *gin.Context) {
		posts, err := repo.Post.SyncPostsByKeywords(&repo.Group, &repo.Keyword, &repo.Wall, &repo.Post, &repo.Notification, nil, 0)
		if err != nil {
			c.JSON(500, "Error")
			return
		}
		c.JSON(200, posts)
	}

}

func GetPosts(repo *store.Repository) gin.HandlerFunc {

	return func(c *gin.Context) {
		posts, err := repo.Post.GetAll()
		if err != nil {
			c.JSON(500, "Ошибка получения постов")
		}
		c.JSON(200, posts)
	}
}

func GetAuthor(repo *store.Repository) gin.HandlerFunc {
	r := struct {
		AuthorId string `json:"author_id"`
	}{}

	return func(c *gin.Context) {
		err := c.BindJSON(&r)
		if err != nil || r.AuthorId == "" {
			c.JSON(500, "Invalid request")
			return
		}

		id, err := strconv.Atoi(r.AuthorId)
		if err != nil {
			log.Println(err)
			c.JSON(500, "Invalid request")
			return
		}

		a := &entity.Author{}

		if id > 0 {
			a, err = repo.Wall.GetPostAuthor(id)
			if err != nil {
				log.Println(err)

				c.JSON(500, "Invalid request")
				return
			}
		} else {
			a, err = repo.Group.GetGroupAsAuthor(id)
			if err != nil {
				log.Println("локально группы не найдено")
				g, err := repo.Wall.GetGroupById(int(math.Abs(float64(id))))
				if err != nil {
					log.Println(err)
					c.JSON(500, "Error")
					return
				}

				newa := entity.Author{}

				newa.Id = g.Id
				newa.Name = g.Name
				newa.Photo = g.PhotoUrl
				newa.FullUrl = g.FullUrl

				a = &newa
			}

		}

		c.JSON(200, a)
	}
}

func GetPost(repo *store.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParams := c.Request.URL.Query()
		params := make(map[string]string)

		for key, values := range queryParams {
			if len(values) > 0 {
				params[key] = values[0]
			}
		}
		posts, err := repo.Post.GetPost(params)
		if err != nil {
			log.Println(err)
			c.JSON(500, "error")
			return
		}

		c.JSON(200, posts)
	}
}

func UpdatePost(repo *store.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		hash := c.Param("hash")
		if len(hash) <= 0 {
			c.JSON(400, "Invalid request")
			return
		}

		var updateData map[string]interface{}
		err := c.ShouldBind(&updateData)
		if err != nil {
			c.JSON(400, "Invalid data")
			return
		}

		err = repo.Post.Update(updateData, hash)
		if err != nil {
			log.Println(err)
			c.JSON(500, err)
			return
		}
		c.JSON(200, "OK")
	}
}
