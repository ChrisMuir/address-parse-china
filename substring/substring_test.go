package substring

import (
	"testing"
)

func TestGetAllSubstrings_SingleRune(t *testing.T) {
	m := GetAllSubstrings("你")
	if len(m) != 0 {
		t.Fatalf("expected empty map for input length 1, got %d", len(m))
	}
}

func TestGetAllSubstrings_Example(t *testing.T) {
	s := "这是一个测试"
	m := GetAllSubstrings(s)

	expected := map[string]int{
		"这是":     0,
		"这是一":    0,
		"这是一个":   0,
		"这是一个测":  0,
		"这是一个测试": 0,
		"是一":     1,
		"是一个":    1,
		"是一个测":   1,
		"是一个测试":  1,
		"一个":     2,
		"一个测":    2,
		"一个测试":   2,
		"个测":     3,
		"个测试":    3,
		"测试":     4,
	}

	for k, want := range expected {
		idx, ok := m[k]
		if !ok {
			t.Fatalf("expected substring %q to be present", k)
		}
		if idx != want {
			t.Fatalf("substring %q: expected index %d, got %d", k, want, idx)
		}
	}
}

func TestGetAllSubstrings_MaxLenAndIndices(t *testing.T) {
	s := "abcdefghijklmnopqrst" // 20 runes
	runes := []rune(s)
	m := GetAllSubstrings(s)
	maxSubStringSeen := -1
	maxSubString := "blah"

	for k, idx := range m {
		l := len([]rune(k))
		if l < MIN_SUBSTRING_LEN || l > MAX_SUBSTRING_LEN {
			t.Fatalf("substring %q has invalid length %d", k, l)
		}
		if idx < 0 || idx >= len(runes) {
			t.Fatalf("invalid index %d for substring %q", idx, k)
		}
		if idx+l > len(runes) {
			t.Fatalf("substring %q length %d at index %d out of bounds", k, l, idx)
		}
		if string(runes[idx:idx+l]) != k {
			t.Fatalf("substring %q does not match runes slice at index %d", k, idx)
		}
		if l > maxSubStringSeen {
			maxSubStringSeen = l
			maxSubString = k
		}
	}
	if maxSubStringSeen != MAX_SUBSTRING_LEN {
		t.Fatalf("expected max substring length to be %v, instead got %v. Max substring returned is: %v", MAX_SUBSTRING_LEN, maxSubStringSeen, maxSubString)
	}
}
