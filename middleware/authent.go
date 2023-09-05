package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/meles-zawude-e/Emp_Manag_Sys_IN_SantimPay/model"
)

const  SecretKey = "secret"

func Authentication(c *fiber.Ctx) error {
	var emp model.Employee
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience:  []string{},
		ID:        "",
		IssuedAt:  &jwt.Time{},
		Issuer:    emp.ID.String(),
		NotBefore: &jwt.Time{},
		Subject:   "",
	})
	token, err := claim.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "Could not login"})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{"message": "success"})
}
