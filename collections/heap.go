/*
MIT License

Copyright (c) 2019 Andrew C. Young

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package collections

import (
	"github.com/vaelen/go-util/sort"
)

var _ Collection = &Heap{}
var _ Queue = &Heap{}

// Heap represents a Minimum Heap
type Heap struct {
	data []interface{}
	comp sort.Comparator
}

// NewHeap constructs an empty heap.
func NewHeap(comp sort.Comparator) *Heap {
	return &Heap{
		data: make([]interface{}, 0),
		comp: comp,
	}
}

// NewHeapFromSlice constructs a new heap from the given slice.
// It creates a copy of the slice so that the original slice is
// not affected.
func NewHeapFromSlice(list []interface{}, comp sort.Comparator) *Heap {
	heap := &Heap{
		data: make([]interface{}, len(list)),
		comp: comp,
	}

	copy(heap.data, list)

	for i := (len(list) / 2) - 1; i >= 0; i-- {
		heap.heapify(i)
	}

	return heap
}

// Add adds a new item to the heap
func (h *Heap) Add(value interface{}) {
	h.data = append(h.data, value)
	if len(h.data) > 1 {
		// Check the parent node and swap if this node is smaller.
		h.siftUp(len(h.data) - 1)
	}
}

// Remove returns the first item from the heap
func (h *Heap) Remove() interface{} {
	if len(h.data) == 0 {
		return nil
	}
	// Remove the first value
	value := h.data[0]
	// Move the last value to the front
	h.data[0] = h.data[len(h.data)-1]
	// Shrink the array
	h.data = h.data[:len(h.data)-1]
	// Sift the front value down to its correct place
	h.heapify(0)
	return value
}

// Size returns the number of items in the heap
func (h *Heap) Size() int {
	return len(h.data)
}

func (h *Heap) Push(value interface{}) {
	h.Add(value)
}

func (h *Heap) Pop() interface{} {
	return h.Remove()
}

// heapify will recursively swap values in an list to ensure a valid heap
func (h *Heap) heapify(index int) {
	smallest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < len(h.data) && h.comp(h.data[left], h.data[smallest]) < 0 {
		smallest = left
	}

	if right < len(h.data) && h.comp(h.data[right], h.data[smallest]) < 0 {
		smallest = right
	}

	if smallest != index {
		// swap and check again
		h.data[smallest], h.data[index] =
			h.data[index], h.data[smallest]
		h.heapify(smallest)
	}
}

// siftUp will recursively swap values towards the top of a list
// to ensure a valid heap.
func (h *Heap) siftUp(index int) {
	if index == 0 {
		return
	}

	parent := ((index + 1) / 2) - 1

	if h.comp(h.data[index], h.data[parent]) < 0 {
		// If the child is larger than the parent, swap
		h.data[index], h.data[parent] =
			h.data[parent], h.data[index]
		h.siftUp(parent)
	}
}
