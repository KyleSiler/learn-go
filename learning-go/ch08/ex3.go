package main

import "fmt"

type LinkedList[T comparable] struct {
	Head *LLNode[T]
	Tail *LLNode[T] // TODO: Reimplement methods with a tail. It simplifies some operations
}

func (ll *LinkedList[T]) Add(value T) {
	newNode := &LLNode[T]{Value: value}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail.Next = newNode
		ll.Tail = newNode
	}
}

func (ll *LinkedList[T]) Insert(value T, index int) {
	newNode := &LLNode[T]{Value: value}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
		return
	}

	if index == 0 {
		newNode.Next = ll.Head
		ll.Head = newNode
		return
	}

	curNode := ll.Head
	for i := 1; i < index; i++ {
		if curNode.Next == nil {
			curNode.Next = newNode
			ll.Tail = newNode
			return
		}
		curNode = curNode.Next
	}

	newNode.Next = curNode.Next
	curNode.Next = newNode
	if curNode == ll.Tail {
		ll.Tail = newNode
	}
}

func (ll *LinkedList[T]) Index(value T) int {
	i := 0

	for curNode := ll.Head; curNode != nil; curNode = curNode.Next {
		if curNode.Value == value {
			return i
		}
		i++
	}

	return -1
}

type LLNode[T comparable] struct {
	Value T
	Next  *LLNode[T]
}

func main() {
	l := &LinkedList[int]{}
	l.Add(5)
	l.Add(10)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))

	l.Insert(100, 0)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	l.Insert(200, 1)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(200))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Insert(300, 10)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Add(400)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Insert(500, 6)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}
}
