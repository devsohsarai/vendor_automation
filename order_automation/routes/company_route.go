package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gowaves/order_automaiton/controllers"
)

func CompanyRoute(group fiber.Router, jwtMiddleware, companyCodeMiddleware, authorizeMiddleware fiber.Handler) {
	//All routes related to users comes here

	group.Post("/company", controllers.CreateCompany)
	group.Post("/company/auth", controllers.AuthProcess)
	group.Get("/company/:mobile", controllers.GetCompanyDetails)
	group.Get("/companies", companyCodeMiddleware, jwtMiddleware, controllers.GetAllCompanies)
	group.Get("/companydetail/:mobile", companyCodeMiddleware, jwtMiddleware, authorizeMiddleware, controllers.GetACompany)
	group.Get("/stack", companyCodeMiddleware, jwtMiddleware, controllers.Stack)
}
