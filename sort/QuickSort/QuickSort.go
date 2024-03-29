package QuickSort

func QuickSort(a []int) []int {
	sort(a)
	return a
}

func sort(a []int) {
	if len(a) <= 1 {
		return
	}
	mid := a[0]
	head, tail := 0, len(a)-1
	for i := 1; i <= tail; {
		if a[i] > mid {
			a[i], a[tail] = a[tail], a[i]
			tail--
		} else {
			a[i], a[head] = a[head], a[i]
			head++
			i++
		}
	}
	sort(a[:head])
	sort(a[head+1:])
}
