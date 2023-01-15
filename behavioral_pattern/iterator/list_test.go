package iterator

import "testing"

func TestList(t *testing.T) {
	list := NewArrayList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	iterator := list.Iterator()
	for iterator.HasNext() {
		t.Log(iterator.CurrentItem())
		iterator.Next()
	}
}
