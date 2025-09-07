package sorting

import (
	"fmt"
)

type IntSlice []int

func (s *IntSlice) BubbleSort() {
	arr := *s
	if len(arr) <= 1 {
		fmt.Println("Already sorted.")
		return
	}
	for i := 0; i < (len(arr) - 1); i++ {
		var noSwap bool = true
		for j := 0; j < (len(arr) - i - 1); j++ {
			if arr[j] < arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				noSwap = false
			}
		}
		if noSwap {
			break
		}
	}
}
