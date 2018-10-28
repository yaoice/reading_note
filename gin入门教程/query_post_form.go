package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.POST("/post", Post)

	router.Run(":9999")
}

func Post(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
		"page": page,
		"name": name,
		"message": message,
	})
}
