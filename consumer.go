package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConsumerHandler(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}
