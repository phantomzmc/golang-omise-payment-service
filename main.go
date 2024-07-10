package main

import (
	"log"

	"github.com/biskitsx/go-fiber-api/payment"
	"github.com/gofiber/fiber/v3"
)



func main() {
    // Initialize a new Fiber app
    app := fiber.New()

    // Define a route for the GET method on the root path '/'
    app.Get("/", func(c fiber.Ctx) error {
        // Send a string response to the client
        return c.SendString("Hello, World 👋!")
    })

	app.Get("/payment", func(c fiber.Ctx) error {
		var p payment.Payment
		response := p.GetPayment()
		return c.JSON(response)
	})

	// app.Get( "/payment/:id", func(c fiber.Ctx) error {
	// 	var p payment.Payment
	// 	result := p.GetPaymentById(c.Params("id"))
	// 	return c.JSON(result)
	// })

	app.Get( "/payments", func(c fiber.Ctx) error {
		var p payment.Payment
		result := p.GetPaymentById()
		return c.JSON(result)
	})

	app.Post( "/payment", func(c fiber.Ctx) error {
		var p payment.Payment
		response := p.CreatePayment()
		return c.JSON(response)
	})



    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}