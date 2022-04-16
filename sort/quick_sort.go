package sort

import "fmt"

//来源于WIKI 网站
func QuicWiki(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	data[head] = mid
	QuicWiki(data[:head])
	QuicWiki(data[head+1:])
	return data
}

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
	high_part = append(high_part, middle_part...)
	high_part = append(high_part, low_part...)
	return high_part
	//low_part = QuickSort1(low_part)
	//high_part = QuickSort1(high_part)
	//
	//low_part = append(low_part, middle_part...)
	//low_part = append(low_part, high_part...)
	//
	//return low_part
}

func QuickSort2(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	head, tail := 0, len(arr)-1
	mid := arr[0]
	for head < tail {
		for mid < arr[tail] && head < tail {
			tail--
		}
		if head < tail {
			arr[head] = arr[tail]
			head++
		}

		for mid > arr[head] && head < tail {
			head++
		}
		if head < tail {
			arr[tail] = arr[head]
			tail--
		}
	}
	arr[head] = mid
	fmt.Println(arr)
	QuickSort2(arr[:head])
	QuickSort2(arr[head+1:])
	return arr
}
