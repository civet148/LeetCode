package main

import "fmt"

func main() {
	elems := []int{32, 12, 34, 42, 88, 312, 18, 1, 0, 99}
	Sort(elems)
}

/*
算法：插入排序
原理：
空间复杂度：
时间复杂度：O(n^2)
是否稳定：
公式：
for(i:=0;i<n;i++) {
	for (j:=i+1;j<n;j++) {swap(i,j)}
}
*/

func Sort(elems []int) {
	n := len(elems)
	loop := 0
	fmt.Printf("插入排序：开始 %v \n", elems)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			loop++
			if elems[i] > elems[j] {
				elems[j], elems[i] = elems[i], elems[j]
			}
		}
	}
	fmt.Printf("插入排序：结束 %v 循环 %v\n", elems, loop)
}
