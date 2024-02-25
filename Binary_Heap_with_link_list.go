// It's simple binary heap which work with linked list

package main

import "fmt"

type LinkList struct {
	data int
	next *LinkList
}

type BinaryHeap struct {
	array []int
}

func NewBinaryHeap() *BinaryHeap {
	return &BinaryHeap{
		array: []int{},
	}
}

func (h *BinaryHeap) Insert(value int) {
	h.array = append(h.array, value)
	h.heapifyUp(len(h.array) - 1)
}

func (h *BinaryHeap) heapifyUp(index int) {
	parent := (index - 1) / 2
	for parent >= 0 && h.array[parent] < h.array[index] {
		h.array[parent], h.array[index] = h.array[index], h.array[parent]
		index = parent
		parent = (index - 1) / 2
	}
}

func main() {
	linkList := &LinkList{data: 1}
	linkList.next = &LinkList{data: 2}
	linkList.next.next = &LinkList{data: 33}
	linkList.next.next.next = &LinkList{data: 4}
	linkList.next.next.next.next = &LinkList{data: 5}
	linkList.next.next.next.next.next = &LinkList{data: 3}

	heap := NewBinaryHeap()
	heap.Insert(linkList.data)
	heap.Insert(linkList.next.data)
	heap.Insert(linkList.next.next.data)
	heap.Insert(linkList.next.next.next.data)
	heap.Insert(linkList.next.next.next.next.data)
	heap.Insert(linkList.next.next.next.next.next.data)

	fmt.Println(heap.array)
}
