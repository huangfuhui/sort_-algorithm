package main

// 1.把第一个元素看成有序的数列
// 2.把未排序数列中的一个元素和有序数列中的元素从右到左做对比，每一个比它大的元素都往右挪一位，
//   直到找到左边比它大，右边比它小的位置后，插入该元素
// 3.重复步骤二，最终得到一个有序的数列
func insertion(arr []int) []int {
	length := len(arr)

	for i := 1; i < length; i++ {
		current := arr[i]
		j := i - 1
		for j >= 0 && arr[j] < current {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = current
	}

	return arr
}
