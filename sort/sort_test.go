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

func testSort(size int, alreadySorted bool) func(t *testing.T) {
	return func(t *testing.T) {
		list := make([]interface{}, size)
		for name, sort := range sortFunctions {
			for index := range list {
				list[index] = index
			}
			if !alreadySorted {
				if len(list) == 2 {
					// Special case
					list[0], list[1] = list[1], list[0]
				} else {
					shuffle(list)
				}
			}
			t.Run(name, testSortFunction(sort, list, IntComparator))
		}
	}
}

func TestSort(t *testing.T) {
	t.Run("Empty", testSort(0, false))
	t.Run("1 Item", testSort(1, false))
	t.Run("2 Items, Not Sorted", testSort(2, false))
	t.Run("2 Items, Sorted", testSort(2, true))
	t.Run("3 Items, Not Sorted", testSort(3, false))
	t.Run("3 Items, Sorted", testSort(3, true))
	t.Run("10 Items, Not Sorted", testSort(10, false))
	t.Run("10 Items, Sorted", testSort(10, true))
	t.Run("100 Items, Not Sorted", testSort(100, false))
	t.Run("100 Items, Sorted", testSort(100, true))
	t.Run("1000 Items, Not Sorted", testSort(1000, false))
	t.Run("1000 Items, Sorted", testSort(1000, true))
}
