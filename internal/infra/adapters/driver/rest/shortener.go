package rest

import (
	"strings"

	"github.com/gofiber/fiber/v2/log"

	"github.com/0xMik3/link-metrics/internal/application/services/shortener/dtos"
	"github.com/0xMik3/link-metrics/internal/domain"
	"github.com/gofiber/fiber/v2"
)

func (r *RestHandler) ShortenUrl(c *fiber.Ctx) error {
	var url dtos.ShortenUrlDto
	if err := c.BodyParser(&url); err != nil {
		log.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}
	key, err := r.Shortener.Create(&domain.Url{
		Url:  url.Url,
		Name: url.Name,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot short url",
		})
	}
	return c.JSON(
		fiber.Map{
			"key": key,
		},
	)
}

func (r *RestHandler) GetUrl(c *fiber.Ctx) error {
	xForwardedFor := c.Get("X-Forwarded-For", "")
	ips := []string{}
	if xForwardedFor != "" {
		ips = strings.Split(xForwardedFor, ",")
	}
	clientIp := "86.106.87.229"
	if len(ips) > 0 {
		clientIp = ips[0]
	}
	key := c.Params("key")
	referer := c.Get("Referer", "anonymous")

	url, err := r.Shortener.GetByKey(key)
	if err != nil {
		if err.Error() == "not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Url not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot retreive url",
		})
	}

	r.Shortener.HandleClick(url.Id, clientIp, referer)

	return c.JSON(
		fiber.Map{
			"url": url.Url,
		},
	)
}
