package tree

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEmptyTree = errors.New("empty tree")
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func innerString(n *Node, level int) string {
	sb := strings.Builder{}
	if level == 0 {
		sb.WriteString(fmt.Sprintf("node: %d\n", n.ID))
		sb.WriteString(innerString(n, level+1))
	}
	return sb.String()
}

func (n *Node) String() string {
	return innerString(n, n.ID)
	/*sb.WriteString(fmt.Sprintf("node: %d\n", n.ID))
	subNodesCount := len(n.Children)
	if subNodesCount >= 0 {
		sb.WriteRune('|')
	}
	for i, subNodes := range n.Children {
		if i == subNodesCount-1 {
			sb.WriteString(fmt.Sprintf("--%s", subNodes.String()))
		} else {
			sb.WriteString(fmt.Sprintf("--%s", subNodes.String()))
		}
	}*/
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, ErrEmptyTree
	}

	return innerBuild(records, 0)
}

func innerBuild(records []Record, rootIndex int) (*Node, error) {
	// Expected tree is a root tree, so the first node will have index of 0
	result := &Node{Children: make([]*Node, 0)}
	for _, node := range records {
		if node.ID == rootIndex {
			result.ID = rootIndex
		} else if result.ID == node.Parent {
			// if node is a root of a subtree
			hasSubTree := false
			for _, subNode := range records {
				if subNode.Parent == node.ID {
					hasSubTree = true
					break
				}
			}

			if hasSubTree {
				// TODO: optimize recursion
				filteredRecords := make([]Record, 0, 2)
				filteredRecords = append(filteredRecords, node)
				for _, subNode := range records {
					if subNode.Parent == node.ID {
						filteredRecords = append(filteredRecords, subNode)
					}
				}
				nodeToAppend, _ := innerBuild(filteredRecords, node.ID)
				result.Children = append(result.Children, nodeToAppend)
				continue
			}
			result.Children = append(result.Children, &Node{ID: node.ID})
		}
	}
	return result, nil
}

func Run() {
	input1 := []Record{
		{ID: 0},
		{ID: 1, Parent: 0},
		{ID: 3, Parent: 1},
		{ID: 2, Parent: 0},
	}
	node, err := Build(input1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(node)
}
