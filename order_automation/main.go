package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gowaves/order_automaiton/configs"
	"github.com/gowaves/order_automaiton/routes"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app)

	app.Listen(":6000")
}
