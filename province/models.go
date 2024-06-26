package province

type Province struct {
	ProvinceName string
	ProvinceCode int
}

var NameMap = makeProvinceNameMap()
var CodeMap = makeProvinceCodeMap()

func makeProvinceNameMap() map[string]int {
	provinceMap := make(map[string]int)
	for _, prov := range Provinces {
		provinceMap[prov.ProvinceName] = prov.ProvinceCode
	}
	return provinceMap
}

func makeProvinceCodeMap() map[int]Province {
	provMap := make(map[int]Province)
	for _, currProv := range Provinces {
		provMap[currProv.ProvinceCode] = currProv
	}
	return provMap
}

var Provinces = []Province{
	{
		"北京",
		11,
	},
	{
		"天津",
		12,
	},
	{
		"河北",
		13,
	},
	{
		"山西",
		14,
	},
	{
		"内蒙古",
		15,
	},
	{
		"辽宁",
		21,
	},
	{
		"吉林",
		22,
	},
	{
		"黑龙江",
		23,
	},
	{
		"上海",
		31,
	},
	{
		"江苏",
		32,
	},
	{
		"浙江",
		33,
	},
	{
		"安徽",
		34,
	},
	{
		"福建",
		35,
	},
	{
		"江西",
		36,
	},
	{
		"山东",
		37,
	},
	{
		"河南",
		41,
	},
	{
		"湖北",
		42,
	},
	{
		"湖南",
		43,
	},
	{
		"广东",
		44,
	},
	{
		"广西壮族",
		45,
	},
	{
		"海南",
		46,
	},
	{
		"重庆",
		50,
	},
	{
		"四川",
		51,
	},
	{
		"贵州",
		52,
	},
	{
		"云南",
		53,
	},
	{
		"西藏",
		54,
	},
	{
		"陕西",
		61,
	},
	{
		"甘肃",
		62,
	},
	{
		"青海",
		63,
	},
	{
		"宁夏回族",
		64,
	},
	{
		"新疆维吾尔",
		65,
	},
}
