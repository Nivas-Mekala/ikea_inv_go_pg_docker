package routes

import (
	"github.com/Nivas-Mekala/ikea_inv_go_pg_docker/database"
	"github.com/Nivas-Mekala/ikea_inv_go_pg_docker/internal/models"
	"github.com/gofiber/fiber/v2"
)

func SaveProducts(c *fiber.Ctx) error {

	var productReq models.ProductRequest
	var data models.Product

	if err := c.BodyParser(&productReq); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	for _, inv := range productReq.Products {
		data.Product_Name = inv.Name

		for _, art := range inv.Contain_Articles {
			data.Amount_Of = art.Amount_Of
			data.Article_Id = art.Article_Id
			database.Database.DB.Create(&data)
		}

	}

	return c.Status(200).SendString("Inventory Data saved successfully !!")

}

func GetAllProducts(c *fiber.Ctx) error {
	proucts := []models.Product{}

	database.Database.DB.Find(&proucts)
	return c.Status(200).JSON(proucts)
}
