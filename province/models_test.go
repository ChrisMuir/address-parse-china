package province

import (
	"reflect"
	"testing"
)

func TestMakeProvinceNameMap(t *testing.T) {
	pMap := makeProvinceNameMap()
	testVal := "河北"
	resp := pMap[testVal]
	expected := 13
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Province map, expected %v, instead got %v, on input %v", expected, resp, testVal)
	}
}

func TestMakeProvinceCodeMap(t *testing.T) {
	ProvinceNameMap := makeProvinceNameMap()
	ProvinceCodeMap := makeProvinceCodeMap()
	testVal := "河北"
	provCode := ProvinceNameMap[testVal]
	resp := ProvinceCodeMap[provCode]
	expected := Province{
		"河北",
		13,
	}
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("Province code map, expected %v, instead got %v, on input %v", expected, resp, testVal)
	}
}
