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

func TestQueue(t *testing.T) {
	data := []int{1, 2, 3}
	reverseData := []int{3, 2, 1}
	t.Run("FIFOQueue", testQueue(&FIFOQueue{}, data, data))
	t.Run("Stack", testQueue(&Stack{}, data, reverseData))
	t.Run("Heap", testQueue(NewHeap(sort.IntComparator), data, data))
}

func testQueue(queue Queue, inData []int, outData []int) func(t *testing.T) {
	return func(t *testing.T) {
		if queue.Size() != 0 {
			t.Fatalf("Expected starting size to be 0 but it was %v\n", queue.Size())
		}
		if queue.Pop() != nil {
			t.Fatalf("Expected Pop to return nil but it didn't\n")
		}

		for i, v := range inData {
			if queue.Size() != i {
				t.Fatalf("Expected size %v but got %v\n", i, queue.Size())
			}
			queue.Push(v)
		}
		if queue.Size() != len(inData) {
			t.Fatalf("Expected size %v but got %v\n", len(inData), queue.Size())
		}

		if iq, ok := queue.(IterableQueue); ok {
			t.Run("Iterator", testIterator(iq, iq.Iterator(), outData))
		}
		for i, e := range outData {
			if queue.Size() != len(outData)-i {
				t.Fatalf("Expected size %v but got %v\n", i, queue.Size())
			}
			v := queue.Pop()
			if e != v {
				t.Fatalf("Expected popped value to be %v but got %v\n", e, v)
			}
		}
		if queue.Size() != 0 {
			t.Fatalf("Expected size %v but got %v\n", 0, queue.Size())
		}

		if queue.Pop() != nil {
			t.Fatalf("Expected Pop to return nil but it didn't\n")
		}
	}
}
