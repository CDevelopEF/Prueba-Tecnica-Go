package utils
import 	"github.com/gofiber/fiber/v2"


var CodigosDeError = map[string]map[int]string{
	"en": {
		fiber.StatusBadRequest:         "Bad request",
		fiber.StatusUnauthorized:       "Unauthorized",
		fiber.StatusPaymentRequired:    "Payment required",
		fiber.StatusForbidden:          "Forbidden",
		fiber.StatusNotFound:           "Not found",
		fiber.StatusMethodNotAllowed:   "Method not allowed",
		fiber.StatusRequestTimeout:     "Request timeout",
		fiber.StatusInternalServerError: "Internal server error",
		fiber.StatusNotImplemented:     "Not implemented",
		fiber.StatusBadGateway:         "Bad gateway",
		fiber.StatusServiceUnavailable: "Service unavailable",
		fiber.StatusGatewayTimeout:     "Gateway timeout",
	},
	"es": {
		fiber.StatusBadRequest:         "Solicitud incorrecta",
		fiber.StatusUnauthorized:       "No autorizado",
		fiber.StatusPaymentRequired:    "Pago requerido",
		fiber.StatusForbidden:          "Prohibido",
		fiber.StatusNotFound:           "No encontrado",
		fiber.StatusMethodNotAllowed:   "Método no permitido",
		fiber.StatusRequestTimeout:     "Tiempo de espera de la solicitud",
		fiber.StatusInternalServerError: "Error interno del servidor",
		fiber.StatusNotImplemented:     "No implementado",
		fiber.StatusBadGateway:         "Puerta de enlace incorrecta",
		fiber.StatusServiceUnavailable: "Servicio no disponible",
		fiber.StatusGatewayTimeout:     "Tiempo de espera de la puerta de enlace",
	},
}

