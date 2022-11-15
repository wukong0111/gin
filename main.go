package main

import (
	"github.com/gin-gonic/gin"
)
var Router * gin.Engine
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})
	r.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})
	r.Run()
}
