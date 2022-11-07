package examples

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	root *Node
}

//Standard insert fuction to create a new binary tree (Where root is nil)
func (t *BinaryTree) insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &Node{Key: data, Left: nil, Right: nil}
	} else {
		t.insert(data)
	}
	return t
}

//Standard insert function to insert Nodes under a root
func (n *Node) insertNode(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			//If no node exists, create a new node on the right with key of K
			n.Right = &Node{Key: k}
		} else {
			//If node exists, run the code again, until insert
			n.Right.insertNode(k)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.insertNode(k)
		}
	}
}

//func main() {
//	//Tree has to be an address (&)
//	tree := &Node{Key: 100}
//	tree.insert(50)
//}
