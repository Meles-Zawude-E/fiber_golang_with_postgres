package login

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meles-zawude-e/Emp_Manag_Sys_IN_SantimPay/database"
	"github.com/meles-zawude-e/Emp_Manag_Sys_IN_SantimPay/model"
)

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var emp model.Employee
	database.DB.Db.Where("email=?", data["email"]).First(&emp)
	database.DB.Db.Where("role=?", data["role"]).First(&emp)
	database.DB.Db.Where("password=?", data["password"]).First(&emp)

	if emp.Role != data["role"] && emp.Email != data["email"] {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "you have no permission to access this"})

	}
	if emp.Password != data["password"] {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "incorrect password"})
	}
	
	return c.JSON(fiber.Map{"Message": "Successfully Login"})

}




