package main

import (
	"log"

	_ "github.com/haviz000/MyGram-RestAPI/docs"

	"github.com/gin-gonic/gin"
	"github.com/haviz000/MyGram-RestAPI/database"
	"github.com/haviz000/MyGram-RestAPI/route"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const PORT = ":8080"

// @title					MyGram API
// @version					1.0
// @description				This is a MyGram API.
// @host 					localhost:8080
// @BasePath 				/
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	router := gin.Default()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("load env failed")
	}

	database.ConnectDB()
	db := database.GetDB()

	route.Routes(router, db)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(PORT)
}
