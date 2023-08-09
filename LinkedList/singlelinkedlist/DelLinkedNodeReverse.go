package main

import "fmt"

type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

func main() {
	head := &LinkedNode{}
	list := head
	for i := 0; i < 10; i++ {
		if i == 0 {
			list.Val = i
		} else {
			list.Next = &LinkedNode{
				Val: i,
			}
			list = list.Next
		}
	}
	printLinkedList(head)
	head = DelSingleLinkedListNodeReverse(head, 10)
	printLinkedList(head)
}

func printLinkedList(list *LinkedNode) {
	for node := list; node != nil; node = node.Next {
		fmt.Printf("%d ", node.Val)
	}
	fmt.Println()
	fmt.Println("----------------------------------------------------------------")
}

// 删除单链表倒数第n个节点(要求复杂度为O1)
// 思路：使用两个指针，第一个指针从头移动n步暂停移动，此时将第二个指针则指向头节点
// 两个指针变量再次移动至第一个指针为空时，第二个指针指向的节点就是要删除的元素
func DelSingleLinkedListNodeReverse(head *LinkedNode, n int) *LinkedNode {

	var p1 = head
	var p2 = head
	var prev *LinkedNode
	//第一个指针移动n个位置
	for i := 0; i < n; i++ {
		if p1.Next == nil {
			break
		}
		p1 = p1.Next
	}

	//两个指针同时移动直到p1指针为空，此时p2指向要删除的节点
	for {
		if p1.Next == nil { //p1已是最后一个节点
			break
		}
		prev = p2
		p1 = p1.Next
		p2 = p2.Next
		if p1 == nil {
			break
		}
	}
	if prev == nil {
		return head.Next
	}
	prev.Next = p2.Next
	p2.Next = nil
	return head
}
