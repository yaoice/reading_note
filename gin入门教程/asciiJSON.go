package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/someJSON", func(c *gin.Context){
		data := map[string]interface{}{
			"lan": "go语言",
			"tag": "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})
	r.Run(":9999")
}
