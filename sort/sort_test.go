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

import (
	"fmt"
	"math/rand"
	"testing"
)

func shuffle(list []interface{}) {
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
}

func TestIsSorted(t *testing.T) {
	if IsSorted([]interface{}{5, 3, 1, 4, 2}, IntComparator) {
		t.Fatalf("IsSorted() returned true for unsorted list\n")
	}
	if !IsSorted([]interface{}{1, 2, 3, 4, 5}, IntComparator) {
		t.Fatalf("IsSorted() returned false for sorted list\n")
	}
}

func testSortFunction(sort Sort, list []interface{}, comp Comparator) func(t *testing.T) {
	return func(t *testing.T) {
		sort(list, IntComparator)
		if !IsSorted(list, IntComparator) {
			t.FailNow()
		}
	}
}

func testSort(size int) func(t *testing.T) {
	list := make([]interface{}, size)
	return func(t *testing.T) {
		sorted, shuffled := createTestLists(size)
		for name, sort := range sortFunctions {
			t.Run(name, func(t *testing.T) {
				copy(list, sorted)
				t.Run("Sorted", testSortFunction(sort, list, IntComparator))
				copy(list, shuffled)
				t.Run("Shuffled", testSortFunction(sort, list, IntComparator))
			})
		}
	}
}

func TestSort(t *testing.T) {
	sizes := []int{0, 1, 2, 3, 1e1, 1e2, 1e3}
	for _, size := range sizes {
		t.Run(fmt.Sprintf("%d Item(s)", size), testSort(size))
	}
}

func createTestLists(size int) (sorted []interface{}, shuffled []interface{}) {
	sorted = make([]interface{}, size)
	shuffled = make([]interface{}, size)
	for index := range sorted {
		sorted[index] = index
		shuffled[index] = index
	}
	if size == 2 {
		// Special case
		shuffled[0], shuffled[1] = sorted[1], sorted[0]
	} else {
		shuffle(shuffled)
	}
	return sorted, shuffled
}

// This is here to prevent compiler optimizations from artificially affecting the benchmarking tests.
var testList []interface{}

func benchmarkSortFunction(sort Sort, input []interface{}, comp Comparator) func(b *testing.B) {
	return func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			list := make([]interface{}, len(input))
			for pb.Next() {
				copy(list, input)
				sort(list, comp)
				testList = list
			}
		})
	}
}

func benchmarkSort(input []interface{}, comp Comparator) func(b *testing.B) {
	return func(b *testing.B) {
		for name, sort := range sortFunctions {
			b.Run(name, benchmarkSortFunction(sort, input, comp))
		}
	}
}

func benchmarkSortSize(size int) func(b *testing.B) {
	return func(b *testing.B) {
		sorted, shuffled := createTestLists(size)
		b.Run("Shuffled", benchmarkSort(shuffled, IntComparator))
		b.Run("Sorted", benchmarkSort(sorted, IntComparator))
	}
}

func BenchmarkSort(b *testing.B) {
	sizes := []int{0, 1, 2, 3, 1e1, 1e2, 1e3, 1e4}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d Item(s)", size), benchmarkSortSize(size))
	}
}
