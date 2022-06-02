package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CountHandler(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}
