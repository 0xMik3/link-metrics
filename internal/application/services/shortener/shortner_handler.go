package shortener

import (
	"math/rand"
	"time"

	"github.com/0xMik3/link-metrics/internal/domain"
)

func (s *ShortenerService) Create(url *domain.Url) error {
	key := generate_key()

	url.Key = key
	url.TotalClicks = 0

	return nil
}

func generate_key() string {
	const alpha_num = "abcdefghijklmnopqrstuvwxyz0123456789"
	const length = 6

	rand.New(rand.NewSource(time.Now().UnixNano()))
	shortKey := make([]byte, length)
	for i := range shortKey {
		shortKey[i] = alpha_num[rand.Intn(len(alpha_num))]
	}
	return string(shortKey)
}
