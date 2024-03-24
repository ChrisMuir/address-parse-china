package filter

import (
	"github.com/ChrisMuir/address-parse-china/city"
	"github.com/ChrisMuir/address-parse-china/county"
	"github.com/ChrisMuir/address-parse-china/models"
)

func Cities(matches []city.City, geoInfo models.GeoLocation) city.City {
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

func Counties(matches []county.County, geoInfo models.GeoLocation) county.County {
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
