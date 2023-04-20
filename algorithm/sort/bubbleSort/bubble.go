package bubblesort

import "fmt"

func BubbleSort(arr []int) {
	num := len(arr)
	for i := 0; i < num; i++ {
		fmt.Println()
		fmt.Printf("i:%d", i)
		fmt.Println()

		for j := 0; j < num-i-1; j++ {
			fmt.Printf("j:%d--j+1:%d", arr[j], arr[j+1])
			fmt.Println()
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
