package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedList struct {
	sentinel *Node
	size     int
}

type listMethods interface {
	Size() int
	InsertFront()
	Append()
	String()
	Search()
}

func insertFirstNode(prevNode, newNode *Node) {
	newNode.prev = prevNode
	newNode.next = prevNode
	prevNode.prev = newNode
	prevNode.next = newNode
}

func (list *LinkedList) Size() int {
	return list.size
}

func (list *LinkedList) insertFront(node *Node, item int) {
	newNode := &Node{value: item}
	if list.Size() == 0 {
		insertFirstNode(node, newNode)
	} else {
		head := list.sentinel.next
		newNode.next = head
		head.prev = newNode
		newNode.prev = list.sentinel
		list.sentinel.next = newNode
	}
	list.size++
}

func (list *LinkedList) InsertFront(item int) {
	list.insertFront(list.sentinel, item)
}

func (list *LinkedList) insertBack(node *Node, item int) {
	newNode := &Node{value: item}
	if list.Size() == 0 {
		insertFirstNode(node, newNode)
	} else {
		tail := list.sentinel.prev
		tail.next = newNode
		newNode.prev = tail
		newNode.next = list.sentinel
		list.sentinel.prev = newNode
	}
	list.size++
}

func (list *LinkedList) Append(item int) {
	list.insertBack(list.sentinel, item)
}

func (list *LinkedList) String() string {
	elements := make([]string, list.Size()+2)

	elements = append(elements, "[")

	for node := list.sentinel.next; node != list.sentinel; node = node.next {
		elements = append(elements, strconv.Itoa(node.value))
	}

	elements = append(elements, "]")

	return strings.Join(elements, " ")
}

func SentinelSetup() *Node {
	sentinel := &Node{value: -1}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	return sentinel
}

func main() {

	var prompt string = "current size of list"

	sentinel := SentinelSetup()

	list := LinkedList{sentinel: sentinel}
	fmt.Println(prompt, list.Size())

	list.InsertFront(1)
	fmt.Println(prompt, list.Size())

	list.Append(2)
	fmt.Println(prompt, list.Size())

	list.InsertFront(3)
	fmt.Println(prompt, list.Size())

	fmt.Println(list.String())
}
