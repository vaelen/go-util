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

func TestSort(t *testing.T) {
	var comp Comparator
	list := [10]interface{}{}
	sortedList := [10]interface{}{}

	f := func(t *testing.T) {
		for name, sort := range sortFunctions {
			shuffle(list[:])
			t.Run(name, func(t *testing.T) {
				sort(list[:], comp)
				if list != sortedList {
					t.FailNow()
				}
			})
		}
	}

	// Ints
	comp = IntComparator
	for index := range list {
		list[index] = index
	}
	copy(sortedList[:], list[:])
	t.Run("Ints", f)

	// Uints
	comp = UintComparator
	for index := range list {
		list[index] = uint(index)
	}
	copy(sortedList[:], list[:])
	t.Run("Ints", f)
}
