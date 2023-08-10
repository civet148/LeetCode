package main

import "fmt"

func main() {
	elems := []int{32, 12, 34, 42, 88, 312, 12, 1, 0, 99}
	fmt.Printf("快速排序：开始 %v \n", elems)
	Sort(elems)
	fmt.Printf("快速排序：结束 %v \n", elems)
}

/*
算法：快速排序
原理：
空间复杂度：
时间复杂度：O(logN)
是否稳定：
公式:
qsort(elems) {
	for head<trip {swap(head+1,trip)}
    qsort(elems[:head])
    qsort(elems[head+1:])

}
*/

func Sort(elems []int) {

	if len(elems) < 2 {
		return
	}
	value := elems[0]
	head, trip := 0, len(elems)-1
	for head < trip {
		if elems[head+1] > value { //当前元素大于标尺元素换到右边(trip--)
			elems[head+1], elems[trip] = elems[trip], elems[head+1]
			trip--
		} else if elems[head+1] < value { //当前元素小于标尺元素换到左边(head++)
			elems[head], elems[head+1] = elems[head+1], elems[head]
			head++
		} else { //当前元素等于标尺元素(head++)
			head++
		}
	}
	Sort(elems[:head])
	Sort(elems[head+1:])
}
