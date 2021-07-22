package hangul

var (
	start = rune(44032)
	end   = rune(55204)
)

func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false
	for _, v := range s {
		if start >= v && v < end {
			index := int(v - start)
			result = (index % numEnds) != 0
		}
	}
	return result
}
