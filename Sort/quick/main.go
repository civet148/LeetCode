package main

import "fmt"

func main() {
	elems := []int{32, 12, 34, 42, 88, 312, 12, 1, 0, 99}
	Sort(elems)
}

/*
算法：快速排序
原理：
空间复杂度：
时间复杂度：O(logN)
是否稳定：
*/

func Sort(elems []int) {
	fmt.Printf("快速排序：开始 %v \n", elems)

	fmt.Printf("快速排序：结束 %v \n", elems)
}
