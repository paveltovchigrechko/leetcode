func twoSum(nums []int, target int) []int {
	numsToPos := make(map[int]int, len(nums))
	result := []int{}
	for i, num := range nums {
		numToTarget := target - num

		secondNum, ok := numsToPos[numToTarget]
		if !ok {
			numsToPos[num] = i
			continue
		} else {
			result = append(result, secondNum, i)
			break
		}
	}

	return result
}