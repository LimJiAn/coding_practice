package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{3, &ListNode{}}
	list2 := &ListNode{2, &ListNode{}}
	mergeTwoLists(list1, list2)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var newList = &ListNode{}
	var out = newList
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			newList.Next = list1
			list1 = list1.Next
			newList = newList.Next
		} else {
			newList.Next = list2
			list2 = list2.Next
			newList = newList.Next
		}
	}
	if list1 != nil {
		newList.Next = list1
	} else if list2 != nil {
		newList.Next = list2
	}
	return out.Next
}
