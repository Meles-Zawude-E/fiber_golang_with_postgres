package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/meles-zawude-e/Emp_Manag_Sys_IN_SantimPay/database"
	"github.com/meles-zawude-e/Emp_Manag_Sys_IN_SantimPay/model"
)

// register customer
func RegisterCustomer(c *fiber.Ctx) error {
	db := database.DB.Db
	emp := new(model.Employee)
	//store it
	err := c.BodyParser(emp)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "Failed", "Message": "Something is wrong", "data": err})
	}
	err = db.Create(&emp).Error
	if err != nil {
		return c.Status(500).JSONP(fiber.Map{
			"status": "Failed", "Message": "Failed to connect Database", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{
		"status": "Success", "Message": "Employeeis successfuly created", "Data": emp})
}

// check all customer
func GetAllCustomer(c *fiber.Ctx) error {
	db := database.DB.Db
	var emps []model.Employee
	//find all employees
	db.Find(&emps)
	if len(emps) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": "Failed", "Message": "Employee not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{
		"status": "success", "Message": "Employee is found", "data": emps})

}

// get only one customer
func GetCustomerByID(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")

	var emp model.Employee
	db.Find(&emp, "id=?", id)
	if emp.ID == uuid.Nil {
		return c.Status(404).JSONP(fiber.Map{
			"status": "failed", "Message": "employee is not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{
		"status": "success", "Message": "Employee is found", "data": emp})

}

//Update Customer

func UpdateCustomer(c *fiber.Ctx) error {
	type updateCustomer struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Phone    string `json:"phone"`
		Photo    string `json:"photo"`
	}
	db := database.DB.Db
	var emp model.Employee

	id := c.Params("id")
	db.Find(&emp, "id=?", id)
	if emp.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "failed", "Message": "employee not found", "data": nil})
	}
	var updateEmployeeData updateCustomer
	err := c.BodyParser(&updateEmployeeData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "fail", "Message": "someting is wrong", "data": err})
	}
	emp.Username = updateEmployeeData.Username
	emp.Email = updateEmployeeData.Email
	emp.Photo = updateEmployeeData.Photo
	emp.Phone = updateEmployeeData.Phone

	db.Save(&emp)
	return c.Status(200).JSON(fiber.Map{
		"status": "success", "Message": "employee is successfuly updated", "data": emp})
}

//Delete customer

func DeleteCustomer(c *fiber.Ctx) error {
	db := database.DB.Db
	var emp model.Employee
	id := c.Params("id")

	db.Find(&emp, "id=?", id)
	if emp.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "failed", "Message": "Not get employee", "data": nil})
	}
	err := db.Delete(&emp, "id=?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "fail", "Message": "Failed to delete", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "Success", "Message": "Successfully deleted", "data": emp})

}


