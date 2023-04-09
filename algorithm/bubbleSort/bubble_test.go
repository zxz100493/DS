package bubblesort

import (
	"fmt"
	"testing"
)

// bubble sort test
func TestBubbleSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	fmt.Println("before", arr)
	BubbleSort(arr)
	fmt.Println("after", arr)
}
