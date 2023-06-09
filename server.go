package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"database/sql"
	"github.com/lib/pq"
)

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
   return c.SendString("Hello")
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
   return c.SendString("Hello")
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
   return c.SendString("Hello")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
   return c.SendString("Hello")
}

func main() {

	// set db
	dbStr := "postgresql://casio:casio@localhost:5432/todos?sslmode=disable"
	db, err := sql.Open("postgres", dbStr)
	if err != nil {
		log.Fatal(err)
	}
	// init app
	app := fiber.New()
	// set handlers
	app.Get("/", func(c *fiber.Ctx) error {
       		return indexHandler(c, db)
   	})

   	app.Post("/", func(c *fiber.Ctx) error {
       		return postHandler(c, db)
   	})

   	app.Put("/update", func(c *fiber.Ctx) error {
       		return putHandler(c, db)
   	})

  	app.Delete("/delete", func(c *fiber.Ctx) error {
       		return deleteHandler(c, db)
  	})
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
	
}
