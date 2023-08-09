package main

import "fmt"

func main() {
	var v = 4
	var elements = []int{1, 2, 3, 4, 5, 4, 6}
	fmt.Printf("elements %v\n", elements)
	elements = RemoveElements(elements, v)
	fmt.Printf("elements %v\n", elements)
}

// 从切片中原地移除指定值（允许额外占用O1空间）并返回切片长度
func RemoveElements(elements []int, val int) []int {
	for i := 0; i < len(elements); {
		if elements[i] == val {
			elements = append(elements[:i], elements[i+1:]...)
		} else {
			i++
		}
	}
	return elements
}
