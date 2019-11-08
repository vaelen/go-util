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

var _ Sort = Insertion

func init() {
	registerSort("Insertion", Insertion)
}

// Insertion performs an in-place insertion sort of the given list.
func Insertion(list []interface{}, comp Comparator) {
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

	// For longer lists, iteratively insert the first value of the
	// remaining list into the location where it is properly sorted.

	// Because 1 item is always sorted, we assume that the first item
	// in the list is a sorted list of 1 item and start from the value
	// at the next position.
	for pos := 1; pos < len(list); pos++ {
		// Get the next value to be sorted
		value := list[pos]
		index := -1
		// Starting from the position before the current one,
		// find where this value needs to go to be sorted
		// with regards to the values to the left of it.
		for i := pos - 1; i >= 0 && comp(list[i], value) > 0; i-- {
			// As long as the value at position i is larger than
			// our target value, shift the value up to the
			// next position, making room for the new value.
			list[i+1] = list[i]
			index = i
		}
		if index >= 0 {
			// Insert our value at the place where it should go.
			list[index] = value
		}
	}
}
