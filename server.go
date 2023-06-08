package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// init app
	app := fiber.New()
	// set handlers
	app.Get("/", indexHandler)
	app.Post("/", postHandler)
	app.Put("/update", putHandler)
	app.Delete("/delete", deleteHandler)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Spintf(":%v", port)))
}
