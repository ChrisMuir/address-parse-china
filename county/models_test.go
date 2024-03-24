package county

import (
	"reflect"
	"testing"
)

func TestMakeCountyMap(t *testing.T) {
	cMap := makeCountyMap()
	testVal := "朝阳"
	resp := cMap[testVal]
	expected := []County{
		{
			"朝阳",
			110105,
			1101,
			11,
		},
		{
			"朝阳",
			211321,
			2113,
			21,
		},
		{
			"朝阳",
			220104,
			2201,
			22,
		},
	}
	if !reflect.DeepEqual(resp, expected) {
		t.Errorf("County map, expected %v, instead got %v, on input %v", expected, resp, testVal)
	}
}
