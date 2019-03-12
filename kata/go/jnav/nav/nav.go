package nav

import (
	"fmt"
	"strings"
)

type Navigator struct {
	Root *Tree
	Node *Tree
}

func NewNavigator() *Navigator {
	return &Navigator{}
}

func (n *Navigator) Load(data []byte) error {
	root, err := LoadJSONTree(data)
	if err != nil {
		return err
	}

	n.Root = root
	n.Node = root
	return nil
}

func (n *Navigator) CurrentNode() *Tree {
	return n.Node
}

func (n *Navigator) ChangeNode(path string) error {
	switch path {
	case ".":
		return nil
	case "..":
		if parent := n.Node.Parent(); parent != nil {
			n.Node = parent
		}

		return nil
	}

	current := n.Node
	if strings.HasPrefix(path, ".") {
		current = n.Root
	}

	for _, step := range strings.Split(path, ".") {
		child := current.Child(step)
		if child == nil {
			return fmt.Errorf("%s: no such node", path)
		}

		current = child
	}

	n.Node = current
	return nil
}

func (n *Navigator) List() []*Tree {
	return n.Node.Children()
}
