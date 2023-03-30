package httpGin

import "github.com/gin-gonic/gin"

func HelloWorld() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
