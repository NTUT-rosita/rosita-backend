package sliceConverter

func Int16ToInt(o []int16) []int {
	result := make([]int, len(o))
	for i, v := range o {
		result[i] = int(v)
	}
	return result
}
