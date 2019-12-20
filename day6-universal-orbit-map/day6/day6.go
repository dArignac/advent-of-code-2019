package day6

import (
	"errors"
	"strings"
)

// Node represents a space object
type Node struct {
	Name     string
	Children []*Node
	Parent   *Node
}

// BulkCreateTree creates a tree from the given, unsorted values representing object relations
func BulkCreateTree(values []string) (root Node) {
	// FIXME need to find COM)XXX
	// and start from there,
	// then search for XXX)YYY and continue

	searchForLeftObject := func(list []string, name string) []string {
		for _, value := range list {
			values := strings.Split(value, ")")
			if values[0] == name {
				return values
			}
		}
		return []string{}
	}

	root = Node{Name: "COM", Parent: nil, Children: nil}

	// find COM as starting point
	startValues := searchForLeftObject(values, "COM")

	// FIXME continue

	return
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
		right = &Node{Name: nodes[1], Parent: left, Children: nil}
		left.Children = append(left.Children, right)
	}

	return nil
}

func findNodeByName(root *Node, name string) *Node {
	if root.Name == name {
		return root
	}

	for _, child := range root.Children {
		if child.Name == name {
			return child
		}

		if len(child.Children) > 0 {
			result := findNodeByName(child, name)
			if result != nil {
				return result
			}
		}
	}

	return nil
}
