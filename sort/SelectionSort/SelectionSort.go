package SelectionSort

func SelectionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		min := a[i]
		minkey := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < min {
				min = a[j]
				minkey = j
			}
		}
		a[i],a[minkey]=a[minkey],a[i]
	}
	return a
}
