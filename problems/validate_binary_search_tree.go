package problems

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func IsValidBST(root *Node) bool {

	if root == nil {
		return true
	}

	var isValid bool
	rootLeft := new(Node)
	rootRight := new(Node)

	if root.Left != nil {
		if root.Key < root.Left.Key {
			return false
		} else {
			rootLeft = root.Left
			isValid = IsValidBST(rootLeft)
		}
	}

	if root.Right != nil {
		if root.Key > root.Right.Key {
			return false
		} else {
			rootRight = root.Right
			isValid = IsValidBST(rootRight)
		}
	}

	return isValid
}
