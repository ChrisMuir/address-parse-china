package substring

const MIN_SUBSTRING_LEN = 2
const MAX_SUBSTRING_LEN = 14

// Get every substring of the input location string, from len 2 thru len 14
// Returns a map, in which keys are substrings, values are the starting index of the substring within the input string
// Example input: "这是一个测试"
// Output: map[这是:0 这是一:0 这是一个:0 这是一个测:0 这是一个测试:0 是一:1 是一个:1 是一个测:1 是一个测试:1 一个:2 一个测:2 一个测试:2 个测:3 个测试:3 测试:4]
func GetAllSubstrings(location string) map[string]int {
	// Build rune->byte offsets so substrings can be sliced without re-encoding runes repeatedly
	offsets := make([]int, 0, 32)
	for i := range location {
		offsets = append(offsets, i)
	}
	offsets = append(offsets, len(location))
	runeLen := len(offsets) - 1
	if runeLen < MIN_SUBSTRING_LEN {
		return map[string]int{}
	}

	b := []byte(location)
	n := len(b)

	// prefix rolling hash over bytes to allow O(1) substring hash
	const base uint64 = 1315423911
	h := make([]uint64, n+1)
	pow := make([]uint64, n+1)
	pow[0] = 1
	for i := 0; i < n; i++ {
		h[i+1] = h[i]*base + uint64(b[i])
		pow[i+1] = pow[i] * base
	}

	maxRune := MAX_SUBSTRING_LEN
	if maxRune > runeLen {
		maxRune = runeLen
	}
	estimated := runeLen * (maxRune - MIN_SUBSTRING_LEN + 1)
	type sample struct{ idx, start, length int }
	hashMap := make(map[uint64]sample, estimated)
	collisions := make(map[string]int)

	// the unique length values of all of the pkg data strings
	subStringLengths := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14}
	for _, L := range subStringLengths {
		for i := 0; i <= runeLen-L; i++ {
			startByte := offsets[i]
			endByte := offsets[i+L]
			Lb := endByte - startByte
			hash := h[endByte] - h[startByte]*pow[Lb]
			key := hash ^ (uint64(Lb) << 32)
			if s, ok := hashMap[key]; !ok {
				hashMap[key] = sample{idx: i, start: startByte, length: Lb}
			} else {
				if s.length == Lb {
					equal := true
					for k := 0; k < Lb; k++ {
						if b[s.start+k] != b[startByte+k] {
							equal = false
							break
						}
					}
					if equal {
						continue
					}
				}
				ks := string(b[startByte:endByte])
				if _, ok := collisions[ks]; !ok {
					collisions[ks] = i
				}
			}
		}
	}

	res := make(map[string]int, len(hashMap)+len(collisions))
	for _, s := range hashMap {
		res[string(b[s.start:s.start+s.length])] = s.idx
	}
	for k, v := range collisions {
		if _, ok := res[k]; !ok {
			res[k] = v
		}
	}
	return res
}
