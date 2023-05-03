package main

import (
	"fmt"

	"github.com/Rama-85/Apiproject/database"
	"github.com/Rama-85/Apiproject/models"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/Employee", EmployeesList)
	app.Post("/Employee", CreateEmployee)
	app.Get("/Employee/:id", GetProductById)
	app.Delete("/Employee/:id", DeleteEmployee)
	app.Put("/Employee/:id", UpdateEmployee)
}

func EmployeesList(c *fiber.Ctx) error {
	emps := []models.Employee{}
	database.DB.Db.Find(&emps)
	return c.Status(200).JSON(emps)
}
func CreateEmployee(c *fiber.Ctx) error {
	emp := new(models.Employee)
	if err := c.BodyParser(emp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&emp)
	return c.Status(200).JSON(emp)
}

func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	var emp models.Employee

	database.DB.Db.Find(&emp, "id = ?", id)

	if emp.id == ' ' {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No not present", "data": nil})
	}

	var uemp models.Employee

	err := c.BodyParser(&uemp)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	emp.name = uemp.name
	emp.salary = uemp.salary

	database.DB.Db.Save(&emp)
	return c.JSON(fiber.Map{"data": emp})

}

func GetProductById(c *fiber.Ctx) error {
	id := c.Params("id")
	var emps models.Employee
	result := database.DB.Db.Find(&emps, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&emps)
}
func DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	var demp models.Employee
	if result := database.DB.Db.First(&demp, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	database.DB.Db.Delete(&demp)
	return c.Status(200).JSON(&demp)
}

func ProjectList(c *fiber.Ctx) error {
	projects := []models.Projects{}
	database.DB.Db.Find(&projects)
	return c.Status(200).JSON(projects)
}
