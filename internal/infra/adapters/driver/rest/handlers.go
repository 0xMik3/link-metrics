package rest

import (
	"net"

	"github.com/0xMik3/link-metrics/internal/infra/adapters/driven/fiber"

	f "github.com/gofiber/fiber/v2"
)

type RestHandler struct {
	Fiber *fiber.FiberServer
}

func NewRestHandler() *RestHandler {
	fiber := fiber.NewFiberApp()
	return &RestHandler{
		Fiber: fiber,
	}
}

func (r *RestHandler) InitializeRoutes() {
	api := r.Fiber.App.Group("api")
	api.Get("/ping", func(c *f.Ctx) error {
		return c.SendString("pong")
	})
}

func (r *RestHandler) Start(l net.Listener) error {
	return r.Fiber.App.Listener(l)
}
