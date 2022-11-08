package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return isValidHelper(root, nil, nil)

}

func isValidHelper(root *TreeNode, min, max interface{}) bool {

	if root == nil {
		return true
	}

	if (min != nil && root.Val <= min.(int)) || (max != nil && root.Val >= max.(int)) {
		return false
	}

	return isValidHelper(root.Left, min, root.Val) && isValidHelper(root.Right, root.Val, max)

}
