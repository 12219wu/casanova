package quickSort

func QuickSort1(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	median := arr[0]
	//fmt.Printf("par: %v\n", median)

	low_part := make([]int, 0, len(arr))
	high_part := make([]int, 0, len(arr))
	middle_part := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			low_part = append(low_part, item)
		case item == median:
			middle_part = append(middle_part, item)
		case item > median:
			high_part = append(high_part, item)
		}
	}

	low_part = QuickSort1(low_part)
	high_part = QuickSort1(high_part)

	low_part = append(low_part, middle_part...)
	low_part = append(low_part, high_part...)

	return low_part
}
