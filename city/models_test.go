package city

import (
	"reflect"
	"testing"
)

func TestMakeCityNameMap(t *testing.T) {
	cMap := makeCityNameMap()
	testVal := "辖区"
	resp := cMap[testVal]
	expected := []City{
		{
			"辖区",
			1101,
			11,
		},
		{
			"辖区",
			1201,
			12,
		},
		{
			"辖区",
			3101,
			31,
		},
		{
			"辖区",
			5001,
			50,
		},
	}
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("City map, expected %v, instead got %v, on input %v", expected, resp, testVal)
	}
}

func TestMakeCityCodeMap(t *testing.T) {
	cityNameMap := makeCityNameMap()
	cityCodeMap := makeCityCodeMap()
	testVal := "石家庄"
	nameVal := cityNameMap[testVal]
	resp := cityCodeMap[nameVal[0].CityCode]
	expected := City{
		"石家庄",
		1301,
		13,
	}
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("City code map, expected %v, instead got %v, on input %v", expected, resp, testVal)
	}
}
