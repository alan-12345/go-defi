package array

import "go_defi/utils/constants"

func IndexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func TokenIndexOf(element constants.Token, data []constants.Token) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
