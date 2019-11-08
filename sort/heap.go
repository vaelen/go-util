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

package sort

var _ Sort = Heap

func init() {
	registerSort("Heap", Heap)
}

// Heap performs an in-place heap sort of the given list.
func Heap(list []interface{}, comp Comparator) {
	// Lists of 1 or 0 elements are already sorted
	if len(list) < 2 {
		return
	}

	// For a list of only 2 elements, swap if they are out of order
	if len(list) == 2 {
		if comp(list[0], list[1]) > 0 {
			list[0], list[1] = list[1], list[0]
		}
		return
	}

	// For longer lists, first convert the slice into a heap
	for i := (len(list) / 2) - 1; i >= 0; i-- {
		heapify(list, comp, i)
	}

	// The value at index 0 is now the largest value,
	// so swap it with the end of the list and iterate
	for i := len(list) - 1; i >= 0; i-- {
		list[i], list[0] = list[0], list[i]
		heapify(list[:i], comp, 0)
	}
}

// Recursively swap values around to ensure a valid heap
func heapify(list []interface{}, comp Comparator, index int) {
	largest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < len(list) && comp(list[left], list[largest]) > 0 {
		largest = left
	}

	if right < len(list) && comp(list[right], list[largest]) > 0 {
		largest = right
	}

	if largest != index {
		list[largest], list[index] = list[index], list[largest]
		heapify(list, comp, largest)
	}
}
