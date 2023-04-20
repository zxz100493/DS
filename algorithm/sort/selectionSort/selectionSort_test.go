package selectionsort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	// list := []int{1, 6, 23, 5, 7, 2, 43}
	// SelectionSort(list)
	// fmt.Println(list)

	list4 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6}
	SelectGoodSort(list4)
	fmt.Println(list4)
}
