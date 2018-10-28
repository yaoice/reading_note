package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	router := gin.Default()
	html := template.Must(template.ParseFiles("file1", "file2"))
	router.SetHTMLTemplate(html)
	router.Run(":9999")
}
