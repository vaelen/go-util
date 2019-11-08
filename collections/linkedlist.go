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

type linkedListNode struct {
	next     *linkedListNode
	previous *linkedListNode
	value    interface{}
}

var _ List = &LinkedList{}

// LinkedList implements a doubly linked list.
// Items can be removed or added from the start or end in constant time.
// Items can be iterated over either forward or backward in linear time.
type LinkedList struct {
	start  *linkedListNode
	end    *linkedListNode
	length int
}

// Size returns the number of items in the list
func (l *LinkedList) Size() int {
	return l.length
}

// Iterator returns an iterator that iterates over the items in
// the list from first to last without changing the list.
func (l *LinkedList) Iterator() Iterator {
	return l.ForwardIterator()
}

// ForwardIterator returns an iterator that iterates over the items in
// the list from first to last
func (l *LinkedList) ForwardIterator() ListIterator {
	return &linkedListForwardIterator{
		list:     l,
		next:     l.start,
		previous: nil,
	}
}

// ReverseIterator returns an iterator that iterates over the items in
// the list from last to first
func (l *LinkedList) ReverseIterator() ListIterator {
	return &linkedListReverseIterator{
		list:     l,
		next:     nil,
		previous: l.end,
	}
}

// AddFirst adds an item to the start of the list.
func (l *LinkedList) AddFirst(value interface{}) {
	l.start = &linkedListNode{
		next:  l.start,
		value: value,
	}

	if l.start.next != nil {
		l.start.next.previous = l.start
	}

	if l.end == nil {
		l.end = l.start
	}

	l.length++
}

// AddLast adds an item to the end of the list.
func (l *LinkedList) AddLast(value interface{}) {
	l.end = &linkedListNode{
		previous: l.end,
		value:    value,
	}

	if l.end.previous != nil {
		l.end.previous.next = l.end
	}

	if l.start == nil {
		l.start = l.end
	}

	l.length++
}

// RemoveFirst removes an item from the start of the list.
func (l *LinkedList) RemoveFirst() interface{} {
	if l.start == nil {
		return nil
	}

	value := l.start.value

	l.start = l.start.next

	if l.start == nil {
		l.end = nil
	}

	l.length--

	return value
}

// RemoveLast removes an item from the end of the list.
func (l *LinkedList) RemoveLast() interface{} {
	if l.end == nil {
		return nil
	}

	value := l.end.value

	l.end = l.end.previous

	if l.end == nil {
		l.start = nil
	}

	l.length--

	return value
}

var _ Iterator = &linkedListForwardIterator{}
var _ ListIterator = &linkedListForwardIterator{}

type linkedListForwardIterator struct {
	list     *LinkedList
	next     *linkedListNode
	previous *linkedListNode
	last     *linkedListNode
}

func (i *linkedListForwardIterator) HasNext() bool {
	return i.next != nil
}

func (i *linkedListForwardIterator) Next() interface{} {
	if i.next == nil {
		return nil
	}

	i.last = i.next
	value := i.last.value
	i.next = i.last.next
	i.previous = i.last

	return value
}

func (i *linkedListForwardIterator) HasPrevious() bool {
	return i.previous != nil
}

func (i *linkedListForwardIterator) Previous() interface{} {
	if i.previous == nil {
		return nil
	}

	i.last = i.previous
	value := i.last.value
	i.next = i.last
	i.previous = i.last.previous

	return value
}

var _ Iterator = &linkedListReverseIterator{}
var _ ListIterator = &linkedListReverseIterator{}

type linkedListReverseIterator linkedListForwardIterator

func (i *linkedListReverseIterator) HasNext() bool {
	return (*linkedListForwardIterator)(i).HasPrevious()
}

func (i *linkedListReverseIterator) Next() interface{} {
	return (*linkedListForwardIterator)(i).Previous()
}

func (i *linkedListReverseIterator) HasPrevious() bool {
	return (*linkedListForwardIterator)(i).HasNext()
}

func (i *linkedListReverseIterator) Previous() interface{} {
	return (*linkedListForwardIterator)(i).Next()
}
