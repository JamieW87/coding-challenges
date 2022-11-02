package examples

type node struct {
	val  int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	//Create a temporary node, second, then create the new node, n, as the head
	second := l.head
	l.head = n
	//Point new head to the old head
	l.head.next = second
	//Have added something so need to increase length
	l.length++
}

func AddNode() {
	//Create new create some new nodes
	myList := linkedList{}
	node1 := &node{val: 8}
}
