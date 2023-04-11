package selectionsort

func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		min := arr[i]
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < min {
				min = arr[j]
				minIndex = j
			}
		}
		if i != minIndex {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}
