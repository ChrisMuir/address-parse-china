package substring

const MIN_SUBSTRING_LEN = 2
const MAX_SUBSTRING_LEN = 15

// Get every substring of the input location string, from len 2 thru len 14
// Returns a map, in which keys are substrings, values are the starting index of the substring within the input string
// Example input: "这是一个测试"
// Output: map[这是:0 这是一:0 这是一个:0 这是一个测:0 这是一个测试:0 是一:1 是一个:1 是一个测:1 是一个测试:1 一个:2 一个测:2 一个测试:2 个测:3 个测试:3 测试:4]
func GetAllSubstrings(location string) map[string]int {
	var currSubStr []rune
	loc := []rune(location)
	maxLen := getMaxSubStrLen(len(loc))
	subStrMap := make(map[string]int)
	for i := 0; i < len(loc); i++ {
		currSubStr = make([]rune, maxLen)
		for j, subStrCounter := i, 0; j < (maxLen + i); j, subStrCounter = j+1, subStrCounter+1 {
			if j == len(loc) {
				break
			}
			currSubStr[subStrCounter] = loc[j]
			if subStrCounter+1 < MIN_SUBSTRING_LEN {
				continue
			}
			subStrMap[string(currSubStr[:subStrCounter+1])] = i
		}
	}
	return subStrMap
}

func getMaxSubStrLen(locationLen int) int {
	var maxLen int
	if locationLen > MAX_SUBSTRING_LEN {
		maxLen = MAX_SUBSTRING_LEN
	} else {
		maxLen = locationLen
	}
	return maxLen
}
