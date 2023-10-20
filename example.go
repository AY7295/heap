package heap

import (
	"fmt"
)

func ExampleMaxHeap() {
	ints := []int{10, 29, 38, 47, 56, 65, 74, 83, 92, 10}
	h := NewMaxHeap(ints...)
	fmt.Println(*h.Peak())
	for !h.Empty() {
		fmt.Println(h.Pop())
	}
}

func ExampleMinHeap() {
	ints := []int{10, 29, 38, 47, 56, 65, 74, 83, 92, 10}
	h := NewMinHeap(ints...)
	fmt.Println(*h.Peak())
	for !h.Empty() {
		fmt.Println(h.Pop())
	}
}

func ExampleCustomizedType() {
	// sort by age ASC
	cs1 := genCustomized()
	h1 := NewHeap(func(x, y customized) bool {
		return x.age < y.age
	}, cs1...)
	for !h1.Empty() {
		fmt.Println(h1.Pop())
	}

	// sort by salary DESC
	cs2 := genCustomized()
	h2 := NewHeap(func(x, y customized) bool {
		return x.salary > y.salary
	}, cs2...)
	for !h2.Empty() {
		fmt.Println(h2.Pop())
	}

}

type customized struct {
	name     string
	age      int
	salary   float64
	location string
}

func genCustomized() []customized {
	cs := make([]customized, 10)
	for i := 0; i < len(cs); i++ {
		cs[i] = customized{
			name:     fmt.Sprintf("name%d", i),
			age:      i,
			salary:   float64(100.0 - i),
			location: "location",
		}
	}

	return cs
}
