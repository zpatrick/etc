package main

import (
	"fmt"
	"log"
)

type Node struct {
	ID          int
	Left, Right *Node
}

func (n *Node) DFT() {
	fmt.Println(n.ID)
	if n.Left != nil {
		n.Left.DFT()
	}
	if n.Right != nil {
		n.Right.DFT()
	}
}

func (n *Node) BFT() {
	parents := []*Node{n}
	fmt.Printf("Level 0: %v\n", n.ID)
	for level := 1; len(parents) > 0; level++ {
		var children []*Node
		for _, p := range parents {
			for _, c := range []*Node{p.Left, p.Right} {
				if c != nil {
					children = append(children, c)
				}
			}
		}

		fmt.Printf("Level %d: ", level)
		for _, c := range children {
			fmt.Printf("%v ", c.ID)
		}
		fmt.Println()

		parents = children
	}
}

func (n *Node) ShortestPath(to int) ([]int, bool) {
	path := []int{n.ID}
	if n.ID == to {
		return path, true
	}

	for _, node := range []*Node{n.Left, n.Right} {
		if node != nil {
			p, ok := node.ShortestPath(to)
			if ok {
				return append(path, p...), true
			}
		}
	}

	return nil, false
}

func (n *Node) LCA(n1, n2 int) (int, error) {
	path1, ok := n.ShortestPath(n1)
	if !ok {
		return 0, fmt.Errorf("Node %d doesn't exist", n1)
	}

	path2, ok := n.ShortestPath(n2)
	if !ok {
		return 0, fmt.Errorf("Node %d doesn't exist", n2)
	}

	var lca int
	for i := 0; i < len(path1) && i < len(path2); i++ {
		if path1[i] == path2[i] {
			lca = path1[i]
			continue
		}

		break
	}

	return lca, nil
}

func IsBST(n *Node) bool {
	if n == nil {
		return true
	}

	if n.Left != nil && n.Left.ID > n.ID {
		return false
	}

	if n.Right != nil && n.Right.ID < n.ID {
		return false
	}

	return IsBST(n.Left) && IsBST(n.Right)
}

type NodeDistance struct {
	Node               *Node
	DistanceFromCenter int
}

func PrintTop(n *Node) {
	// always print the root
	fmt.Println(n.ID)
	farthestDistancePrintedLeft := 0
	farthestDistancePrintedRight := 0

	parents := []NodeDistance{{Node: n, DistanceFromCenter: 0}}
	for len(parents) > 0 {
		// gather children at current level
		children := []NodeDistance{}
		for _, parent := range parents {
			if parent.Node.Left != nil {
				child := NodeDistance{
					Node:               parent.Node.Left,
					DistanceFromCenter: parent.DistanceFromCenter - 1,
				}

				children = append(children, child)
			}

			if parent.Node.Right != nil {
				child := NodeDistance{
					Node:               parent.Node.Right,
					DistanceFromCenter: parent.DistanceFromCenter + 1,
				}

				children = append(children, child)
			}
		}

		// determine which children to print
		for _, child := range children {
			if child.DistanceFromCenter < farthestDistancePrintedLeft {
				fmt.Println(child.Node.ID)
				farthestDistancePrintedLeft = child.DistanceFromCenter
			}

			if child.DistanceFromCenter > farthestDistancePrintedRight {
				fmt.Println(child.Node.ID)
				farthestDistancePrintedRight = child.DistanceFromCenter
			}
		}

		// move to the next level
		parents = children
	}
}

func TopView(n *Node) {

}

func main() {
	root := &Node{
		ID: 1,
		Left: &Node{
			ID:    2,
			Left:  &Node{ID: 4},
			Right: &Node{ID: 5},
		},
		Right: &Node{
			ID:    3,
			Left:  &Node{ID: 6},
			Right: &Node{ID: 7},
		},
	}

	lca, err := root.LCA(2, 5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("LCA:", lca)

	fmt.Println("View from the top: ")
	root = &Node{
		ID: 0,
		Right: &Node{
			ID: 1,
			Left: &Node{
				ID: 0,
				Left: &Node{
					ID: -1,
					Left: &Node{
						ID: -2,
					},
				},
			},
		},
	}
	PrintTop(root)

	root = &Node{
		ID: 2,
		Left: &Node{
			ID:   1,
			Left: &Node{ID: 5},
		},
		Right: &Node{
			ID: 3,
		},
	}

	fmt.Println("Is BST:", IsBST(root))
	fmt.Println("DFT:")
	root.DFT()
	fmt.Println("BFT:")
	root.BFT()
}
