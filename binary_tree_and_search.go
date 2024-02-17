//It's a simple case which demonstrate binary tree
//Also function searchElement which have a O(log n) complexity 

package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func NewNode(data int) *Node {
	return &Node{Data: data}
}
func searchElement(findelement int, root *Node) bool {
	if root == nil {
		return false
	}
	if root.Data == findelement {
		return true
	}
	if root.Data > findelement {
		return searchElement(findelement, root.Left)
	}
	return searchElement(findelement, root.Right)

}
func main() {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)

	var element int
	fmt.Scan(&element)

	if searchElement(element, root) {
		fmt.Println("Элемент содержится в дереве")
	} else {
		fmt.Println("Элемент не содержится в дереве")
	}

	fmt.Println("Пример бинарного дерева:")
	fmt.Println("     ", root.Data)
	fmt.Println("    / \\")
	fmt.Println("   ", root.Left.Data, " ", root.Right.Data)
	fmt.Println("  / \\")
	fmt.Println(root.Left.Left.Data, " ", root.Left.Right.Data)
}
