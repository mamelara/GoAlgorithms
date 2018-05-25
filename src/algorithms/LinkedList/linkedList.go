package main

import "fmt"

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
	ToString()
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

func (list *LinkedList) ToString() []byte {

	addSizeOfBrackets := 2
	listContents := make([]byte, list.Size()+addSizeOfBrackets)
	node := list.sentinel.next
	listContents[0] = byte("[")

	i := 1
	for ; node != list.sentinel; node = node.next {
		listContents[i] = byte(node.value)
		i++
	}

	listContents[(list.Size()+addSizeBrackets)-1] = byte("]")

	return listContents
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

	fmt.Println(list.ToString())
}
