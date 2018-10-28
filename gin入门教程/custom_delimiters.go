package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Delims("{[{", "}]}")
	r.LoadHTMLGlob("/path/to/templates")
}
