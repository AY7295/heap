package heap

import (
	"container/heap"
)

type container[T any] struct {
	elems   []T
	compare Compare[T]
}

func (ci *container[T]) Len() int {
	return len(ci.elems)
}

func (ci *container[T]) Less(i, j int) bool {
	return ci.compare(ci.elems[i], ci.elems[j])
}

func (ci *container[T]) Swap(i, j int) {
	ci.elems[i], ci.elems[j] = ci.elems[j], ci.elems[i]
}

func (ci *container[T]) Push(x any) {
	ci.elems = append(ci.elems, x.(T))
}

func (ci *container[T]) Pop() any {
	if ci.Len() == 0 {
		panic("heap is empty")
	}
	defer func() {
		ci.elems = ci.elems[:ci.Len()-1]
	}()
	return ci.elems[ci.Len()-1]
}

type containerHeap[T any] struct {
	*container[T]
}

func (m *containerHeap[T]) Push(x T) {
	heap.Push(m.container, x)
}

func (m *containerHeap[T]) Pop() T {
	return heap.Pop(m.container).(T)
}

func (m *containerHeap[T]) Peak() *T {
	if m.Empty() {
		return nil
	}

	return &m.elems[0]
}

func (m *containerHeap[T]) Empty() bool {
	return m.Len() == 0
}
