package main

// 1.将整个待排序数列(R1, R2, R3, ..., Rn)按增量d划分成d个子数列，其中第i个子数列为(Ri, Ri+d, Ri+2d, ..., Ri+kd)
// 2.分别对各个子数列进行一次插入排序
//
// 增量d的取值数列称之为增量数列，基于增量数列的降序特点，希尔排序也被称为缩小增量排序
// 按照增量数列d[0...t-1]对待排序数列arr作t趟希尔排序
func shell(arr []int, d []int, t int) []int {
	for k := 0; k < t; k++ {
		shellInsert(&arr, d[k])
	}

	return arr
}

// 对排序数列作一趟增量为dk的希尔排序
func shellInsert(arr *[]int, dk int) {
	length := len(*arr)

	// 进行插入排序
	for i := 0; i < length-dk; i++ {
		if (*arr)[i+dk] < (*arr)[i] {
			tmp := (*arr)[i+dk]

			j := i + dk
			for j-dk >= 0 && tmp < (*arr)[j-dk] {
				j -= dk
				(*arr)[j+dk] = (*arr)[j]
			}

			(*arr)[j] = tmp
		}
	}
}
