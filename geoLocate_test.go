package address_parse_china

import (
	"github.com/ChrisMuir/address-parse-china/city"
	"github.com/ChrisMuir/address-parse-china/county"
	"github.com/ChrisMuir/address-parse-china/models"
	"github.com/ChrisMuir/address-parse-china/province"
	"reflect"
	"testing"
)

func TestGeoLocate(t *testing.T) {
	tests := []struct {
		input    []string
		expected []models.GeoLocation
	}{
		{
			input: []string{"大连市甘井子区南关岭街道姚工街101号"},
			expected: []models.GeoLocation{
				{
					Address:      "大连市甘井子区南关岭街道姚工街101号",
					Province:     "辽宁",
					ProvinceCode: 21,
					City:         "大连",
					CityCode:     2102,
					County:       "甘井子",
					CountyCode:   210211,
				},
			},
		},
		{
			input: []string{"大连市甘井子区南关岭街道姚工街101号", "浙江省杭州市余杭区径山镇小古城村"},
			expected: []models.GeoLocation{
				{
					Address:      "大连市甘井子区南关岭街道姚工街101号",
					Province:     "辽宁",
					ProvinceCode: 21,
					City:         "大连",
					CityCode:     2102,
					County:       "甘井子",
					CountyCode:   210211,
				},
				{
					Address:      "浙江省杭州市余杭区径山镇小古城村",
					Province:     "浙江",
					ProvinceCode: 33,
					City:         "杭州",
					CityCode:     3301,
					County:       "余杭",
					CountyCode:   330110,
				},
			},
		},
		{
			input: []string{"我喜欢猫"},
			expected: []models.GeoLocation{
				{
					Address:      "我喜欢猫",
					Province:     "",
					ProvinceCode: 0,
					City:         "",
					CityCode:     0,
					County:       "",
					CountyCode:   0,
				},
			},
		},
		{
			input: []string{"I like cats"},
			expected: []models.GeoLocation{
				{
					Address:      "I like cats",
					Province:     "",
					ProvinceCode: 0,
					City:         "",
					CityCode:     0,
					County:       "",
					CountyCode:   0,
				},
			},
		},
	}

	for _, tt := range tests {
		resp := GeoLocate(tt.input)
		if !reflect.DeepEqual(resp, tt.expected) {
			t.Errorf("from GeoLocate(), expected %v, instead got %v, on input %v", tt.expected, resp, tt.input)
		}
	}
}

func BenchmarkGeoLocate(b *testing.B) {
	locs := []string{
		"大连市甘井子区南关岭街道姚工街101号",
	}
	for i := 0; i < b.N; i++ {
		GeoLocate(locs)
	}
}

func TestGetProvinceData(t *testing.T) {
	resp := GetProvinceData()

	// Check len
	expectedLen := 31
	if len(resp) != expectedLen {
		t.Errorf("Province data expected len %v, instead got %v", expectedLen, len(resp))
	}

	// Check one expected value
	expectedVal := province.Province{
		"上海",
		31,
	}
	seen := false
	for _, currProv := range resp {
		if currProv.ProvinceName == expectedVal.ProvinceName && currProv.ProvinceCode == expectedVal.ProvinceCode {
			seen = true
			break
		}
	}
	if !seen {
		t.Errorf("Province data expected to see value %v, but value not found in pkg data", expectedVal)
	}
}

func TestGetCityData(t *testing.T) {
	resp := GetCityData()

	// Check len
	expectedLen := 341
	if len(resp) != expectedLen {
		t.Errorf("City data expected len %v, instead got %v", expectedLen, len(resp))
	}

	// Check one expected value
	expectedVal := city.City{
		"晋城",
		1405,
		14,
	}
	seen := false
	for _, currCity := range resp {
		if currCity.CityName == expectedVal.CityName &&
			currCity.CityCode == expectedVal.CityCode &&
			currCity.ProvinceCode == expectedVal.ProvinceCode {
			seen = true
			break
		}
	}
	if !seen {
		t.Errorf("City data expected to see value %v, but value not found in pkg data", expectedVal)
	}
}

func TestGetCountyDataData(t *testing.T) {
	resp := GetCountyData()

	// Check len
	expectedLen := 2978
	if len(resp) != expectedLen {
		t.Errorf("County data expected len %v, instead got %v", expectedLen, len(resp))
	}

	// Check one expected value
	expectedVal := county.County{
		"河北唐山海港经济开发",
		130274,
		1302,
		13,
	}
	seen := false
	for _, currCounty := range resp {
		if currCounty.CountyName == expectedVal.CountyName &&
			currCounty.CountyCode == expectedVal.CountyCode &&
			currCounty.CityCode == expectedVal.CityCode &&
			currCounty.ProvinceCode == expectedVal.ProvinceCode {
			seen = true
			break
		}
	}
	if !seen {
		t.Errorf("County data expected to see value %v, but value not found in pkg data", expectedVal)
	}
}
