package sort_algorithms

// 1.首先在一轮比较中选出最小的元素，然后将它和该轮数列的第一个元素交换
// 2.对未排序的剩下元素不断重复步骤一，最终得到一个有序数列
func selection(arr []int) []int {
	length := len(arr)

	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}
