package BucketSort

func BucketSort(a []int) []int{
	var b = [10]int{0}
	var c []int
	for _, i := range a {
		b[i] = 1
	}
	for key, i := range b {
		if i == 1 {
			c = append(c, key)
		}
	}
	return c
}
