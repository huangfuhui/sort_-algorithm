package sort_algorithms

// 1.比较相邻的两个元素，如果第一个比第二个大，则交换它们的位置
// 2.一轮下来，最大的元素会排在数列的最后方
// 3.对剩下未排序的元素不断重复步骤一和二，最终得出一个有序的数列
func bubble(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}
