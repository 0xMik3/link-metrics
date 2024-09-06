package domain

type IpLocation struct {
	Ip          string `json:"ip"`
	CountryName string `json:"country_name"`
	CountryCode string `json:"country_code2"`
}
