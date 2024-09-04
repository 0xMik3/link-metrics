package ports

import "github.com/0xMik3/link-metrics/internal/domain"

type ShortenerRepo interface {
	Create(url *domain.Url) error
}
