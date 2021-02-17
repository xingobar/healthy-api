package healthy_api

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"healthy-api/router"
	"log"
)

func main() {
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	blood := r.Group("/blood")
	router.BloodRouter(blood)

	r.Run(":9999")
}