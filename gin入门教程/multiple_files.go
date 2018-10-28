package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["file[]"]

		for _, file := range files {
			log.Println(file.Filename)
			// Upload the file to specific dst.
			// c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"files uploaded": len(files),
		})
	})
	router.Run(":9999")
}
