package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/superdie2014/blogbackend/database"
	"github.com/superdie2014/blogbackend/routes"
)

func main()  {
	
	database.Connect()
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file!")
	}

	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":"+port)
}