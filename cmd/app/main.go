package main

import (
	"log"

	"github.com/Nivas-Mekala/ikea_inv_go_pg_docker/database"
	"github.com/Nivas-Mekala/ikea_inv_go_pg_docker/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectToDatabase()
	app := fiber.New()
	setUpRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func setUpRoutes(app *fiber.App) {
	app.Get("/ikea", welcome)

	app.Post("/ikea/saveInventory", routes.SaveInventory)
	app.Get("/ikea/getAllInventory", routes.GetAllInventory)
	app.Get("/ikea/getInventory/:id", routes.GetInventory)
	app.Put("/ikea/updateInventory/:id", routes.UpdateInventory)
	app.Delete("/ikea/removeInventory/:id", routes.DeleteInventory)

	app.Post("/ikea/saveProducts", routes.SaveProducts)
	app.Get("/ikea/getProducts", routes.GetAllProducts)
}

func welcome(c *fiber.Ctx) error {
	return c.Status(200).SendString("Welcome To IKEA")
}
