package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meles-zawude-e/Emp_Manag_Sys_IN_SantimPay/handler"
	"github.com/meles-zawude-e/Emp_Manag_Sys_IN_SantimPay/login"
)

func SetupRouter(app *fiber.App) {
	//grouping customer
	api := app.Group("/api")
	v1 := api.Group("/customer")

	v1.Get("/", handler.GetAllCustomer)
	v1.Get("/:id", handler.GetCustomerByID)
	v1.Post("/", handler.RegisterCustomer)
	v1.Put("/:id", handler.UpdateCustomer)
	v1.Delete("/:id", handler.DeleteCustomer)

	api.Post("/login", login.Login)

}
