package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gowaves/order_automaiton/configs"
	"github.com/gowaves/order_automaiton/middleware"
	"github.com/gowaves/order_automaiton/routes"
	"github.com/gowaves/order_automaiton/utils"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "eSeller Order Automaiton",
	})

	app.Use(cors.New())

	//run database
	configs.ConnectDB()

	// Create a route group for API version 1 ("/api/v1").
	v1Group := app.Group("/api/v1", logger.New())

	// Define a map to store comp_code and mobile based on companies
	companySecretKeys, err := utils.GetCompanySecretKeys()

	if err != nil {
		fmt.Println("Error getting company secret keys:", err)
		return
	}

	// Define a middleware
	companyCodeMiddleware := middleware.CompanyCodeMiddleware(companySecretKeys)
	jwtMiddleware := middleware.JwtMiddleware()
	authorizeMiddleware := middleware.AuthorizationMiddleware()

	//routes with middleware
	routes.CompanyRoute(v1Group, jwtMiddleware, companyCodeMiddleware, authorizeMiddleware)

	// Defer the disconnect and handle errors gracefully
	defer func() {
		if err := configs.DB.Disconnect(context.Background()); err != nil {
			fmt.Println("Error disconnecting from MongoDB:", err)
		} else {
			fmt.Println("Connection is disconnect! 111")
		}
	}()

	app.Listen(":6000")
}
