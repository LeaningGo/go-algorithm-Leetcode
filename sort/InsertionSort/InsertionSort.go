package InsertionSort

func InsertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		for j := i - 1; j >= 0; j-- {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
				i--
			} else {
				break
			}
		}
	}
	return a
}
