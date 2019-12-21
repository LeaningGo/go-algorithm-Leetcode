package BubbleSort

import "fmt"

func BubbleSort(a []int)[]int{
	for i:=1;i<len(a);i++{
		for j:=0;j<len(a)-i;j++{
			if a[j]<a[j+1]{
				a[j],a[j+1]=a[j+1],a[j]
				fmt.Println(a)
			}
		}
	}
	return a
}
