package domain

import "time"

type Url struct {
	Id          int64
	Url         string `xorm:"text"`
	Name        string `xorm:"index"`
	Key         string `xorm:"varchar(8) unique index"`
	TotalClicks int64
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
}
