package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func main() {
	route := gin.Default()
	route.Any("/testing", formHandler)
	route.Run(":9999")
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
}