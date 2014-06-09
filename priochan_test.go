package priochan

import (
	"container/heap"
	"testing"
)

func TestMessageHeap(t *testing.T) {
	dict := map[int]string{1: "one", 2: "two", 3: "three"}

	h := &messageHeap{}
	heap.Init(h)
	heap.Push(h, message{"one", 1})
	heap.Push(h, message{"three", 3})
	heap.Push(h, message{"two", 2})

	AssertForTest(t, (*h)[0].body, "three")

	for i := 3; i >= 1; i-- {
		msg := heap.Pop(h)
		AssertForTest(t, msg, dict[i])
		AssertForTest(t, h.Len(), i-1)
	}
}
