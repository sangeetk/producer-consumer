package main

import (
	"github.com/gin-gonic/gin"
)

const (
	MinThreshold = 10
	MaxThreshold = 100
)

var ItemCount int

func main() {

	r := gin.Default()

	r.GET("/consume", ConsumerHandler)
	r.GET("/count", CountHandler)

	r.Run()
}
