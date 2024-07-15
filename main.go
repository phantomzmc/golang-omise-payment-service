package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	modelResponse "github.com/phantomzmc/omise-go-payment-service/model"
	"github.com/phantomzmc/omise-go-payment-service/payment"
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

	app.Get( "/payment/:id", func(c fiber.Ctx) error {
		var p payment.Payment
		result := p.GetPaymentById(c.Params("id"))
		return c.JSON(result)
	})

	app.Get( "/payments", func(c fiber.Ctx) error {
		var p payment.Payment
		result := p.GetPaymentList()
		return c.JSON(result)
	})

	app.Post( "/payment", func(c fiber.Ctx) error {
		reqPayment := new(payment.RequestPayment)
		if err := c.Bind().JSON(reqPayment); err != nil {
			log.Println(err)
            return err
        }
		var p payment.Payment
		response := p.CreatePayment(reqPayment)
		return c.JSON(response)
	})

	app.Post( "/charge-by-source", func(c fiber.Ctx) error {
		reqChargeBySource := new(payment.RequestChargeBySource)
		if err := c.Bind().JSON(reqChargeBySource); err != nil {
			log.Println(err)
            return err
        }
		var p payment.Payment
		result := p.ChargeBySource(reqChargeBySource)
		modelResponse := modelResponse.CommonResponse{
			Status: "success",
			Message: "success",
			Data: result,
		}
		return c.JSON(modelResponse)
	})

	app.Post( "/charge-by-token", func(c fiber.Ctx) error {
		reqChargeByToken := new(payment.RequestChargeByToken)
		if err := c.Bind().JSON(reqChargeByToken); err != nil {
			log.Println(err)
            return err
        }
		var p payment.Payment
		result := p.ChargeByToken(reqChargeByToken)
		modelResponse := modelResponse.CommonResponse{
			Status: "success",
			Message: "success",
			Data: result,
		}
		return c.JSON(modelResponse)
	})

	app.Get("/check/:id", func(c fiber.Ctx) error {
		var p payment.Payment
		result := p.CheckStatusByChargeId(c.Params("id"))
		modelResponse := modelResponse.CommonResponse{
			Status: "success",
			Message: "success",
			Data: result,
		}
		return c.JSON(modelResponse)
	})

	app.Get("/sync", func(c fiber.Ctx) error {
		// TODO: business logic
		print("call to payment sync")
		return c.Redirect().To("http://localhost:3000/th/payment?status=success")
	})

    // Start the server on port 8080
    log.Fatal(app.Listen(":8080"))
}