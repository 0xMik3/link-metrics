package fiber

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type FiberServer struct {
	App *fiber.App
}

func NewFiberApp() *FiberServer {
	return &FiberServer{
		App: fiber.New(),
	}
}

func (f *FiberServer) Start(port string) error {
	log.Info("App ready in port:", port)
	return f.App.Listen(fmt.Sprintf(":%v", port))
}
