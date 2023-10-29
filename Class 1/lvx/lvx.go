package lvx

import (
	"math/rand"
)

var Num = rand.Intn(100)

func Find() int {

	//var num int
	//num = rand.Intn(100)

	var left int
	var right int
	var mid int

	left = 0
	right = 99
	mid = (left + right) / 2

	for left <= right {
		if Num == mid {
			return mid
		} else if Num > mid {
			left = left + 1
		} else if Num < mid {
			right = right - 1
		}
		mid = (left + right) / 2
	}
	return -1
}
