package models

//const MIN_SUBSTRING_LEN = 2
//const MAX_SUBSTRING_LEN = 15

type GeoLocation struct {
	Address      string
	Province     string
	ProvinceCode int
	City         string
	CityCode     int
	County       string
	CountyCode   int
}
