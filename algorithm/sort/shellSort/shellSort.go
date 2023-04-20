package shellsort

import "fmt"

// 增量序列折半的希尔排序
func ShellSort(list []int) {
	// 数组长度
	n := len(list)

	// 每次减半，直到步长为 1
	for step := n / 2; step >= 1; step /= 2 {
		// 开始插入排序，每一轮的步长为 step
		fmt.Printf("step------:%d\n", step)
		for i := step; i < n; i += step {
			fmt.Printf("i:%d\n", i)
			for j := i - step; j >= 0; j -= step {
				fmt.Printf("j--:%d\n", j)
				// 满足插入那么交换元素
				if list[j+step] < list[j] {
					list[j], list[j+step] = list[j+step], list[j]
					continue
				}
				break
			}
		}
	}
}

/* func ShellSort(arr []int) {
	n := len(arr) // 获取数组长度
	gap := n / 2  // 定义初始间隔为数组长度的一半
	for gap > 0 { // 当间隔不为0时，继续进行排序
		for i := gap; i < n; i++ { // 对每个间隔进行插入排序
			temp := arr[i]                      // 存储当前要插入的元素
			j := i                              // 定义插入的起始位置
			for j >= gap && arr[j-gap] > temp { // 向前比较并移动元素
				arr[j] = arr[j-gap]
				j -= gap // 更新插入位置
			}
			arr[j] = temp // 插入当前元素
		}
		gap /= 2 // 更新间隔
	}
} */
