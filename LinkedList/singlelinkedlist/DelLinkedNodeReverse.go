package singlelinkedlist

type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

//删除单链表倒数第n个节点(要求复杂度为O1)
//思路：使用两个指针，第一个指针从头移动n步暂停移动，此时将第二个指针则指向头节点
//两个指针变量再次移动至第一个指针为空时，第二个指针指向的节点就是要删除的元素
func DelSingleLinkedListNodeReverse(head *LinkedNode, nReverse int) *LinkedNode {

	var first = head
	var second = head
	var prev *LinkedNode
	for i := 0; i < nReverse; i++ {
		first = first.Next
	}

	for {
		first = first.Next
		prev = second
		second = second.Next
		if first == nil {
			break
		}
	}
	prev.Next = second.Next
	second.Next = nil
	return head
}
