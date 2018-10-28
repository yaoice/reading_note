package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", login1)
//		v1.POST("/submit", submit)
//		v1.POST("/read", read)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", login2)
//		v2.POST("/submit", submit)
//		v2.POST("/read", read)
	}
	router.Run(":9999")
}

func login1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"login1": "haha",
	})
}

func login2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"login2": "haha",
	})
}


