package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
     router := gin.Default()
     router.GET("/get", getting)
/*
     router.POST("/Post", posting)
     router.PUT("/Put", putting)
     router.DELETE("/Delete", deleting)
     router.PATCH("/Patch", patching)
     router.HEAD("/Head", head)
     router.OPTIONS("/Options", options)
*/
     router.Run(":9999")
}

func getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "getting",
	})
}
