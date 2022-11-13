package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const keyRequestId = "requestId"

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	r.Use(func(c *gin.Context) {
		// log latency, response code, log latency
		s := time.Now()
		c.Next()
		logger.Info("Incoming request", zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)))

	}, func(c *gin.Context) {
		c.Set(keyRequestId, rand.Int())
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "ping",
		}
		if rid, exists := c.Get("requestId"); exists {
			h["requestId"] = rid
		}
		c.JSON(200, h)
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})

	r.Run()
}
