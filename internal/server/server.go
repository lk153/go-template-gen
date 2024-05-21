package server

import "github.com/gin-gonic/gin"

func init() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})
	// listen and serve on 0.0.0.0:8080
	// on windows "localhost:8080"
	// can be overriden with the PORT env var
	r.Run()
}
