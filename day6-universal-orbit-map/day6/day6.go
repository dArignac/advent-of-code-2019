package day6

import (
	"errors"
	"strings"
)

// Node represents a space object
type Node struct {
	name     string
	children []*Node
	parent   *Node
}

func BulkInsertIntoTree() {

}

func insertIntoTree(root *Node, value string) error {
	// split string value of objects
	nodes := strings.Split(value, ")")

	// search left object in tree, must exist, if not throw
	left := findNodeByName(root, nodes[0])
	if left == nil {
		return errors.New("left node not already in tree")
	}

	// search right object in tree, if not exist, create it
	right := findNodeByName(root, nodes[1])
	if right == nil {
		right = &Node{name: nodes[1], parent: left, children: nil}
		left.children = append(left.children, right)
	}

	return nil
}

func findNodeByName(root *Node, name string) *Node {
	if root.name == name {
		return root
	}

	for _, child := range root.children {
		if child.name == name {
			return child
		}

		if len(child.children) > 0 {
			result := findNodeByName(child, name)
			if result != nil {
				return result
			}
		}
	}

	return nil
}
