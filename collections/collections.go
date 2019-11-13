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

// An Iterator iterates through the values in a collection
type Iterator interface {
	// HasNext returns true if there are more items to iterate over.
	HasNext() bool
	// Next returns the next value in the iterator.
	// It returns nil if the iterator has reached the end of the
	// collection.
	Next() interface{}
}

// A ListIterator is an ordered iterator.
// It can iterate forwards and backwards as well as modify the list.
type ListIterator interface {
	Iterator
	// HasPrevious returns true if there are previous items to
	// iterate over.
	HasPrevious() bool
	// Previous returns the previous value from the iterator.
	// It returns nil if the iterator has reached the start of the
	// collection.
	Previous() interface{}
	// Add inserts an item into the list at the current location
	// Add(interface{})
	// Set replaces the last item returned by either next or previous
	// Set(interface{})
	// Remove removees the last item returned by either next or previous
	// Remove()
}

// A Collection is a collection of items with no guaranteed order
type Collection interface {
	// Size returns the number of items in the list
	Size() int
}

// An Iterable is a collection that can be iterated over
type Iterable interface {
	// Iterator returns an iterator that does not modify the list
	Iterator() Iterator
}

// A List is an iterable collection
type List interface {
	Collection
	Iterable
}

// A Queue is a type of list where new items are always added and removed
// in a specific order.
type Queue interface {
	Collection
	// Push adds an item to the queue
	Push(value interface{})
	// Pop removes an item from the queue
	Pop() interface{}
}

// An IterableQueue is an iterable queue
type IterableQueue interface {
	Queue
	Iterable
}
