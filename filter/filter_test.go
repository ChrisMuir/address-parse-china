package filter

import (
	"github.com/ChrisMuir/address-parse-china/city"
	"github.com/ChrisMuir/address-parse-china/county"
	"github.com/ChrisMuir/address-parse-china/models"
	"reflect"
	"testing"
)

var cityOne = city.City{
	CityName:     "辖区",
	CityCode:     1101,
	ProvinceCode: 11,
}

var cityTwo = city.City{
	CityName:     "唐山",
	CityCode:     1302,
	ProvinceCode: 13,
}

var countyOne = county.County{
	CountyName:   "东城",
	CountyCode:   110101,
	CityCode:     1101,
	ProvinceCode: 11,
}

var countyTwo = county.County{
	CountyName:   "栾城",
	CountyCode:   130111,
	CityCode:     1301,
	ProvinceCode: 13,
}

func TestCities(t *testing.T) {
	tests := []struct {
		testName string
		matches  []city.City
		geoInfo  models.GeoLocation
		expected city.City
	}{
		{
			testName: "In Cities(), no city matches",
			matches:  []city.City{},
			geoInfo:  models.GeoLocation{},
			expected: city.City{},
		},
		{
			testName: "In Cities(), return first matching city, no province match",
			matches:  []city.City{cityOne, cityTwo},
			geoInfo:  models.GeoLocation{},
			expected: cityOne,
		},
		{
			testName: "In Cities(), return first matching city with a province match",
			matches:  []city.City{cityOne, cityTwo},
			geoInfo: models.GeoLocation{
				ProvinceCode: 13,
			},
			expected: cityTwo,
		},
		{
			testName: "In Cities(), return no matching city because none match the province",
			matches:  []city.City{cityOne, cityTwo},
			geoInfo: models.GeoLocation{
				ProvinceCode: 15,
			},
			expected: city.City{},
		},
	}

	for _, tt := range tests {
		resp := Cities(tt.matches, tt.geoInfo)
		if !reflect.DeepEqual(resp, tt.expected) {
			t.Errorf("from filter(), for test %v, expected %v, instead got %v", tt.testName, tt.expected, resp)
		}
	}
}

func TestCounties(t *testing.T) {
	tests := []struct {
		testName string
		matches  []county.County
		geoInfo  models.GeoLocation
		expected county.County
	}{
		{
			testName: "In Counties(), no county matches",
			matches:  []county.County{},
			geoInfo:  models.GeoLocation{},
			expected: county.County{},
		},
		{
			testName: "In Counties(), return first matching county, no province or city match",
			matches:  []county.County{countyOne, countyTwo},
			geoInfo:  models.GeoLocation{},
			expected: countyOne,
		},
		{
			testName: "In Counties(), return first matching country with a province match",
			matches:  []county.County{countyOne, countyTwo},
			geoInfo: models.GeoLocation{
				ProvinceCode: 13,
			},
			expected: countyTwo,
		},
		{
			testName: "In Counties(), return first matching county with a city match",
			matches:  []county.County{countyOne, countyTwo},
			geoInfo: models.GeoLocation{
				CityCode: 1301,
			},
			expected: countyTwo,
		},
		{
			testName: "In Counties(), return no matching city because none match the city or province",
			matches:  []county.County{countyOne, countyTwo},
			geoInfo: models.GeoLocation{
				ProvinceCode: 15,
				CityCode:     1501,
			},
			expected: county.County{},
		},
	}

	for _, tt := range tests {
		resp := Counties(tt.matches, tt.geoInfo)
		if !reflect.DeepEqual(resp, tt.expected) {
			t.Errorf("from filter(), for test %v, expected %v, instead got %v", tt.testName, tt.expected, resp)
		}
	}
}
