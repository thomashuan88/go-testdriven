package main

import (
	"github.com/gin-gonic/gin"
)

// go get -u github.com/gin-gonic/gin
// go get -u go.uber.org/zap
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
