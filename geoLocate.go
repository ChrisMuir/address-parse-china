package address_parse_china

import (
	"github.com/ChrisMuir/address-parse-china/city"
	"github.com/ChrisMuir/address-parse-china/county"
	"github.com/ChrisMuir/address-parse-china/filter"
	"github.com/ChrisMuir/address-parse-china/interpolate"
	"github.com/ChrisMuir/address-parse-china/match"
	"github.com/ChrisMuir/address-parse-china/models"
	"github.com/ChrisMuir/address-parse-china/province"
	"github.com/ChrisMuir/address-parse-china/substring"
)

// GeoLocate takes in an array of Chinese address/business strings, returns the Province, Provincial Code, City,
// City Code, County, and County Code for each array element.
func GeoLocate(locations []string) []models.GeoLocation {
	var geoLocations = make([]models.GeoLocation, len(locations))
	for idx, location := range locations {
		geoLocations[idx] = getGeoLocation(location)
	}
	return geoLocations
}

// Given an input Chinese location/address string, determine the Province,
// City, and County of the input string (and the associated geocodes for all
// three).
func getGeoLocation(location string) models.GeoLocation {
	var geoInfo models.GeoLocation
	geoInfo.Address = location

	// Get every possible substring of location string. Substring len 2 thru 14 (14 is the longest county)
	// For each of the three substring matching steps below (province, city, county), follow this logic (using province
	//  as an example):
	// if number of location substrings is greater than then number of provinces, then:
	//     Look each province up in the location substring map
	// else if number of provinces if greater than the number of location substrings, then:
	//     Look each location substring up in the province map
	subStrMap := substring.GetAllSubstrings(location)

	// Province substring matches
	matchingProvince := match.Province(subStrMap)
	geoInfo.Province = matchingProvince.ProvinceName
	geoInfo.ProvinceCode = matchingProvince.ProvinceCode

	// City substring matches
	matchingCities := match.Cities(subStrMap)
	matchingCity := filter.Cities(matchingCities, geoInfo)
	geoInfo.City = matchingCity.CityName
	geoInfo.CityCode = matchingCity.CityCode

	// County substring matches
	matchingCounties := match.Counties(subStrMap)
	matchingCounty := filter.Counties(matchingCounties, geoInfo)
	geoInfo.County = matchingCounty.CountyName
	geoInfo.CountyCode = matchingCounty.CountyCode

	// Fill in missing fields using geocode interpolation
	geoInfo = interpolate.InterpolateMissingValues(geoInfo, matchingCity, matchingCounty)

	return geoInfo
}

func GetProvinceData() []province.Province {
	return province.Provinces
}

func GetCityData() []city.City {
	return city.Cities
}

func GetCountyData() []county.County {
	return county.Counties
}
