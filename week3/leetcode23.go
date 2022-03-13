package main


type ListNode struct {
	Val int
	Next *ListNode
}

func merge2list(l1, l2*ListNode)*ListNode{
	if l1 == nil {
		return l2
	}else if l2 == nil {
		return l1
	}
	protectNode := new(ListNode)
	if l1.Val < l2.Val {
		protectNode.Next = l1
	}else  {
		protectNode.Next = l2
	}
	headNode := protectNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			headNode.Next = l1
			l1 = l1.Next
		}else {
			headNode.Next = l2
			l2 = l2.Next
		}
		headNode = headNode.Next
	}
	if l1 != nil {
		headNode.Next = l1
	}
	if l2 != nil {
		headNode.Next = l2
	}
	return protectNode.Next
}

func merge(lists []*ListNode, left, right int) *ListNode{
	if left == right {
		return lists[left]
	}
	if left > right {
		return nil
	}
	mid := (left + right) >> 1
	return merge2list(merge(lists, left, mid), merge(lists, mid + 1, right))
}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists) - 1)
}

func main() {
	node1 := new(ListNode)
	node1.Val = 1
	node2 := new(ListNode)
	node2.Val = 4
	node3 := new(ListNode)
	node3.Val = 5
	node1.Next = node2
	node2.Next = node3

	node4 := new(ListNode)
	node4.Val = 1
	node5 := new(ListNode)
	node5.Val = 3
	node6 := new(ListNode)
	node6.Val = 4
	node4.Next = node5
	node5.Next = node6

	node7 := new(ListNode)
	node7.Val = 2
	node8 := new(ListNode)
	node8.Val = 6
	node7.Next = node8

	l := []*ListNode{node1, node4, node8}
	mergeKLists(l)
}