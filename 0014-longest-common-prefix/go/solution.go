func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	if len(strs[0]) == 0 {
		return ""
	}

	slices.Sort(strs)
	slices.Reverse(strs)

	commonChars := []string{}
	breakFound := false

	for i, c := range strs[0] {
		if breakFound {
			break
		}
		if i > 0 {
			commonChars = append(commonChars, string(strs[0][i-1]))
		}
		for _, str := range strs[1:] {
			if i == len(str) {
				breakFound = true
				break
			} else if rune(str[i]) != c {
				breakFound = true
				break
			}
		}
	}

	if !breakFound {
		commonChars = append(commonChars, string(strs[0][len(strs[0])-1]))
	}

	commonPrefix := strings.Join(commonChars, "")
	return commonPrefix
}