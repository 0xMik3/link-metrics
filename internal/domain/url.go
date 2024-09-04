package domain

import "time"

type Url struct {
	Id          int64
	Url         string `xorm:"text"`
	Key         string `xorm:"varchar(6) unique index"`
	TotalClicks int64
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
}
