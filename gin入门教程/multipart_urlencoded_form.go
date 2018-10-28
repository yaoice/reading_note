package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.POST("/form_post", formPost)

	router.Run(":9999")
}

func formPost(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "test")
	c.JSON(http.StatusOK, gin.H{
		"status": "posted",
		"message": message,
		"nick": nick,
	})
}



