package main

import "fmt"

func main() {
	elems := []int{32, 12, 34, 42, 88, 312, 18, 1, 0, 99}
	Sort(elems)
}

/*
算法：冒泡排序
原理：经历n次遍历把每次左边大于右边的元素进行交换（每轮第n-i个元素被确定位置）
空间复杂度：
时间复杂度：O(n^2)
是否稳定：
公式：
for(i:=0;i<n-1;i++){
  for(j:=0;j<n-i-1;j++) {swap(j,j+1}
}
*/

func Sort(elems []int) {
	n := len(elems)
	loop := 0
	fmt.Printf("冒泡排序：开始 %v \n", elems)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			loop++
			if elems[j] > elems[j+1] {
				elems[j], elems[j+1] = elems[j+1], elems[j]
			}
		}
	}
	fmt.Printf("冒泡排序：结束 %v 循环 %v\n", elems, loop)
}
