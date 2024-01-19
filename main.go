package main

import (
	"log"
	"os"
	// "strings"
	// "regexp"

	kompas "github.com/dhupee/Indonesia-News-Aggregator/kompas"
	detik "github.com/dhupee/Indonesia-News-Aggregator/detik"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env")
		log.Println("Using default environment variables")
	}

	// Get the port from the environment variables
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	// Create a new Fiber app
	app := fiber.New(fiber.Config{
		ServerHeader: "Indonesia-News-Aggregator",
	})

	v1 := app.Group("/v1") // add v1 grouping to manage if needed

	// Define the route handler for the root path and non-source-specific routes
	v1.Get("/", RootHandler)
	// v1.Get("/search", SearchHandler)
	v1.Get("/ping", func(c *fiber.Ctx) error { // health check
		return c.SendString("pong")
	})

	// Define the route handlers for Kompas endpoints
	v1.Get("/kompas/index", kompas.KompasIndexHandler)
	v1.Get("/kompas/news", kompas.KompasNewsHandler)

	// Define the route handlers for Detik endpoints
	v1.Get("/detik/news", detik.DetikNewsHandler)

	// Start the app on the specified port
	log.Fatal(app.Listen(":"+port))
}

func RootHandler(c *fiber.Ctx) error {
	return c.SendFile("./assets/welcome.txt")
}

// TODO: this one next, maybe just use universal solution
// TODO: W.I.P
// func SearchHandler(c *fiber.Ctx) error {
// 	keyword := c.Get("keyword")
// 	if keyword == "" {
// 		return c.SendString("Please specify keyword")
// 	}

// 	result := kompas.Search(keyword)
// 	return c.SendString("You search for " + result)
// }
