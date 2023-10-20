package heap

import "container/heap"

// Heap is an implement for heap sort
type Heap[T any] interface {
	// Push return a value into heap
	Push(T)
	// Pop return a value from top of heap,
	//  Attention: pop an empty heap will panic
	Pop() T
	// Peak return the top value of elements in heap
	Peak() *T
	// Len return the length of elements in heap
	Len() int
	// Empty return true if heap is empty
	Empty() bool
}

// Compare : every customized type should hava Compare function
type Compare[T any] func(x, y T) bool

// comparator contains all basic types
type comparator interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

func less[T comparator](x, y T) bool {
	return x < y
}

func more[T comparator](x, y T) bool {
	return x > y
}

func NewMinHeap[T comparator](elems ...T) Heap[T] {
	return NewHeap(less[T], elems...)
}

func NewMaxHeap[T comparator](elems ...T) Heap[T] {
	return NewHeap(more[T], elems...)
}

// NewHeap comparator is required for customized type
func NewHeap[T any](comparator Compare[T], elems ...T) Heap[T] {
	return NewContainerHeap(comparator, elems...)
}

func NewContainerHeap[T any](comparator Compare[T], elems ...T) Heap[T] {
	h := &containerHeap[T]{
		container: &container[T]{
			compare: comparator,
			elems:   elems,
		},
	}
	heap.Init(h.container)

	return h
}
