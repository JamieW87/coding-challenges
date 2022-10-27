package problems

//You are given the heads of two sorted linked lists list1 and list2.
//
//Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
//
//Return the head of the merged linked list.
//
//Example 1:
//
//Input: list1 = [1,2,4], list2 = [1,3,4]
//Output: [1,1,2,3,4,4]

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	ln := &ListNode{
		Val:  0,
		Next: nil,
	}

	return ln

}

//Solution:
//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil && l2 != nil { return l2 }
//	if l1 != nil && l2 == nil { return l1 }
//	if l1 == nil && l2 == nil { return nil }
//	newNode := new(ListNode)
//	if l1.Val >= l2.Val {
//		newNode.Val = l2.Val
//		l2 = l2.Next
//		newNode.Next = mergeTwoLists(l1, l2)
//	} else {
//		newNode.Val = l1.Val
//		l1 = l1.Next
//		newNode.Next = mergeTwoLists(l1, l2)
//	}
//	return newNode
//}
