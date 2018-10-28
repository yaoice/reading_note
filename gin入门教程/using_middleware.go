package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Default With the Logger and Recovery middleware already attached
	//r := gin.Default()

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Per route middleware, you can add as many as you desire.
	r.GET("/benchmark", MyBenchLogger, benchEndpoint)

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/v1")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(AuthRequired)
	{
		authorized.POST("/login", loginEndpoint)
//		authorized.POST("/submit", submitEndpoint)
//		authorized.POST("/read", readEndpoint)

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":9999")
}

func MyBenchLogger(c *gin.Context) {

}

func benchEndpoint(c *gin.Context) {

}

func AuthRequired(c *gin.Context) {

}

func loginEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "login")
}

func analyticsEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "analytics")
}
