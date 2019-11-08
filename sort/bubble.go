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

var _ Sort = Bubble

func init() {
	registerSort("Bubble", Bubble)
}

// Bubble performs an in-place bubble sort of the given list.
func Bubble(list []interface{}, comp Comparator) {
	lastIndex := len(list) - 1
	for {
		changes := 0
		for index := range list {
			if index == lastIndex {
				break
			}
			a := list[index]
			b := list[index+1]
			if comp(a, b) > 0 {
				list[index+1] = a
				list[index] = b
				changes++
			}

		}
		if changes == 0 {
			break
		}
		lastIndex--
	}
}
