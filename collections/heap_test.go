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
	"testing"

	"github.com/vaelen/go-util/sort"
)

func TestHeap(t *testing.T) {
	input := []interface{}{3, 5, 1, 0, 2, 4}
	output := []interface{}{0, 1, 2, 3, 4, 5}
	comp := sort.IntComparator

	t.Run("Empty", func(t *testing.T) {
		heap := NewHeap(comp)
		for _, v := range input {
			heap.Add(v)
		}
		testHeap(t, heap, output)
	})

	t.Run("Copied", func(t *testing.T) {
		heap := NewHeapFromSlice(input, comp)
		testHeap(t, heap, output)
	})
}

func testHeap(t *testing.T, heap *Heap, output []interface{}) {
	if heap.Size() != len(output) {
		t.Fatalf("Expected heap size to be %v but it was %v\n",
			len(output), heap.Size())
	}
	for i, v := range output {
		x := heap.Remove()
		if x != v {
			t.Fatalf("Expected value at index %v to be %v but it was %v\n",
				i, v, x)
		}
	}
}
