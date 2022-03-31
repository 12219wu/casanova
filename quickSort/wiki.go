package quickSort

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
