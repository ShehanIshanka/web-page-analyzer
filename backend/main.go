package main

import (
	"log"
	"net/http"

    "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"backend/analyzer"
)

func main() {
	r := gin.Default()

	// For CORS
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/analyze", func(c *gin.Context) {
		var req struct {
			URL string `json:"url"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		data, err := analyzer.AnalyzeURL(req.URL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	})

	
	log.Fatal(r.Run(":8080"))
}
