package main

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/generate-uuid", func(c *gin.Context) {
		id := uuid.New()

		c.JSON(http.StatusOK, gin.H{
			"uuid": id.String(),
		})
	})

	router.Run(":8080")
}
