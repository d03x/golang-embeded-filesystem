package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Logging(app *fiber.Ctx) error {
	log.Println(string(app.Request().URI().FullURI()))
	return app.Next()
}
