package interpolate

import (
	"github.com/ChrisMuir/address-parse-china/city"
	"github.com/ChrisMuir/address-parse-china/county"
	"github.com/ChrisMuir/address-parse-china/models"
	"reflect"
	"testing"
)

var countyOne = county.County{
	CountyName:   "东城",
	CountyCode:   110101,
	CityCode:     1101,
	ProvinceCode: 11,
}

var cityOne = city.City{
	CityName:     "辖区",
	CityCode:     1101,
	ProvinceCode: 11,
}

func TestInterpolateMissingValues(t *testing.T) {
	tests := []struct {
		testName    string
		geoInfo     models.GeoLocation
		cityMatch   city.City
		countyMatch county.County
		expected    models.GeoLocation
	}{
		{
			testName:    "Missing city and province, infer them from county",
			geoInfo:     models.GeoLocation{CountyCode: countyOne.CountyCode},
			cityMatch:   city.City{},
			countyMatch: countyOne,
			expected: models.GeoLocation{
				Province:     "北京",
				ProvinceCode: 11,
				City:         "辖区",
				CityCode:     1101,
				CountyCode:   countyOne.CountyCode,
			},
		},
		{
			testName:    "Missing city, infer it from county",
			geoInfo:     models.GeoLocation{CountyCode: countyOne.CountyCode, ProvinceCode: 11},
			cityMatch:   city.City{},
			countyMatch: countyOne,
			expected: models.GeoLocation{
				ProvinceCode: 11,
				City:         "辖区",
				CityCode:     1101,
				CountyCode:   countyOne.CountyCode,
			},
		},
		{
			testName:    "Missing province, infer it from city",
			geoInfo:     models.GeoLocation{CityCode: cityOne.CityCode},
			cityMatch:   cityOne,
			countyMatch: county.County{},
			expected: models.GeoLocation{
				Province:     "北京",
				ProvinceCode: 11,
				CityCode:     1101,
			},
		},
		{
			testName:    "Everything filled in, nothing to infer",
			geoInfo:     models.GeoLocation{CountyCode: 123, CityCode: 456, ProvinceCode: 789},
			cityMatch:   city.City{},
			countyMatch: county.County{},
			expected: models.GeoLocation{
				ProvinceCode: 789,
				CityCode:     456,
				CountyCode:   123,
			},
		},
		{
			testName:    "Nothing filled in, nothing to infer",
			geoInfo:     models.GeoLocation{},
			cityMatch:   city.City{},
			countyMatch: county.County{},
			expected:    models.GeoLocation{},
		},
	}

	for _, tt := range tests {
		resp := InterpolateMissingValues(tt.geoInfo, tt.cityMatch, tt.countyMatch)
		if !reflect.DeepEqual(resp, tt.expected) {
			t.Errorf("from InterpolateMissingValues(), for test %v, expected %v, instead got %v", tt.testName, tt.expected, resp)
		}
	}
}
