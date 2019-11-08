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

var _ Sort = Merge

func init() {
	registerSort("Merge", Merge)
}

// Merge performs an in-place merge sort on the given list.
func Merge(list []interface{}, comp Comparator) {
	// Empty and single item lists are already sorted
	if len(list) <= 1 {
		return
	}

	// For a list of only 2 elements, swap if they are out of order
	if len(list) == 2 {
		if comp(list[0], list[1]) > 0 {
			list[0], list[1] = list[1], list[0]
		}
		return
	}

	// For longer lists, divide list into two pieces and recurse
	middle := len(list) / 2

	leftList := list[0:middle]
	rightList := list[middle:]

	Merge(leftList, comp)
	Merge(rightList, comp)

	// Now merge the results togther into dest
	leftIndex := 0
	rightIndex := 0
	destIndex := 0
	dest := make([]interface{}, len(list))

	for {
		leftValue := leftList[leftIndex]
		rightValue := rightList[rightIndex]

		if comp(leftValue, rightValue) > 0 {
			// Right value is larger
			dest[destIndex] = rightValue
			rightIndex++
			destIndex++
		} else {
			// Left value is smaller or equal
			dest[destIndex] = leftValue
			leftIndex++
			destIndex++
		}

		if leftIndex == len(leftList) {
			// Finished with left list
			// Copy remainder of right list
			copy(dest[destIndex:], rightList[rightIndex:])
			break
		} else if rightIndex == len(rightList) {
			// Finished with left list
			// Copy remainder of right list
			copy(dest[destIndex:], leftList[leftIndex:])
			break
		}

	}

	// Copy the sorted list back into the original list
	copy(list, dest)

}
