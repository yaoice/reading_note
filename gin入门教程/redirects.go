package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	// listen and serve on 0.0.0.0:8080
	r.Run(":9999")
}
