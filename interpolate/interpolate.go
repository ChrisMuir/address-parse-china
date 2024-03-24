package interpolate

import (
	"github.com/ChrisMuir/address-parse-china/city"
	"github.com/ChrisMuir/address-parse-china/county"
	"github.com/ChrisMuir/address-parse-china/models"
	"github.com/ChrisMuir/address-parse-china/province"
)

// Uses child matches to infer parent values
func InterpolateMissingValues(geoInfo models.GeoLocation, cityMatch city.City, countyMatch county.County) models.GeoLocation {
	// If county was matched and city was not matched, infer city from county match
	if geoInfo.CityCode == 0 && geoInfo.CountyCode > 0 {
		geoInfo.CityCode = countyMatch.CityCode
		geoInfo.City = city.CodeMap[countyMatch.CityCode].CityName
	}

	// If city or county was matched and province was not matched, infer province from city or county match
	if geoInfo.ProvinceCode == 0 && geoInfo.CityCode > 0 {
		if countyMatch.ProvinceCode > 0 {
			geoInfo.ProvinceCode = countyMatch.ProvinceCode
		} else {
			geoInfo.ProvinceCode = cityMatch.ProvinceCode
		}
		geoInfo.Province = province.CodeMap[geoInfo.ProvinceCode].ProvinceName
	}
	return geoInfo
}
