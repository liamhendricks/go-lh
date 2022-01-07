package sorting

// O(nlogn)
func mergeSort(data []int) []int {
	if len(data) == 1 {
		return data
	}

	left := mergeSort(data[:len(data)/2])
	right := mergeSort(data[len(data)/2:])
	return mergeArrs(left, right)
}

func mergeArrs(left, right []int) []int {
	newArr := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			newArr = append(newArr, left[i])
			i++
		} else {
			newArr = append(newArr, right[j])
			j++
		}
	}

	for k := i; k < len(left); k++ {
		newArr = append(newArr, left[k])
	}

	for k := j; k < len(right); k++ {
		newArr = append(newArr, right[k])
	}

	return newArr
}
