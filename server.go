package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"database/sql"
)

func indexHandler(c *fiber.Ctx) error {
	return c.SendString("INDEX")
}

func postHandler(c *fiber.Ctx) error {
	return c.SendString("POST")
}

func putHandler(c *fiber.Ctx) error {
	return c.SendString("PUT")
}

func deleteHandler(c *fiber.Ctx) error {
	return c.SendString("DELETE")
}

func main() {

	// set db
	dbStr := "postgresql://<username>:<password>@<database_ip>/todos?sslmode=disable"
	db, err := sql.Open("postgres", dbStr)
	if err != nil {
		log.Fatal(err)
	}
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

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
	
}
