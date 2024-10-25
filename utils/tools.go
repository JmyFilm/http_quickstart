package utils

// RemoveDuplicates 切片去重
func RemoveDuplicates[T comparable](input []T) (ret []T) {
	uniqueMap := make(map[T]struct{})

	for idx := 0; idx < len(input); idx++ {
		if _, ok := uniqueMap[input[idx]]; !ok {
			uniqueMap[input[idx]] = struct{}{}
			ret = append(ret, input[idx])
		}
	}

	return
}
