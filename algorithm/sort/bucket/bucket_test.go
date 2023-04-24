package bucket

import (
	"fmt"
	"testing"
)

func TestBucketSort(t *testing.T) {
	arr := []int{5, 3, 2, 8, 6, 4}
	maxVal := 8

	sortedArr := BucketSort(arr, maxVal)
	fmt.Println(sortedArr) // 输出 [2 3 4 5 6 8]
}
