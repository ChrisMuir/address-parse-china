package match

import (
	"math"
	"strings"

	"github.com/ChrisMuir/address-parse-china/city"
	"github.com/ChrisMuir/address-parse-china/county"
	"github.com/ChrisMuir/address-parse-china/province"
)

func Province(location string) province.Province {
	var match province.Province
	var lowestSubStrIdx = math.MaxInt
	for _, currProv := range province.Provinces {
		idx := strings.Index(location, currProv.ProvinceName)
		if idx >= 0 && idx < lowestSubStrIdx {
			match.ProvinceName = currProv.ProvinceName
			match.ProvinceCode = currProv.ProvinceCode
			lowestSubStrIdx = idx
		}
	}
	return match
}

func Cities(location string) []city.City {
	var matches []city.City
	var lowestSubStrIdx = math.MaxInt
	var lowestCityName string
	for _, currCity := range city.Cities {
		idx := strings.Index(location, currCity.CityName)
		if idx >= 0 && idx < lowestSubStrIdx {
			matches = []city.City{currCity}
			lowestSubStrIdx = idx
			lowestCityName = currCity.CityName
			continue
		}
		if currCity.CityName == lowestCityName {
			matches = append(matches, currCity)
		}
	}
	return matches
}

func Counties(location string) []county.County {
	var matches []county.County
	var lowestSubStrIdx = math.MaxInt
	var lowestCountyName string
	for _, currCounty := range county.Counties {
		idx := strings.Index(location, currCounty.CountyName)
		if idx >= 0 && idx < lowestSubStrIdx {
			matches = []county.County{currCounty}
			lowestSubStrIdx = idx
			lowestCountyName = currCounty.CountyName
			continue
		}
		if currCounty.CountyName == lowestCountyName {
			matches = append(matches, currCounty)
		}
	}
	return matches
}
