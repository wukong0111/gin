package main

import (
	"github.com/gin-gonic/gin"
)
var Router * gin.Engine
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "<h1>OK</h1>")
	})
	r.POST("/", func(c *gin.Context) {
		c.String(200, "<h1>OK</h1>")
	})
	r.Run()
}
