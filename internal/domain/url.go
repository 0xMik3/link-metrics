package domain

import "time"

type Url struct {
	Id          int64
	Url         string
	Key         string `xorm:"unique, index"`
	TotalClicks int64
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
}
