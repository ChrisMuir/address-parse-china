package address_parse_china

import (
	"github.com/ChrisMuir/address-parse-china/city"
	"github.com/ChrisMuir/address-parse-china/county"
	"github.com/ChrisMuir/address-parse-china/province"
)

const MIN_SUBSTRING_LEN = 2
const MAX_SUBSTRING_LEN = 15

// Takes in an array of Chinese address/business strings, returns the Province, Provincial Code, City, City Code,
// County, and County Code for each array element.
func GeoLocate(locations []string) []GeoLocation {
	var geoLocations = make([]GeoLocation, len(locations))
	for idx, location := range locations {
		geoLocations[idx] = getGeoLocation(location)
	}
	return geoLocations
}

// Given an input Chinese location/address string, determine the Province,
// City, and County of the input string (and the associated geocodes for all
// three).
func getGeoLocation(location string) GeoLocation {
	var geoInfo GeoLocation
	geoInfo.Address = location
	loc := stringToRuneArray(location)

	// Get every possible substring of location string. Substring len 2 thru 14 (14 is the longest county)
	// For each of the three substring matching steps below (province, city, county), follow this logic (using province
	//  as an example):
	// if number of location substrings is greater than then number of provinces, then:
	//     Look each province up in the location substring map
	// else if number of provinces if greater than the number of location substrings, then:
	//     Look each location substring up in the province map
	subStrMap := getAllSubstrings(loc)

	// Province:
	matchingProvince := getProvinceMatches(subStrMap)
	geoInfo.Province = matchingProvince.ProvinceName
	geoInfo.ProvinceCode = matchingProvince.ProvinceCode

	// City
	matchingCities := getCityMatches(subStrMap)
	matchingCity := filterCityMatches(matchingCities, geoInfo)
	geoInfo.City = matchingCity.CityName
	geoInfo.CityCode = matchingCity.CityCode

	// County
	matchingCounties := getCountyMatches(subStrMap)
	matchingCounty := filterCountyMatches(matchingCounties, geoInfo)
	geoInfo.County = matchingCounty.CountyName
	geoInfo.CountyCode = matchingCounty.CountyCode

	return geoInfo
}

// Get every substring of the input location string, from len 2 thru len 14
func getAllSubstrings(location []rune) map[string]int {
	maxLen := getMaxSubStrLen(location)
	subStrMap := make(map[string]int)
	for i := 0; i < maxLen; i++ {
		currSubStr := []rune{}
		for j := i; j < maxLen; j++ {
			currSubStr = append(currSubStr, location[j])
			if len(currSubStr) < MIN_SUBSTRING_LEN {
				continue
			}
			subStrMap[string(currSubStr)] = i
		}
	}
	return subStrMap
}

func getMaxSubStrLen(location []rune) int {
	var maxLen int
	if len(location) > MAX_SUBSTRING_LEN {
		maxLen = MAX_SUBSTRING_LEN
	} else {
		maxLen = len(location)
	}
	return maxLen
}

func stringToRuneArray(s string) []rune {
	return []rune(s)
}

func getProvinceMatches(subStrMap map[string]int) province.Province {
	var match province.Province
	var lowestSubStrIdx = 9999
	if len(subStrMap) > len(province.Map) {
		for currProv, currProvCode := range province.Map {
			subStrIdx, ok := subStrMap[currProv]
			if !ok {
				continue
			}
			if subStrIdx < lowestSubStrIdx {
				match.ProvinceName = currProv
				match.ProvinceCode = currProvCode
				lowestSubStrIdx = subStrIdx
			}
		}
	} else {
		for currSubStr, currSubStrIdx := range subStrMap {
			provCode, ok := province.Map[currSubStr]
			if !ok {
				continue
			}
			if currSubStrIdx < lowestSubStrIdx {
				match.ProvinceName = currSubStr
				match.ProvinceCode = provCode
				lowestSubStrIdx = currSubStrIdx
			}
		}
	}
	return match
}

func getCityMatches(subStrMap map[string]int) []city.City {
	var matches []city.City
	var lowestSubStrIdx = 9999
	if len(subStrMap) > len(city.Map) {
		for currCitySubStr, currCities := range city.Map {
			subStrIdx, ok := subStrMap[currCitySubStr]
			if !ok {
				continue
			}
			if subStrIdx < lowestSubStrIdx {
				matches = currCities
				lowestSubStrIdx = subStrIdx
			}
		}
	} else {
		for currSubStr, currSubStrIdx := range subStrMap {
			cities, ok := city.Map[currSubStr]
			if !ok {
				continue
			}
			if currSubStrIdx < lowestSubStrIdx {
				matches = cities
				lowestSubStrIdx = currSubStrIdx
			}
		}
	}
	return matches
}

func getCountyMatches(subStrMap map[string]int) []county.County {
	var matches []county.County
	var lowestSubStrIdx = 9999
	if len(subStrMap) > len(county.Map) {
		for currCountySubStr, currCounties := range county.Map {
			subStrIdx, ok := subStrMap[currCountySubStr]
			if !ok {
				continue
			}
			if subStrIdx < lowestSubStrIdx {
				matches = currCounties
				lowestSubStrIdx = subStrIdx
			}
		}
	} else {
		for currSubStr, currSubStrIdx := range subStrMap {
			counties, ok := county.Map[currSubStr]
			if !ok {
				continue
			}
			if currSubStrIdx < lowestSubStrIdx {
				matches = counties
				lowestSubStrIdx = currSubStrIdx
			}
		}
	}
	return matches
}

func filterCityMatches(matches []city.City, geoInfo GeoLocation) city.City {
	// If there are no city substring matches, return nil
	if len(matches) < 1 {
		return city.City{}
	}

	// If there is no matching province, return the first city substring match
	if geoInfo.ProvinceCode == 0 {
		return matches[0]
	}

	// If there is a matching province, then the first two digits of any city match must be equal to the province code, in
	// order to be the right match.
	if geoInfo.ProvinceCode > 0 {
		for _, currCity := range matches {
			if currCity.ProvinceCode == geoInfo.ProvinceCode {
				return currCity
			}
		}
	}

	return city.City{}
}

func filterCountyMatches(matches []county.County, geoInfo GeoLocation) county.County {
	// If there are no county substring matches, return nil
	if len(matches) < 1 {
		return county.County{}
	}

	// If there are no matching province or city, return the first county substring match
	if geoInfo.ProvinceCode == 0 && geoInfo.CityCode == 0 {
		return matches[0]
	}

	// If there is a matching city, then the first four digits of any county match must be equal to the city code, in
	// order to be the right match.
	if geoInfo.CityCode > 0 {
		for _, currCounty := range matches {
			if currCounty.CityCode == geoInfo.CityCode {
				return currCounty
			}
		}
	}

	// If there is a matching province, then the first two digits of any county match must be equal to the province code, in
	// order to be the right match.
	if geoInfo.ProvinceCode > 0 {
		for _, currCounty := range matches {
			if currCounty.ProvinceCode == geoInfo.ProvinceCode {
				return currCounty
			}
		}
	}

	return county.County{}
}
