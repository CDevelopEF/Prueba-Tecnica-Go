package middlewares
import (
	"strings"
	"github.com/gofiber/fiber/v2"

)

func obtenerIdiomaUsuario(c *fiber.Ctx) string {
	idioma := c.Get("Accept-Language")
	if idioma == "" {
		return "ru"
	}
	idiomasPreferidos := strings.Split(idioma, ";")
	primerIdioma := strings.Split(idiomasPreferidos[0], ";")[0]
	return primerIdioma
}

func LanguageMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		idiomaUsuario := obtenerIdiomaUsuario(c)
		c.Locals("idioma", idiomaUsuario)
		return c.Next()
	}
}
// Un middleware que recupere de la base de datos el estado actual de la conexión.
func ConnectionStatusMiddleware() fiber.Handler {
	return func( c *fiber.Ctx) error {
		return c.Next()
	}
}