package ports

import "github.com/0xMik3/link-metrics/internal/domain"

type ShortnerService interface {
	Create(url *domain.Url) (string, error)
	GetByKey(key string) (*domain.Url, error)
	UpdateTotalClicks(id int64) error
}
