package quickSort

import "fmt"

func QuickSort2(arr []int) []int {
	if len(arr)<=1{
		return arr
	}
	head,tail:=0,len(arr)-1
	mid:=arr[0]
	for ;head<tail;{
		for ;mid<arr[tail]&&head<tail;{
			tail--
		}
		arr[head]=arr[tail]
		head++
		for ;mid>arr[head];{
			head++
		}
		arr[]

	}
	arr[head]=mid
	fmt.Println(arr)
	QuickSort2(arr[:head])
	QuickSort2(arr[head+1:])
	return arr
}