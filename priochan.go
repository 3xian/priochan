package priochan

import (
//"container/heap"
//"sync"
)

type PrioChan struct {
	messages messageHeap
}

func NewPrioChan(chans ...Chan) *PrioChan {
	return NewPrioChanWithSlice(chans)
}

func NewPrioChanWithSlice(chans []Chan) *PrioChan {
	return nil // TODO
}

func (pc *PrioChan) Select() (msg interface{}) {
	return nil // TODO
}

type message struct {
	body     interface{}
	priority int
}

// A messageHeap is a max-heap of priority messages.
type messageHeap []message

func (h messageHeap) Len() int           { return len(h) }
func (h messageHeap) Less(i, j int) bool { return h[i].priority > h[j].priority }
func (h messageHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *messageHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(message))
}

func (h *messageHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x.body
}
