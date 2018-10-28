package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		// Upload the file to specific dst.
		// c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"uploaded": file.Filename,
		})
	})
	router.Run(":9999")
}
