package shortener

import (
	"context"

	"github.com/0xMik3/link-metrics/internal/application/ports"
)

type ShortenerService struct {
	globalCtx     context.Context
	shortenerRepo ports.ShortenerRepo
}

func NewShortenerService(ctx context.Context, shortnerRepo ports.ShortenerRepo) ports.ShortnerService {
	return &ShortenerService{
		globalCtx:     ctx,
		shortenerRepo: shortnerRepo,
	}
}
