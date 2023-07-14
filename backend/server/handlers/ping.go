package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddPingRoutes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	}
}
