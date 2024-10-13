package routes

import (
	"strconv"

	"github.com/Nivas-Mekala/ikea_inv_go_pg_docker/database"
	"github.com/Nivas-Mekala/ikea_inv_go_pg_docker/internal/models"
	"github.com/gofiber/fiber/v2"
)

func SaveInventory(c *fiber.Ctx) error {
	var inputRequest models.InventoryRequest
	var data models.Inventory

	if err := c.BodyParser(&inputRequest); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	for _, inv := range inputRequest.Inventory {
		data.Article_Id = inv.Article_Id
		data.Name = inv.Name
		data.Stock = inv.Stock

		database.Database.DB.Create(&data)
	}

	return c.Status(200).SendString("Inventory Data saved successfully !!")
}

func GetAllInventory(c *fiber.Ctx) error {
	inventory := []models.Inventory{}
	database.Database.DB.Find(&inventory)

	return c.Status(200).JSON(inventory)
}

func findInventory(id int, inventory *models.Inventory) error {
	database.Database.DB.Find(&inventory, "article_id=?", strconv.Itoa(id))
	return nil
}

func GetInventory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	response := models.Inventory{}
	if err := findInventory(id, &response); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON(response)
}

func UpdateInventory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	response := models.Inventory{}

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findInventory(id, &response); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	updateInv := models.Inventory{}
	if err := c.BodyParser(&updateInv); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	database.Database.DB.Save(&updateInv)

	return c.Status(200).JSON(updateInv)

}

func DeleteInventory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var inventory models.Inventory
	if err := findInventory(id, &inventory); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.DB.Delete(&inventory)

	return c.Status(200).SendString("Inventory has been deleted")
}
