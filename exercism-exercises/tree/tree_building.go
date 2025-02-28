package tree

import (
	"errors"
	"fmt"
	"strings"
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

func (n *Node) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("node: %d\n", n.ID))
	for i, v := range n.Children {
		if i == len(n.Children)-1 {
			sb.WriteString(fmt.Sprintf("|--node: %d", v.ID))
		} else {
			sb.WriteString(fmt.Sprintf("|--node: %d\n", v.ID))
		}
	}
	return sb.String()
}

var (
	ErrEmptyTree = errors.New("empty tree")
)

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, ErrEmptyTree
	}

	// Expected tree is a root tree, so the first node will have index of 0
	result := &Node{Children: make([]*Node, 0)}
	for i, v := range records {
		isRoot := false
		for _, j := range records {
			if j.Parent == i {
				isRoot = true
				break
			}
		}

		if v.ID == 0 && isRoot {
			result.ID = v.ID
		} else if result.ID == v.Parent {
			//test, err := Build(records)
			//fmt.Println(test, err)
			result.Children = append(result.Children, &Node{ID: v.ID})
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
