package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	// 创建链表
	list := &List{}
	// 添加元素
	list.Add(1)
	list.Add(2)
	list.Add(3)

}

// ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}

// 判断链表有环
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
