package sorting

func (s *IntSlice) MergeSort(low, high int) {
	if low >= high {
		return
	}
	mid := low + (high-low)/2
	s.MergeSort(low, mid)
	s.MergeSort(mid+1, high)
	s.merge(low, mid, high)
}

func (s *IntSlice) merge(low, mid, high int) {
	arr := *s
	n1 := mid - low + 1
	n2 := high - mid

	lArr := make([]int, n1)
	rArr := make([]int, n2)
	for i := 0; i < n1; i++ {
		lArr[i] = arr[low+i]
	}
	for i := 0; i < n2; i++ {
		rArr[i] = arr[mid+1+i]
	}

	temp := make([]int, 0, n1+n2)
	i, j := 0, 0
	for i < n1 && j < n2 {
		if lArr[i] >= rArr[j] {
			temp = append(temp, lArr[i])
			i++
		} else {
			temp = append(temp, rArr[j])
			j++
		}
	}

	if i < n1 {
		temp = append(temp, lArr[i:]...)
	}
	if j < n2 {
		temp = append(temp, rArr[j:]...)
	}

	for i := 0; i < len(temp); i++ {
		arr[low+i] = temp[i]
	}
	// copy(arr[low:high+1], temp)
}
