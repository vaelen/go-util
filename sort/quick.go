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

var _ Sort = Quick

func init() {
	registerSort("Quick", Quick)
}

// Quick performs an in-place quicksort of the given list.
func Quick(list []interface{}, comp Comparator) {
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

	// For longer lists, perform a Quicksort
	// First choose a pivot using the median-of-three method
	a, b, c := list[0], list[len(list)/2], list[len(list)-1]
	if comp(a, b) > 0 {
		b, a = a, b
	}
	if comp(b, c) > 0 {
		b, c = c, b
	}
	if comp(a, b) > 0 {
		b, a = a, b
	}

	// The pivot value will be the median of the three values
	pivot := b

	// Now divide the list into two buckets based on the pivot value
	// This is done by looking at both ends of the list and moving
	// towards the center.  When the pivot is in its final sorted
	// position, the top and bottom indexes will be equal.

	bottom := 0
	top := len(list) - 1

	for bottom < top {
		// Compare the bottom value to the pivot
		x := comp(list[bottom], pivot)
		if x < 0 {
			// In the right bucket, move to the next value
			bottom++
		} else if x > 0 {
			// Move to the other bucket by swapping values
			list[bottom], list[top] = list[top], list[bottom]
			if x > 0 {
				top--
			}
		}

		// Exit the loop if the indexes converge
		if top <= bottom {
			break
		}

		// Compare the top value to the pivot
		x = comp(list[top], pivot)
		if x > 0 {
			// In the right bucket, move to the next value
			top--
		} else if x <= 0 {
			// Move to the other bucket by swapping values
			list[bottom], list[top] = list[top], list[bottom]
			if x < 0 {
				bottom++
			}
		}
	}

	// Finally, sort the two sub arrays (minus the pivot)
	Quick(list[0:bottom], comp)
	Quick(list[top+1:], comp)

}
