package examples

import "fmt"

type node struct {
	val  int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

//Prepends a new node to the list
func (l *linkedList) prepend(n *node) {
	//Create a temporary node, second, then create the new node, n, as the head
	second := l.head
	l.head = n
	//Point new head to the old head
	l.head.next = second
	//Have added something so need to increase length
	l.length++
}

func (l *linkedList) deleteWithValue(value int) {
	previousToDelete := l.head
	//Loops until it finds a match, sets previousToDelete to the next node
	for previousToDelete.next.val != value {
		previousToDelete = previousToDelete.next
	}
	//This sets the node that needs to be deleted (previousToDelete.next) and sets its to the next value after that, removing it from the list.
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

//Prints the current list
func (l *linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%v", toPrint.val)
		toPrint = toPrint.next
		l.length--
		fmt.Println("")
	}
}

func AddNode() {
	//Create new linkedList create some new nodes instances
	myList := linkedList{}
	node1 := &node{val: 8}
	node2 := &node{val: 22}
	node3 := &node{val: 48}
	node4 := &node{val: 20}

	//Prepend them to myList
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)

	//myList.deleteWithValue(48)

	myList.printListData()
}
