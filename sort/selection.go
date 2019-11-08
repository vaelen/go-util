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

var _ Sort = Selection

func init() {
	registerSort("Selection", Selection)
}

// Selection performs an in-place selection sort of the given list.
func Selection(list []interface{}, comp Comparator) {
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

	// For longer lists, iteratively find the smallest value of the
	// remaining list and place it at the front.
	for pos := range list {
		var minValue interface{}
		minPos := -1
		for i, value := range list[pos:] {
			if minValue == nil || comp(value, minValue) < 0 {
				minValue = value
				minPos = i
			}
		}
		list[pos+minPos], list[pos] = list[pos], list[pos+minPos]
	}
}
