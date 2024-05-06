package main

import (
	"example.com/prueba/database"
	midl "example.com/prueba/middlewares"
	appRoutes "example.com/prueba/routes"
	str "example.com/prueba/structs"
	"example.com/prueba/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {

	//Instancia Global de la BD
	database.GetInstance()
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			idioma := c.Locals("idioma").(string)
			codigoError := err.(*fiber.Error).Code
			mensajeError := utils.CodigosDeError[idioma][codigoError]
			return c.Status(codigoError).JSON(str.Mensaje{
				Code:    codigoError,
				Message: mensajeError,
			})
		},
	})
	app.Use(midl.LanguageMiddleware())
	var pingHandler = func(ctx *fiber.Ctx) error {
		value := ctx.Params("name")
		return ctx.Status(200).JSON(fiber.Map{"nombre": value, "idioma": ctx.Locals("idioma").(string)})

	}
	appRoutes.SetupRoutes(app)
	api := app.Group("/api")
	api.Get("/ping/:name", pingHandler)
	app.Listen(":8000")
}
