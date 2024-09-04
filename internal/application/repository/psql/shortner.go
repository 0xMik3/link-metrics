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
