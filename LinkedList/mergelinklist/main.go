package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	l2 := &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 7, Next: nil}}}
	l3 := MergeTwoList(l1, l2)
	for next := l3; next != nil; next = next.Next {
		fmt.Printf("%v ", next.Val)
	}
}

func MergeTwoList(l1, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	head := l3
	//新链表节点赋值（由l1和l2最小的节点构成）
	for l1 != nil && l2 != nil { //某条链表遍历完就退出循环
		if l1.Val <= l2.Val {
			l3.Next = l1
			l1 = l1.Next
		} else {
			l3.Next = l2
			l2 = l2.Next
		}
		l3 = l3.Next
	}
	//两个链表不一样长，l1或l2剩下的节点直接挂在l3的尾部
	if l1 != nil {
		l3.Next = l1
	}
	if l2 != nil {
		l3.Next = l2
	}
	return head.Next
}
