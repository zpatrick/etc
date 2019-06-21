package node

import "fmt"

type Scanner struct {
	index int
	node  *Node
}

func NewScanner(n *Node) *Scanner {
	return &Scanner{node: n}
}

func (s *Scanner) ScanToByte(b byte) bool {
	return s.TraverseFunc(func(bt byte) bool {
		if b == bt {
			return true
		}

		return false
	})
}

func (s *Scanner) Next() bool {
	return s.Move(1)
}

func (s *Scanner) Move(count int) bool {
	var moved int
	return s.TraverseFunc(func(b byte) bool {
		moved++
		return moved < count
	})
}

func (s *Scanner) Peek(length int) (string, bool) {
	currentNode := s.node
	currentIndex := s.index
	defer func() {
		s.node = currentNode
		s.index = currentIndex
	}()

	var res string
	completed := s.TraverseFunc(func(b byte) bool {
		res += string(b)
		if len(res) == length {
			return true
		}

		return false
	})

	return res, completed
}

func (s *Scanner) TraverseFunc(fn func(b byte) bool) bool {
	for s.node != nil {
		if s.index >= len(s.node.Data) {
			s.node = s.node.Next
			s.index = 0
			continue
		}

		fmt.Printf("Node: %v\n", s.node.Data)
		for ; s.index < len(s.node.Data); s.index++ {
			fmt.Printf("node[%v]: %s\n", s.index, string(s.node.Data[s.index]))
			if fn(s.node.Data[s.index]) {
				return true
			}
		}
	}

	return false
}
