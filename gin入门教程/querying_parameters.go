package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", welCome)
	router.Run(":9999")
}

func welCome(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "ice")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	c.String(http.StatusOK, "Hi %s %s", firstname, lastname)
}