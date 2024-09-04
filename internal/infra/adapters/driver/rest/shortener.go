package rest

import (
	"log"

	"github.com/0xMik3/link-metrics/internal/application/services/shortener/dtos"
	"github.com/0xMik3/link-metrics/internal/domain"
	"github.com/gofiber/fiber/v2"
)

func (r *RestHandler) ShortenUrl(c *fiber.Ctx) error {
	var url dtos.ShortenUrlDto
	if err := c.BodyParser(&url); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}
	key, err := r.Shortener.Create(&domain.Url{
		Url: url.Url,
	})

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot create user",
		})
	}
	return c.JSON(
		fiber.Map{
			"key": key,
		},
	)
}
