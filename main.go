package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

const (
	MinThreshold = 10
	MaxThreshold = 20
)

var ItemCount int
var ProductionRate = 1

var Shutdown chan bool

func main() {

	Shutdown = make(chan bool)

	ticker := time.NewTicker(1 * time.Second)
	go Produce(ticker)

	r := gin.Default()

	r.GET("/consume", ConsumerHandler)
	r.GET("/count", CountHandler)

	r.Run()
}
