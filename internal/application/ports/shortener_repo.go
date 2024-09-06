package ports

import "github.com/0xMik3/link-metrics/internal/domain"

type ShortenerRepo interface {
	Create(url *domain.Url) error
	GetByKey(key string) (*domain.Url, error)
	UpdateTotalClicks(id int64) error
	CreateMetric(metric *domain.Metric) error
}
