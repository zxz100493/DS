package quick

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort(list3, 0, len(list3)-1)
	fmt.Println(list3)
}
