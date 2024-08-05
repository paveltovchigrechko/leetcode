func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	str := strconv.Itoa(x)

	reversedRunes := []rune(str)
	for i, j := 0, len(reversedRunes)-1; i < len(reversedRunes)/2; i, j = i+1, j-1 {
		reversedRunes[i], reversedRunes[j] = reversedRunes[j], reversedRunes[i]
	}

	reversedStr := string(reversedRunes)

	return str == reversedStr
}