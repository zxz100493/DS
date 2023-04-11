package selectionsort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	list := []int{1, 6, 23, 5, 7, 2, 43}
	SelectionSort(list)
	fmt.Println(list)
}
