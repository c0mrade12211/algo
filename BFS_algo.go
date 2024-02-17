package main

import (
	"fmt"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func NewNode(data int) *Node {
	return &Node{Data: data}
}

func BFS(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}

	fmt.Println(queue)         //[0xc000010018] pointer to slice nodes
	fmt.Println(queue[0])      //&{1 0xc000010030 0xc000010048} 1 - its data, 0xc000010030 - left child, right child - 0xc000010048
	fmt.Println(queue[0].Left) // Next childs &{2 0xc000010060 0xc000010078}

	for len(queue) > 0 {
		node := queue[0]  // take first node
		queue = queue[1:] // remove first node from queue

		fmt.Print(node.Data, " ")

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func main() {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)

	fmt.Println("Обход в ширину (BFS):")
	BFS(root)
	fmt.Println()

}
