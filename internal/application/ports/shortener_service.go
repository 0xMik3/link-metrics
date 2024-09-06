package ports

import "github.com/0xMik3/link-metrics/internal/domain"

type ShortnerService interface {
	Generate_key() string
	Create(url *domain.Url) (string, error)
	GetByKey(key string) (*domain.Url, error)
	UpdateTotalClicks(id int64) error
	CheckIpLocation(ip string) (*domain.IpLocation, error)
	HandleClick(id int64, ip string, referer string)
}
