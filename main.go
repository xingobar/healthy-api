package main

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

	// 血壓紀錄
	//blood := r.Group("/blood")
	//router.BloodRouter(blood)

	// 設備
	device := r.Group("/device")
	router.DeviceRouter(device)

	r.Run()
}