package shortener

import (
	"context"

	"github.com/0xMik3/link-metrics/internal/application/ports"
	"github.com/ua-parser/uap-go/uaparser"
)

type ShortenerService struct {
	globalCtx     context.Context
	shortenerRepo ports.ShortenerRepo
	parser        *uaparser.Parser
}

func NewShortenerService(ctx context.Context, shortnerRepo ports.ShortenerRepo) ports.ShortnerService {
	return &ShortenerService{
		globalCtx:     ctx,
		shortenerRepo: shortnerRepo,
		parser:        uaparser.NewFromSaved(),
	}
}
