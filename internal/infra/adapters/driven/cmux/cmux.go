package cmux

import (
	"net"

	"github.com/gofiber/fiber/v2/log"
	"github.com/soheilhy/cmux"
)

type CmuxConfig struct {
	httpListener net.Listener
	cmux         cmux.CMux
}

func NewCmuxConfig(port string) *CmuxConfig {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalw("Error starting cmux server", "error", err)
	}
	m := cmux.New(l)
	anyL := m.Match(cmux.Any())
	return &CmuxConfig{
		httpListener: anyL,
		cmux:         m,
	}
}

func (c *CmuxConfig) HttpListener() net.Listener {
	return c.httpListener
}

func (c *CmuxConfig) Start() {
	err := c.cmux.Serve()
	if err != nil {
		log.Fatalf("Error starting cmux server: %v", err)
	}
}
