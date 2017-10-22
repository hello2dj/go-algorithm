package string

func KMPSearch(source, pattern string) int {
	si := 0
	pi := 0
	sourceBytes := []byte(source)
	patternBytes := []byte(pattern)
	nextArray := genNextArray(pattern)
	for si < len(source) && pi < len(pattern) {
		if pi == -1 || (sourceBytes[si] == patternBytes[pi]) {
			si++
			pi++
		} else {
			pi = nextArray[pi]
		}
	}
	if pi == len(pattern) {
		return si - pi
	} else {
		return -1
	}
}

func genNextArray(pattern string) []int {
	patternBytes := []byte(pattern)
	length := len(pattern)
	array := make([]int, length)
	array[0] = -1
	prefixI := -1
	suffixI := 0
	for suffixI < length-1 {
		if prefixI == -1 || (patternBytes[prefixI] == patternBytes[suffixI]) {
			prefixI++
			suffixI++
			if patternBytes[prefixI] != patternBytes[suffixI] {
				array[suffixI] = prefixI
			} else {
				array[suffixI] = array[prefixI]
			}
		} else {
			prefixI = array[prefixI]
		}
	}
	return array
}
