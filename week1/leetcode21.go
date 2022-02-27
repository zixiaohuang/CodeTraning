package main
type ListNode struct {
     Val int
     Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var protectNode, headNode *ListNode
	protectNode = new(ListNode)
	if list1.Val <= list2.Val {
		protectNode.Next = list1
	}else {
		protectNode.Next = list2
	}

	headNode = protectNode
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			headNode.Next = list1
			list1 = list1.Next
		} else {
			headNode.Next = list2
			list2 = list2.Next
		}
		headNode = headNode.Next
	}

	if list1 != nil {
		headNode.Next = list1
	}else if list2 != nil {
		headNode.Next = list2
	}
	return protectNode.Next
}