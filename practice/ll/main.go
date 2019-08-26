package main

import (
	"fmt"
	"strconv"
)

type LinkedList struct {
	head *Node
}

func NewLinkedList(data int) *LinkedList {
	return &LinkedList{
		head: &Node{Data: data},
	}
}

func (l *LinkedList) Insert(data int) {
	current := l.head
	for ; current.Next != nil; current = current.Next {
	}

	current.Next = &Node{Data: data}
}

func (l *LinkedList) Traverse(fn func(i int, n *Node) bool) {
	for i, current := 0, l.head; current != nil; i, current = i+1, current.Next {
		if !fn(i, current) {
			break
		}
	}
}

func (l *LinkedList) Reverse() {
	ptr := l.head
	var newHead *Node
	for ptr != nil {
		current := ptr
		ptr = ptr.Next

		current.Next = newHead
		newHead = current
	}

	l.head = newHead
}

func (l *LinkedList) String() string {
	var s string
	l.Traverse(func(i int, n *Node) bool {
		s += strconv.Itoa(n.Data)
		if n.Next != nil {
			s += "->"
		}

		return true
	})

	return s
}

/*

Swap 2 nodes:

0 -> 1 -> 2 -> 3 -> 4 -> 5

Swap(1, 4)

Iterate to 0, since 0.Next == 1
Iterate to 3, since 3.Next == 4

a := 0.Next
b := 3.Next

0.Next = b
tmp := a.Next
a.Next = b.Next
b.Next = tmp
3.Next = a

tmp = 2
0 -> 4
1 -> 5
4 -> 2

0 -> 4 -> 2 -> 3 -> 1 -> 5
*/

type Node struct {
	Data int
	Next *Node
}

func FindInsersectStart(l1, l2 *LinkedList) *Node {
	nodes := map[int]*Node{}
	l1.Traverse(func(i int, n *Node) bool {
		nodes[n.Data] = n
		return true
	})

	var start *Node
	l2.Traverse(func(i int, n *Node) bool {
		if _, ok := nodes[n.Data]; ok {
			start = n
			return false
		}

		return true
	})

	return start
}

// how would we find all intersections?
/*
	- create a from l1 map[int]{n *Node, index int, visited bool}
	- as you iterate through l2, find

how about w/o a map?
if you use two pointers

*/

// How do you reverse a singly linked list
func Reverse(n *Node) *Node {
	var previous, current *Node
	var head = n
	var pointer = head

	for pointer != nil {
		current = pointer
		pointer = pointer.Next

		current.Next = previous
		previous = current
		head = current
	}

	return head
}

func main() {
	l1 := NewLinkedList(5)
	l1.Insert(8)
	l1.Insert(3)
	l1.Insert(2)
	l1.Insert(9)

	l2 := NewLinkedList(1)
	l2.Insert(0)
	//l2.Insert(3)
	//l2.Insert(2)
	//l2.Insert(9)

	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(FindInsersectStart(l1, l2))

	l3 := NewLinkedList(1)
	l3.Insert(2)
	l3.Insert(3)
	l3.Insert(4)
	fmt.Println("Before: ", l3)
	l3.Reverse()
	fmt.Println("After: ", l3)
}
