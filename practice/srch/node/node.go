package node

type Node struct {
	Data string
	Next *Node
}

func NewNode(data string) *Node {
	return &Node{Data: data}
}

func NewNodes(data ...string) *Node {
	if len(data) == 0 {
		return nil
	}

	head := NewNode(data[0])
	for current, i := head, 1; i < len(data); i++ {
		current.Next = NewNode(data[i])
		current = current.Next
	}

	return head
}

func (n *Node) Contains(query string) bool {
	if len(query) == 0 {
		return true
	}

	scanner := NewScanner(n)
	for scanner.ScanToByte(query[0]) {
		peek, ok := scanner.Peek(len(query))
		if !ok {
			return false
		}

		if peek == query {
			return true
		}
		
		// move over 1 so we don't perpetually ScanToByte
		scanner.Next()
	}

	return false
}
