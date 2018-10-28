package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// simulate some private data
var secrets = gin.H{
	"yao":    gin.H{"email": "yao@bar.com", "phone": "123433"},
	"ice":    gin.H{"email": "ice@example.com", "phone": "666"},
}

func main() {
	r := gin.Default()

	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"yao":  "yao",
		"ice":  "ice",
	}))

	// /admin/secrets endpoint
	// hit "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// Listen and serve on 0.0.0.0:9999
	r.Run(":9999")
}