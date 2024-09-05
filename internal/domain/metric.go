package domain

type Metric struct {
	Id      int64
	Country string
	Referer string
	Device  string
	UrlId   int64 `xorm:"index"`
}

type MetricGroup struct {
	Url  `xorm:"extends"`
	Name string
}

func (MetricGroup) TableName() string {
	return "metric"
}
