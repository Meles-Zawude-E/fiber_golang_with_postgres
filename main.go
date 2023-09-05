package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/meles-zawude-e/Emp_Manag_Sys_IN_SantimPay/database"
	"github.com/meles-zawude-e/Emp_Manag_Sys_IN_SantimPay/router"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{AllowCredentials: true}))

	router.SetupRouter(app)
     
	app.Use(func (c *fiber.Ctx) error {
		return c.SendStatus(404)
	 })
	

	app.Listen(":3000")
}
