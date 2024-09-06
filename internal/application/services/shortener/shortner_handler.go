package shortener

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/0xMik3/link-metrics/internal/domain"
	"github.com/gofiber/fiber/v2/log"
)

func (s *ShortenerService) Generate_key() string {
	const alpha_num = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	const length = 8

	rand.New(rand.NewSource(time.Now().UnixNano()))
	shortKey := make([]byte, length)
	for i := range shortKey {
		shortKey[i] = alpha_num[rand.Intn(len(alpha_num))]
	}
	return string(shortKey)
}

func (s *ShortenerService) Create(url *domain.Url) (string, error) {
	key := s.Generate_key()

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

func (s *ShortenerService) CheckIpLocation(ip string) (*domain.IpLocation, error) {
	ipLocation := domain.IpLocation{}
	requestUrl := fmt.Sprintf("https://api.iplocation.net/?ip=%s", ip)
	res, err := http.Get(requestUrl)
	if err != nil {
		log.Error("could not get data from ip:", err)
		return nil, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&ipLocation)
	return &ipLocation, nil
}

func (s *ShortenerService) HandleClick(id int64, ip string, referer string, userAgent string) {
	var wg sync.WaitGroup
	metric := domain.Metric{
		UrlId:   id,
		Referer: referer,
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.UpdateTotalClicks(id)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		client := s.parser.Parse(userAgent)
		metric.Device = client.Os.Family
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ipLocation, err := s.CheckIpLocation(ip)
		if err == nil {
			metric.CountryCode = ipLocation.CountryCode
			metric.CountryName = ipLocation.CountryName
		}
	}()
	wg.Wait()

	s.shortenerRepo.CreateMetric(&metric)
}
