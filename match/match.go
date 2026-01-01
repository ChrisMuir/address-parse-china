package match

import (
	"math"

	"github.com/ChrisMuir/address-parse-china/city"
	"github.com/ChrisMuir/address-parse-china/county"
	"github.com/ChrisMuir/address-parse-china/province"
)

func Province(subStrMap map[string]int) province.Province {
	var match province.Province
	var lowestSubStrIdx = math.MaxInt
	if len(subStrMap) > len(province.NameMap) {
		for currProv, currProvCode := range province.NameMap {
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
			provCode, ok := province.NameMap[currSubStr]
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

func Cities(subStrMap map[string]int) []city.City {
	var matches []city.City
	var lowestSubStrIdx = math.MaxInt
	if len(subStrMap) > len(city.NameMap) {
		for currCitySubStr, currCities := range city.NameMap {
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
			cities, ok := city.NameMap[currSubStr]
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

func Counties(subStrMap map[string]int) []county.County {
	var matches []county.County
	var lowestSubStrIdx = math.MaxInt
	if len(subStrMap) > len(county.NameMap) {
		for currCountySubStr, currCounties := range county.NameMap {
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
			counties, ok := county.NameMap[currSubStr]
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
