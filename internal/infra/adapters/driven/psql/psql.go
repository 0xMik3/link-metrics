package psql

import (
	"fmt"

	"github.com/0xMik3/link-metrics/internal/config"
	"github.com/go-xorm/xorm"
	"github.com/gofiber/fiber/v2/log"
)

func Connect(config *config.Config) (*xorm.Engine, error) {
	connection_string := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)

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
