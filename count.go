package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CountHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"count": ItemCount,
	})
}
