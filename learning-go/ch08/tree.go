package main

import "fmt"

type TreeRoot[T any] struct {
	rootNode *Node[T]
	compare  func(left T, right T) int
}

func (t *TreeRoot[T]) Append(v T) {
	t.rootNode = t.rootNode.Append(v, t.compare)
}

func (t TreeRoot[T]) Print() []T {
	return t.rootNode.Print()
}

type Node[T any] struct {
	left  *Node[T]
	right *Node[T]
	val   T
}

func (n *Node[T]) Append(v T, compare func(left T, right T) int) *Node[T] {
	if n == nil {
		return &Node[T]{val: v}
	}

	switch r := compare(v, n.val); {
	case r < 0:
		n.left = n.left.Append(v, compare)
	case r > 0:
		n.right = n.right.Append(v, compare)
	}
	return n
}

func (n *Node[T]) Print() []T {
	output := make([]T, 0)
	if n == nil {
		return output
	}

	output = append(output, n.left.Print()...)
	output = append(output, n.val)
	output = append(output, n.right.Print()...)

	return output
}

func main() {
	root := TreeRoot[int]{compare: func(left int, right int) int {
		if left < right {
			return -1
		} else if left == right {
			return 0
		} else {
			return 1
		}
	}}

	root.Append(5)

	fmt.Println(root.Print())

	root.Append(1)

	root.Append(7)
	root.Append(6)
	fmt.Println(root.Print())
}
