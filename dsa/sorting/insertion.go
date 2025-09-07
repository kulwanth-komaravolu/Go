package sorting

import (
	"fmt"
)

func (s *IntSlice) InsertionSort() {
	arr := *s
	if len(arr) <= 1 {
		fmt.Println("Already sorted.")
		return
	}
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] < key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
