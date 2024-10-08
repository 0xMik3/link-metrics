package psql

import (
	"fmt"

	"github.com/0xMik3/link-metrics/internal/config"
	"github.com/0xMik3/link-metrics/internal/domain"
	"github.com/go-xorm/xorm"
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/lib/pq"
)

func Connect(config *config.Config) (*xorm.Engine, error) {
	connection_string := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName, config.DbSslMode)

	engine, err := xorm.NewEngine("postgres", connection_string)
	if err != nil {
		log.Error("could not connect to db:", err)
		return nil, err
	}
	err = engine.Ping()
	if err != nil {
		log.Error("failed to ping db: ", err)
		return nil, err
	}
	log.Info("connected to db succesfully")
	return engine, nil
}

func Sync_tables(en *xorm.Engine) {
	err := en.Sync(new(domain.Url), new(domain.Metric))
	if err != nil {
		log.Error("creation error", err)
		return
	}
	log.Info("Successfully synced")
}
