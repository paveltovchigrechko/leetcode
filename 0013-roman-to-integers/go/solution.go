func romanToInt(s string) int {
	symbolValue := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var num, curValue, nextValue int
	skipValue := false

	for i := len(s) - 1; i >= 0; i-- {
		if skipValue {
			skipValue = false
			continue
		}

		curValue = symbolValue[rune(s[i])]
		if i > 0 {
			nextValue = symbolValue[rune(s[i-1])]
		}

		if curValue > nextValue {
			num += curValue
			num -= nextValue
			skipValue = true
		} else {
			num += curValue
		}
	}

	return num
}