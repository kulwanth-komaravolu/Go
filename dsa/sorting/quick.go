package sorting

func (s *IntSlice) QuickSort(low, high int) {
	if low < high {
		piv := s.partition(low, high)

		s.QuickSort(low, piv-1)
		s.QuickSort(piv+1, high)
	}
}

// Using Lomuto's Partiton Scheme.
func (s *IntSlice) partition(low, high int) int {
	arr := *s
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] >= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
