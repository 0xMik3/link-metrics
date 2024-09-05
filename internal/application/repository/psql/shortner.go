package psql

import (
	"errors"

	"github.com/0xMik3/link-metrics/internal/application/ports"
	"github.com/0xMik3/link-metrics/internal/domain"
	"github.com/go-xorm/xorm"
	"github.com/gofiber/fiber/v2/log"
)

type ShortenerRepository struct {
	engine *xorm.Engine
}

func NewShortenerRepository(e *xorm.Engine) ports.ShortenerRepo {
	return &ShortenerRepository{
		engine: e,
	}
}

func (r *ShortenerRepository) Create(url *domain.Url) error {
	affected, err := r.engine.InsertOne(url)
	if err != nil {
		log.Error("error ocurred while inserting record:", err)
		return err
	}
	if affected < 1 {
		return errors.New("could not insert record in db")
	}
	return nil
}

func (r *ShortenerRepository) GetByKey(key string) (*domain.Url, error) {
	url := domain.Url{
		Key: key,
	}
	has, err := r.engine.Get(&url)
	if err != nil {
		log.Error("error while retreiving url:", err)
		return nil, err
	}
	if !has {
		return nil, errors.New("not found")
	}
	return &url, nil
}

func (r *ShortenerRepository) GetById(id int64) (*domain.Url, error) {
	url := domain.Url{
		Id: id,
	}
	has, err := r.engine.Get(&url)
	if err != nil {
		log.Error("error while retreiving url:", err)
		return nil, err
	}
	if !has {
		return nil, errors.New("not found")
	}
	return &url, nil
}

func (r *ShortenerRepository) UpdateTotalClicks(id int64) error {
	url, err := r.GetById(id)
	if err != nil {
		return err
	}
	affected, err := r.engine.Table(new(domain.Url)).ID(id).Update(map[string]interface{}{"total_clicks": url.TotalClicks + 1})
	if err != nil {
		log.Error("error ocurred while updating total clicks:", err)
		return err
	}
	if affected < 1 {
		log.Error("could not update totalclicks on id:", id)
	}
	return nil
}
