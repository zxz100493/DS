package bucket

import "fmt"

func BucketSort(arr []int, maxVal int) []int {
	buckets := make([]int, maxVal+1)

	for _, v := range arr {
		fmt.Println("v:", v)
		buckets[v]++
	}
	fmt.Println("buckets:", buckets)
	idx := 0
	for i, bucket := range buckets {
		fmt.Println("i:", i, "bucket:", bucket)
		for j := 0; j < bucket; j++ {
			fmt.Println("j:", j, "idx:", idx)
			arr[idx] = i
			idx++
		}
	}
	return arr
}
