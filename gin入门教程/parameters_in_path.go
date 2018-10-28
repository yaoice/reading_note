package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", getUserName)
	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", getUserNameAction)
	router.Run(":9999")
}

func getUserName(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hi %s", name)
}

func getUserNameAction(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}



