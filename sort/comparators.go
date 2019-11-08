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

var _ Comparator = IntComparator
var _ Comparator = Int8Comparator
var _ Comparator = Int16Comparator
var _ Comparator = Int32Comparator
var _ Comparator = Int64Comparator
var _ Comparator = UintComparator
var _ Comparator = Uint8Comparator
var _ Comparator = Uint16Comparator
var _ Comparator = Uint32Comparator
var _ Comparator = Uint64Comparator
var _ Comparator = Float32Comparator
var _ Comparator = Float64Comparator
var _ Comparator = StringComparator

// IntComparator compares ints
func IntComparator(a interface{}, b interface{}) int {
	x := a.(int)
	y := b.(int)
	if x < y {
		return -1
	} else if x == y {
		return 0
	} else {
		return 1
	}
}

// Int8Comparator compares int8s
func Int8Comparator(a interface{}, b interface{}) int {
	x := a.(int8)
	y := b.(int8)
	if x < y {
		return -1
	} else if x == y {
		return 0
	} else {
		return 1
	}
}

// Int16Comparator compares int16s
func Int16Comparator(a interface{}, b interface{}) int {
	x := a.(int16)
	y := b.(int16)
	if x < y {
		return -1
	} else if x == y {
		return 0
	} else {
		return 1
	}
}

// Int32Comparator compares int32s
func Int32Comparator(a interface{}, b interface{}) int {
	x := a.(int32)
	y := b.(int32)
	if x < y {
		return -1
	} else if x == y {
		return 0
	} else {
		return 1
	}
}

// Int64Comparator compares int64s
func Int64Comparator(a interface{}, b interface{}) int {
	x := a.(int64)
	y := b.(int64)
	if x < y {
		return -1
	} else if x == y {
		return 0
	} else {
		return 1
	}
}

// UintComparator compares uints
func UintComparator(a interface{}, b interface{}) int {
	x := a.(uint)
	y := b.(uint)
	if x < y {
		return -1
	} else if x == y {
		return 0
	} else {
		return 1
	}
}

// Uint8Comparator compares uint8s
func Uint8Comparator(a interface{}, b interface{}) int {
	x := a.(uint8)
	y := b.(uint8)
	if x < y {
		return -1
	} else if x == y {
		return 0
	} else {
		return 1
	}
}

// Uint16Comparator compares uint16s
func Uint16Comparator(a interface{}, b interface{}) int {
	x := a.(uint16)
	y := b.(uint16)
	if x < y {
		return -1
	} else if x == y {
		return 0
	} else {
		return 1
	}
}

// Uint32Comparator compares uint32s
func Uint32Comparator(a interface{}, b interface{}) int {
	x := a.(uint32)
	y := b.(uint32)
	if x < y {
		return -1
	} else if x == y {
		return 0
	} else {
		return 1
	}
}

// Uint64Comparator compares uint64s
func Uint64Comparator(a interface{}, b interface{}) int {
	x := a.(uint64)
	y := b.(uint64)
	if x < y {
		return -1
	} else if x == y {
		return 0
	} else {
		return 1
	}
}

// Float32Comparator compares float32s
func Float32Comparator(a interface{}, b interface{}) int {
	x := a.(float32)
	y := b.(float32)
	if x < y {
		return -1
	} else if x > y {
		return 1
	} else {
		return 0
	}
}

// Float64Comparator compares float64s
func Float64Comparator(a interface{}, b interface{}) int {
	x := a.(float64)
	y := b.(float64)
	if x < y {
		return -1
	} else if x > y {
		return 1
	} else {
		return 0
	}
}

// StringComparator compares strings based on code points
func StringComparator(a interface{}, b interface{}) int {
	x := a.(string)
	y := b.(string)
	if x < y {
		return -1
	} else if x > y {
		return 1
	} else {
		return 0
	}
}
