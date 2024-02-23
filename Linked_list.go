// Demonstrate linked list on golang
// Linked list = data, *pointer_to_next_node
// Also realised addHead, addTail
package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func addHead(head *Node, data int) *Node {
	if head == nil {
		return &Node{Data: data, Next: nil}
	}
	node := &Node{Data: data, Next: head}
	return node
}
func addTail(head *Node, data int) *Node {
	if head == nil {
		return &Node{Data: data, Next: nil}
	}
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node{Data: data, Next: nil}
	return head
}
func main() {
	var addHeadNumber int
	var addTailNumber int

	n := &Node{Data: 1}
	n.Next = &Node{Data: 2}
	n.Next.Next = &Node{Data: 3}
	n.Next.Next.Next = &Node{Data: 4}
	n.Next.Next.Next.Next = &Node{Data: 5}
	n.Next.Next.Next.Next.Next = &Node{Data: 6}
	n.Next.Next.Next.Next.Next.Next = &Node{Data: 7}

	fmt.Println("Now elements list -> :", n.Next.Data, n.Next.Next.Data, n.Next.Next.Next.Data, n.Next.Next.Next.Next.Data, n.Next.Next.Next.Next.Next.Data, n.Next.Next.Next.Next.Next.Next.Data)
	fmt.Println("[+] Insert a new head at the beginning of the list ")
	fmt.Scanln(&addHeadNumber)
	n = addHead(n, addHeadNumber)

	fmt.Println("[+] Insert a new tail at the beginning of the list ")
	fmt.Scanln(&addTailNumber)
	n = addTail(n, addTailNumber)
	fmt.Println("Now elements list -> :", n.Next.Data, n.Next.Next.Data, n.Next.Next.Next.Data, n.Next.Next.Next.Next.Data, n.Next.Next.Next.Next.Next.Data, n.Next.Next.Next.Next.Next.Next.Data)

}
