package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.Query("user")
		pwd := c.Query("pwd")
		c.JSON(200, gin.H{
			"success": true,
			"id":      id,
			"user":    user,
			"pwd":     pwd,
		})
	})

	r.POST("/path/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	r.DELETE("/path/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	// r.GET("/path/:id", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"success": true,
	// 	})
	// })

	r.Run(":8899")
}
