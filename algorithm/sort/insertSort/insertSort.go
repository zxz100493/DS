package insertsort

func InsertSort(list []int) {
	n := len(list)
	for i := 1; i <= n-1; i++ {
		deal := list[i]
		j := i - 1
		if deal < list[j] {

			for ; j >= 0 && list[j] > deal; j-- {
				list[j+1] = list[j]
			}
			list[j+1] = deal
		}
	}
}
