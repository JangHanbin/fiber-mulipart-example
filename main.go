package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// Fiber instance
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		// Parse the multipart form:
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}
		// => *multipart.Form
		for k, v := range c.GetReqHeaders() {
			fmt.Println(k, " : ", v)
		}
		println(string(c.Body()))
		// Get all files from "documents" key:
		files := form.File["documents"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			// Save the files to disk:
			err := c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))

			// Check for errors
			if err != nil {
				return err
			}
		}
		return nil
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
