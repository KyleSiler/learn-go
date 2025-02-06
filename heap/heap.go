package heapy

import "fmt"

type Heap struct {
	heap []int
	size int
	max  int
}

func New() Heap {
	return Heap{
		size: 0,
		max:  1,
		heap: make([]int, 1),
	}
}

func (h *Heap) Insert(n int) {
	if h.size == h.max {
		h.max *= 2
		tmp := make([]int, h.max*2)
		copy(tmp, h.heap)
		h.heap = tmp
	}
	h.heap[h.size] = n

	cur := h.size
	for cur > 0 {
		parent := (cur - 1) / 2
		if h.heap[cur] > h.heap[parent] {
			h.heap[cur], h.heap[parent] = h.heap[parent], h.heap[cur]
			cur = parent
		} else {
			break
		}
	}
	h.size++
}

func (h Heap) Print() {
	fmt.Println(h.heap[:h.size-1])
}

func (h Heap) Top() int {
	return h.heap[0]
}

func (h Heap) Remove() int {
	top := h.heap[0]

	h.size -= 1
	h.heap[0] = h.heap[h.size]

	index := 0
	for index < (h.size/2) && index <= h.size {
		cur := h.heap[index]

		l := left(index)
		r := right(index)

		if cur < h.heap[l] || cur < h.heap[r] {
			if h.heap[l] > h.heap[r] {
				h.heap[index] = h.heap[l]
				h.heap[l] = cur
				index = l
			} else {
				h.heap[index] = h.heap[r]
				h.heap[r] = cur
				index = r
			}
		}
	}

	return top
}

func left(n int) int {
	return 2*n + 1
}

func right(n int) int {
	return 2*n + 2
}
