package rest

import (
	"net"

	"github.com/0xMik3/link-metrics/internal/application/ports"
	"github.com/0xMik3/link-metrics/internal/infra/adapters/driven/fiber"

	f "github.com/gofiber/fiber/v2"
)

type RestHandler struct {
	Fiber     *fiber.FiberServer
	Shortener ports.ShortnerService
}

func NewRestHandler(shortener ports.ShortnerService) *RestHandler {
	fiber := fiber.NewFiberApp()
	return &RestHandler{
		Fiber:     fiber,
		Shortener: shortener,
	}
}

func (r *RestHandler) InitializeRoutes() {
	api := r.Fiber.App.Group("api")
	api.Get("/ping", func(c *f.Ctx) error {
		return c.SendString("pong")
	})

	shorten := api.Group("shorten")
	shorten.Post("/", r.ShortenUrl)
	shorten.Get("/:key", r.GetUrl)
}

func (r *RestHandler) Start(l net.Listener) error {
	return r.Fiber.App.Listener(l)
}
