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

var _ Queue = &FIFOQueue{}

// A FIFOQueue is a first-in-first-out queue.
// It is implemented as a doubly linked list.
// It implements the Queue interface.
type FIFOQueue LinkedList

// Size returns the number of items in the queue.
func (l *FIFOQueue) Size() int {
	return (*LinkedList)(l).Size()
}

// Iterator returns an iterator that doesn't modify the queue.
func (l *FIFOQueue) Iterator() Iterator {
	return (*LinkedList)(l).ReverseIterator()
}

// Push adds an item to the back of the queue.
func (l *FIFOQueue) Push(value interface{}) {
	(*LinkedList)(l).AddFirst(value)
}

// Pop removes an item from the front of the queue.
func (l *FIFOQueue) Pop() interface{} {
	return (*LinkedList)(l).RemoveLast()
}
