package linkedlist

import "testing"

func createSentinel() *Node {
	sentinel := &Node{value: -1}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	return sentinel
}

func TestLinkedListSize(t *testing.T) {

	list := LinkedList{sentinel: createSentinel()}

	N := list.Size()
	if N != 0 {
		t.Error("Expected empty list, got", N)
	}
}

func TestLinkedListInsert(t *testing.T) {

	list := LinkedList{sentinel: createSentinel()}

	for i := 1; i <= 3; i++ {
		list.InsertFront(i)
	}

	N := list.Size()
	if N != 3 {
		t.Error("Expected list of size 3, got", N)
	}
}

func TestLinkedListSearch(t *testing.T) {

	list := LinkedList{sentinel: createSentinel()}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			list.InsertFront(i)
		} else {
			list.Append(i)
		}
	}

	if _, found := list.Search(-1); found {
		t.Error("Expected -1 not found, got", found)
	}

	if _, found := list.Search(9); !found {
		t.Error("Expected to find 10, got", found)
	}
}
