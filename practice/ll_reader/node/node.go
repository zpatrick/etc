package node

import "io"

type Node struct {
	Data []byte
	Next *Node
}

func NewNode(data []byte, next *Node) *Node {
	return &Node{
		Data: data,
		Next: next,
	}
}

func NewNodeString(data string, next *Node) *Node {
	return NewNode([]byte(data), next)
}

func NewList(datum ...[]byte) *Node {
	if len(datum) == 0 {
		return nil
	}

	root := NewNode(datum[0], nil)
	for i, current := 1, root; i < len(datum); i, current = i+1, current.Next {
		current.Next = NewNode(datum[i], nil)
	}

	return root
}

func NewListString(datum ...string) *Node {
	d := make([][]byte, len(datum))
	for i, s := range datum {
		d[i] = []byte(s)
	}

	return NewList(d...)
}

func (n *Node) Read(p []byte) (int, error) {
	if n == nil || len(n.Data) == 0 {
		return 0, io.EOF
	}

	for i := 0; i < len(p); i++ {
		if i >= len(n.Data) {
			n = n.Next
			return i, nil
		}

		p[i] =
	}

	if len(p) < len(n.Data) {
		
	}


	// if p >= len(n.Data), move to next node
	// toerhwise, up to n.Data	

	for i := 0; i < len(p); i++ {
			
	}

	if len(p) < len(n.Data) {
		for i := 0; i < len(p); i++ {
			p[i] = n.Data[i]	
		}

		n.Data = n.Data[:len(p)]	
		return len(p), nil
	}
		
	return 0, nil
}
