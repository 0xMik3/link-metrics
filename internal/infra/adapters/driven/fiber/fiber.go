package fiber

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
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

	return f.App.Listen(fmt.Sprintf(":%v", port))
}
