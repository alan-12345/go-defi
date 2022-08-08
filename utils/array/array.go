package array

func IndexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}