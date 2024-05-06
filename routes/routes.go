package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/auth/login", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"activo": "si"})
	})

	// GET /auth/login
	app.Get("/auth/login", func(c *fiber.Ctx) error {
		
		return nil
	})

	// POST /auth
	app.Post("/auth", func(c *fiber.Ctx) error {
		// Lógica para crear un nuevo usuario
		return nil
	})

	// GET /auth/:userId
	app.Get("/auth/:userId", func(c *fiber.Ctx) error {
		// Lógica para devolver la información del usuario con el ID especificado
		return nil
	})

	// PUT /auth/:userId
	app.Put("/auth/:userId", func(c *fiber.Ctx) error {
		// Lógica para modificar la información del usuario con el ID especificado
		return nil
	})

	// DELETE /auth/:userId
	app.Delete("/auth/:userId", func(c *fiber.Ctx) error {
		// Lógica para marcar para borrado lógico al usuario con el ID especificado
		return nil
	})

	// POST /sites
	app.Post("/sites", func(c *fiber.Ctx) error {
		// Lógica para crear un nuevo sitio
		return nil
	})

	// GET /sites
	app.Get("/sites", func(c *fiber.Ctx) error {
		// Lógica para devolver todos los sitios
		return nil
	})

	// GET /sites/:siteId
	app.Get("/sites/:siteId", func(c *fiber.Ctx) error {
		// Lógica para devolver el sitio con el ID especificado
		return nil
	})

	// PUT /sites/:siteId
	app.Put("/sites/:siteId", func(c *fiber.Ctx) error {
		// Lógica para modificar el sitio con el ID especificado
		return nil
	})

	// DELETE /sites/:siteId
	app.Delete("/sites/:siteId", func(c *fiber.Ctx) error {
		// Lógica para marcar para borrado lógico al sitio con el ID especificado
		return nil
	})
}
