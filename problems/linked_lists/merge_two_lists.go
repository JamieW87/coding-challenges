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
	if list1 != nil && list2 == nil {
		return list1
	}
	if list2 != nil && list1 == nil {
		return list2
	}
	if list1 == nil && list2 == nil {
		return nil
	}

	//Create new node type to write the two lists into
	newNode := new(ListNode)

	//If list1s value is bigger than list2s value then write list2s value to the new node
	//Call the function again so that it keeps doing the above, until list1s value is smaller
	//then it writes list2s value to the new node. Returns the new node with the ordered values
	if list1.Val >= list2.Val {
		newNode.Val = list2.Val
		list2 = list2.Next
		newNode.Next = MergeTwoLists(list1, list2)
	} else {
		newNode.Val = list1.Val
		list1 = list1.Next
		newNode.Next = MergeTwoLists(list1, list2)
	}

	return newNode

}
