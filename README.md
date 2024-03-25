# address-parse-china

Golang port of [this](https://github.com/ChrisMuir/geolocChina/) R package

Golang package that takes Chinese location/address stings as input and returns geolocation data for each input string. 
This package does not rely on an external API service to perform geocoding. It is dependency-free.

The functions are designed to work with business name strings and address strings. For each input string, the following 
values are returned: `Province`, `City`, `County`, `Provincial geocode`, `City geocode`, and `County geocode`.

The package is using geo substrings and associated geocodes from [Administrative-divisions-of-China](https://github.com/modood/Administrative-divisions-of-China) 
on GitHub, put together by user [modood](https://github.com/modood).

For more info on geolocating Chinese strings, see [this](https://pdfs.semanticscholar.org/ca9d/2d09d0a2420a7ce398e14ed43f8cd7464705.pdf) 
2016 paper on the subject. The geolocation function works by using both substring matching and geocodes to validate child 
regions and infer parent regions.

## Installation

```
go get github.com/ChrisMuir/address-parse-china
```

## Example Usage

Function `GeoLocate()`

```go
package main

import (
	"encoding/json"
	"fmt"
	apc "github.com/ChrisMuir/address-parse-china"
)

func main() {
	locs := []string{
		"大连市甘井子区南关岭街道姚工街101号",
		"浙江省杭州市余杭区径山镇小古城村",
		"大连御洋食品有限公司",
		"徐州雅莲连锁超市有限公司",
		"四川省南充市阆中市公园路63号",
		"同德县鲜肉蔬菜配送行",
	}
	geoLocs := apc.GeoLocate(locs)
	geoLocsJson, _ := json.Marshal(geoLocs)
	fmt.Println(fmt.Sprintf("%+v", string(geoLocsJson)))
}
```

```json
[
  {
    "Address": "大连市甘井子区南关岭街道姚工街101号",
    "Province": "辽宁",
    "ProvinceCode": 21,
    "City": "大连",
    "CityCode": 2102,
    "County": "甘井子",
    "CountyCode": 210211
  },
  {
    "Address": "浙江省杭州市余杭区径山镇小古城村",
    "Province": "浙江",
    "ProvinceCode": 33,
    "City": "杭州",
    "CityCode": 3301,
    "County": "余杭",
    "CountyCode": 330110
  },
  {
    "Address": "大连御洋食品有限公司",
    "Province": "辽宁",
    "ProvinceCode": 21,
    "City": "大连",
    "CityCode": 2102,
    "County": "",
    "CountyCode": 0
  },
  {
    "Address": "徐州雅莲连锁超市有限公司",
    "Province": "江苏",
    "ProvinceCode": 32,
    "City": "徐州",
    "CityCode": 3203,
    "County": "",
    "CountyCode": 0
  },
  {
    "Address": "四川省南充市阆中市公园路63号",
    "Province": "四川",
    "ProvinceCode": 51,
    "City": "南充",
    "CityCode": 5113,
    "County": "阆中",
    "CountyCode": 511381
  },
  {
    "Address": "同德县鲜肉蔬菜配送行",
    "Province": "青海",
    "ProvinceCode": 63,
    "City": "海南藏族",
    "CityCode": 6325,
    "County": "同德",
    "CountyCode": 632522
  }
]
```

There are also functions that expose the pkg data used for geolocation, `GetProvinceData()`, `GetCityData()`, and `GetCountyData()`

```go
package main

import (
	"encoding/json"
	"fmt"
	apc "github.com/ChrisMuir/address-parse-china"
)

func main() {
	provs := apc.GetProvinceData()
	provsJson, _ := json.Marshal(provs)
	fmt.Println(fmt.Sprintf("%+v", string(provsJson)))
}

```

```json
[
  {
    "ProvinceName": "北京",
    "ProvinceCode": 11
  },
  {
    "ProvinceName": "天津",
    "ProvinceCode": 12
  },
  {
    "ProvinceName": "河北",
    "ProvinceCode": 13
  },
  {
    "ProvinceName": "山西",
    "ProvinceCode": 14
  },
  {
    "ProvinceName": "内蒙古",
    "ProvinceCode": 15
  },
  {
    "ProvinceName": "辽宁",
    "ProvinceCode": 21
  },
  {
    "ProvinceName": "吉林",
    "ProvinceCode": 22
  },
  {
    "ProvinceName": "黑龙江",
    "ProvinceCode": 23
  },
  {
    "ProvinceName": "上海",
    "ProvinceCode": 31
  },
  {
    "ProvinceName": "江苏",
    "ProvinceCode": 32
  },
  {
    "ProvinceName": "浙江",
    "ProvinceCode": 33
  },
  {
    "ProvinceName": "安徽",
    "ProvinceCode": 34
  },
  {
    "ProvinceName": "福建",
    "ProvinceCode": 35
  },
  {
    "ProvinceName": "江西",
    "ProvinceCode": 36
  },
  {
    "ProvinceName": "山东",
    "ProvinceCode": 37
  },
  {
    "ProvinceName": "河南",
    "ProvinceCode": 41
  },
  {
    "ProvinceName": "湖北",
    "ProvinceCode": 42
  },
  {
    "ProvinceName": "湖南",
    "ProvinceCode": 43
  },
  {
    "ProvinceName": "广东",
    "ProvinceCode": 44
  },
  {
    "ProvinceName": "广西壮族",
    "ProvinceCode": 45
  },
  {
    "ProvinceName": "海南",
    "ProvinceCode": 46
  },
  {
    "ProvinceName": "重庆",
    "ProvinceCode": 50
  },
  {
    "ProvinceName": "四川",
    "ProvinceCode": 51
  },
  {
    "ProvinceName": "贵州",
    "ProvinceCode": 52
  },
  {
    "ProvinceName": "云南",
    "ProvinceCode": 53
  },
  {
    "ProvinceName": "西藏",
    "ProvinceCode": 54
  },
  {
    "ProvinceName": "陕西",
    "ProvinceCode": 61
  },
  {
    "ProvinceName": "甘肃",
    "ProvinceCode": 62
  },
  {
    "ProvinceName": "青海",
    "ProvinceCode": 63
  },
  {
    "ProvinceName": "宁夏回族",
    "ProvinceCode": 64
  },
  {
    "ProvinceName": "新疆维吾尔",
    "ProvinceCode": 65
  }
]
```

## Benchmark

```go
package address_parse_china

import "testing"

func BenchmarkGeoLocate(b *testing.B) {
	locs := []string{
		"大连市甘井子区南关岭街道姚工街101号",
	}
	for i := 0; i < b.N; i++ {
		GeoLocate(locs)
	}
}
```
```bash
~/address-parse-china % go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/ChrisMuir/address-parse-china
BenchmarkGeoLocate-8       40614             26821 ns/op
PASS
ok      github.com/ChrisMuir/address-parse-china        1.614s
```