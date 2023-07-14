package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vk-notification-monitor/entity"
	"vk-notification-monitor/store"
)

func GetWallByGroupAddress(repo *store.Repository) gin.HandlerFunc {

	return func(c *gin.Context) {
		wall, _ := repo.Wall.GetWallPostsByDomain("overhearelgorsk")
		keywords := []entity.Keyword{"каталог"}
		repo.Post.GetPostsByKeywords(wall, keywords)
		c.JSON(http.StatusOK, wall)
	}
}
