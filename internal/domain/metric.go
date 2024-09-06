package domain

import "time"

type Metric struct {
	Id          int64
	CountryCode string `xorm:"varchar(2)"`
	CountryName string
	Referer     string
	Device      string
	UrlId       int64     `xorm:"index"`
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
}

type MetricGroup struct {
	Url  `xorm:"extends"`
	Name string
}

func (MetricGroup) TableName() string {
	return "metric"
}
