package linkedlist

import (
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

func (list *LinkedList) Search(item int) (int, bool) {
	var elemt int
	found := false

	for node := list.sentinel.next; node != list.sentinel; node = node.next {
		if node.value == item {
			elemt = node.value
			found = true
		}
	}

	return elemt, found
}

func SentinelSetup() *Node {
	sentinel := &Node{value: -1}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	return sentinel
}
