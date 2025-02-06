package heapy

import "testing"

func TestInsert(t *testing.T) {
	heap := New()

	heap.Insert(5)
	heap.Insert(6)
	if heap.Remove() != 6 {
		t.Fatalf("Expected %d found %d", 6, heap.Top())
	}
	heap.Print()
	heap.Insert(10)
	heap.Print()
}
