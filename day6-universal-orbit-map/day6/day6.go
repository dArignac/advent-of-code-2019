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

// direct orbit = parent
// indirect orbit = the parent's parent count up to root/COM
func (node Node) getOrbitCount() int {
	if node.Parent == nil {
		return 0
	}

	if (*node.Parent).Name == "COM" {
		return 1
	}

	return 1 + (*node.Parent).getOrbitCount()
}

// RecursiveCountOrbits counts the orbits for all nodes down or equal to the given node
func RecursiveCountOrbits(node *Node, counter *int) {
	*counter += (*node).getOrbitCount()
	for _, child := range node.Children {
		RecursiveCountOrbits(child, counter)
	}
}

// CreateTree creates a tree from the given, unsorted values representing object relations
func CreateTree(pairs []string) (root Node) {
	root = Node{Name: "COM", Parent: nil, Children: nil}
	recursiveInsertToTree(&root, pairs)
	return
}

func recursiveInsertToTree(root *Node, pairs []string) {
	failedOnes := []string(nil)
	for _, pair := range pairs {
		err := insertIntoTree(root, pair)
		if err != nil {
			failedOnes = append(failedOnes, pair)
		}
	}
	if len(failedOnes) > 0 {
		recursiveInsertToTree(root, failedOnes)
	}
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
