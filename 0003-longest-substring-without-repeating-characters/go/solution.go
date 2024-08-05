func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	var substring string
	var length int

	for _, c := range s {
		if !strings.ContainsRune(substring, c) {
			substring += string(c)
		} else {
			if len(substring) > length {
				length = len(substring)
			}
			newString := strings.SplitAfterN(substring, string(c), 2)[1]
			substring = newString + string(c)
		}
	}

	if len(substring) > length {
		length = len(substring)
	}
	return length
}