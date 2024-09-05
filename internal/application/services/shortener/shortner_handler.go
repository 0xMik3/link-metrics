package shortener

import (
	"errors"
	"math/rand"
	"time"

	"github.com/0xMik3/link-metrics/internal/domain"
)

func (s *ShortenerService) Create(url *domain.Url) (string, error) {
	key := generate_key()

	url.Key = key
	url.TotalClicks = 0

	err := s.shortenerRepo.Create(url)
	if err != nil {
		return "", err
	}
	return key, nil
}

func (s *ShortenerService) GetByKey(key string) (*domain.Url, error) {
	if len(key) != 8 {
		return nil, errors.New("invalid key")
	}
	url, err := s.shortenerRepo.GetByKey(key)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (s *ShortenerService) UpdateTotalClicks(id int64) error {
	return s.shortenerRepo.UpdateTotalClicks(id)
}

func generate_key() string {
	const alpha_num = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	const length = 8

	rand.New(rand.NewSource(time.Now().UnixNano()))
	shortKey := make([]byte, length)
	for i := range shortKey {
		shortKey[i] = alpha_num[rand.Intn(len(alpha_num))]
	}
	return string(shortKey)
}
