package quick

import "fmt"

func QuickSort(array []int, begin, end int) {
	fmt.Println("quick sort")
	if begin < end {
		loc := partition(array, begin, end)
		QuickSort(array, begin, loc-1)
		fmt.Println("+++++", array)
		QuickSort(array, loc+1, end)
		fmt.Println("---", array)
	}
}

func partition(array []int, begin, end int) int {
	i := begin + 1
	j := end

	for i < j {
		if array[i] > array[begin] {
			array[i], array[j] = array[j], array[i]
			j--
		} else {
			i++
		}
	}

	if array[i] >= array[begin] {
		i--
	}

	array[begin], array[i] = array[i], array[begin]
	return i
}
