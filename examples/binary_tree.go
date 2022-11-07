package examples

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) insert(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			//If no node exists, create a new node on the right with key of K
			n.Right = &Node{Key: k}
		} else {
			//If node exists, run the code again, until insert
			n.Right.insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.insert(k)
		}
	}
}

//func main() {
//	//Tree has to be an address (&)
//	tree := &Node{Key: 100}
//	tree.insert(50)
//}
