package linkedlist

// Define List and Node types here.
type Node struct {
	Val        interface{}
	prev, next *Node
}

type List struct {
	first, last *Node
}

func NewList(args ...interface{}) *List {
	l := new(List)
	var pTemp *Node
	for i, arg := range args {

	}

	return nil
}

/*
node addNode(node head, int value) {
    node temp, p; // declare two nodes temp and p
    temp = createNode(); // assume createNode creates a new node with data = 0 and next pointing to NULL.
    temp->data = value; // add element's value to data part of node
    if (head == NULL) {
        head = temp;     // when linked list is empty
    }
    else {
        p = head; // assign head to p
        while (p->next != NULL) {
            p = p->next; // traverse the list until p is the last node. The last node always points to NULL.
        }
        p->next = temp; // Point the previous last node to the new node created.
    }
    return head;
}
*/

func NewNode(val interface{}, p, n *Node) *Node {
	node := Node{
		Val:  val,
		prev: p,
		next: n,
	}
	return &node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) PushFront(v interface{}) {
	panic("Please implement the PushFront function")
}

func (l *List) PushBack(v interface{}) {
	panic("Please implement the PushBack function")
}

func (l *List) PopFront() (interface{}, error) {
	panic("Please implement the PopFront function")
}

func (l *List) PopBack() (interface{}, error) {
	panic("Please implement the PopBack function")
}

func (l *List) Reverse() {
	panic("Please implement the Reverse function")
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}
