package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World!"})
	})

	router.POST("/hello", func(context *gin.Context) {
		name := context.Query("name")

		if name == "" {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Name is required",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{ // StatusOK, не StatusBadRequest!
			"message": "Hello, " + name,
		})
	})

	router.POST("/user", func(context *gin.Context) {
		name := context.Query("name")
		age := context.Query("age")

		context.JSON(http.StatusOK, gin.H{ // StatusOK, не StatusBadRequest!
			"message": "Hello, " + name + " тебе " + age + " ?",
		})
	})

	router.POST("/user/:id", func(context *gin.Context) {
		id := context.Param("id")
		context.JSON(http.StatusOK, gin.H{
			"user_id": id,
			"message": "hello user, is your id " + id,
		})
	})

	fmt.Println("Server starting on :8080")
	router.Run(":8080") // Добавьте эту строку!
}
