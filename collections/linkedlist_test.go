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
)

func TestLinkedListAddFirst(t *testing.T) {
	list := &LinkedList{}
	for i, v := range []int{1, 2, 3} {
		if list.Size() != i {
			t.Fatalf("Expected size %v but got %v\n", i, list.Size())
		}
		list.AddFirst(v)
	}
	if list.Size() != 3 {
		t.Fatalf("Expected size %v but got %v\n", 3, list.Size())
	}

	e := [3]int{3, 2, 1}
	v := [3]int{}
	for i := range v {
		v[i] = list.RemoveFirst().(int)
	}
	if e != v {
		t.Fatalf("Expected contents to be %v but got %v\n", e, v)
	}
}

func TestLinkedListAddLast(t *testing.T) {
	list := &LinkedList{}
	for i, v := range []int{1, 2, 3} {
		if list.Size() != i {
			t.Fatalf("Expected size %v but got %v\n", i, list.Size())
		}
		list.AddLast(v)
	}
	if list.Size() != 3 {
		t.Fatalf("Expected size %v but got %v\n", 3, list.Size())
	}

	e := [3]int{1, 2, 3}
	v := [3]int{}
	for i := range v {
		v[i] = list.RemoveFirst().(int)
	}
	if e != v {
		t.Fatalf("Expected contents to be %v but got %v\n", e, v)
	}
}

func TestLinkedListRemoveFirst(t *testing.T) {
	list := &LinkedList{}
	for _, v := range []int{1, 2, 3} {
		list.AddLast(v)
	}
	for i, e := range []int{1, 2, 3} {
		if list.Size() != 3-i {
			t.Fatalf("Expected size %v but got %v\n", i, list.Size())
		}
		v := list.RemoveFirst().(int)
		if v != e {
			t.Fatalf("Expected %v but got %v\n", e, v)
		}
	}
}

func TestLinkedListRemoveLast(t *testing.T) {
	list := &LinkedList{}
	for _, v := range []int{1, 2, 3} {
		list.AddLast(v)
	}
	for i, e := range []int{3, 2, 1} {
		if list.Size() != 3-i {
			t.Fatalf("Expected size %v but got %v\n", i, list.Size())
		}
		v := list.RemoveLast().(int)
		if v != e {
			t.Fatalf("Expected %v but got %v\n", e, v)
		}
	}
}

func TestLinkedListIterators(t *testing.T) {
	data := []int{1, 2, 3}
	reverseData := []int{3, 2, 1}

	list := &LinkedList{}
	for _, v := range []int{1, 2, 3} {
		list.AddLast(v)
	}

	t.Run("Iterator", testIterator(list, list.Iterator(), data))
	t.Run("ForwardIterator", testListIterator(list, list.ForwardIterator(), data, reverseData))
	t.Run("ReverseIterator", testListIterator(list, list.ReverseIterator(), reverseData, data))
}

func testIterator(list Collection, iter Iterator, data []int) func(t *testing.T) {
	return func(t *testing.T) {
		for i, e := range data {
			if list.Size() != len(data) {
				t.Fatalf("Expected size %v but got %v (i = %v)\n", 3, list.Size(), i)
			}
			if !iter.HasNext() {
				t.Fatalf("Expected HasNext() to return true but it returned false (i = %v)\n", i)
			}
			v := iter.Next()
			if v != e {
				t.Fatalf("Expected Next() to be %v but got %v (i = %v)\n", e, v, i)
			}
		}
		if iter.HasNext() {
			t.Fatalf("Expected HasNext() to return false but it returned true at the end of the list\n")
		}
		if iter.Next() != nil {
			t.Fatalf("Expected Next() to return nil at the end of the list but it did not.\n")
		}
	}
}

func testListIterator(list List, iter ListIterator, nextData []int, previousData []int) func(t *testing.T) {
	return func(t *testing.T) {
		for i, e := range nextData {
			if list.Size() != len(nextData) {
				t.Fatalf("Expected size %v but got %v (i = %v)\n", 3, list.Size(), i)
			}
			if !iter.HasNext() {
				t.Fatalf("Expected HasNext() to return true but it returned false (i = %v)\n", i)
			}
			v := iter.Next()
			if v != e {
				t.Fatalf("Expected Next() to be %v but got %v (i = %v)\n", e, v, i)
			}
		}
		if iter.HasNext() {
			t.Fatalf("Expected HasNext() to return false but it returned true at the end of the list\n")
		}
		for i, e := range previousData {
			if list.Size() != len(previousData) {
				t.Fatalf("Expected size %v but got %v (i = %v)\n", 3, list.Size(), i)
			}
			if !iter.HasPrevious() {
				t.Fatalf("Expected HasPrevious() to return true but it returned false (i = %v)\n", i)
			}
			v := iter.Previous()
			if v != e {
				t.Fatalf("Expected Previous() to be %v but got %v (i = %v)\n", e, v, i)
			}
		}
		if iter.HasPrevious() {
			t.Fatalf("Expected HasPrevious() to return false but it returned true at the start of the list\n")
		}
	}
}
