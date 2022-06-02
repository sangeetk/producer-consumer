package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConsumerHandler(c *gin.Context) {

	if ItemCount <= 0 {
		ProductionRate++
		c.Status(http.StatusTooEarly)
		return
	}

	ItemCount--

	c.Status(http.StatusOK)
}
