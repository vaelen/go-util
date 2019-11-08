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

import "testing"

func testComparator(comp Comparator, lower interface{}, higher interface{}) func(t *testing.T) {
	return func(t *testing.T) {
		v := comp(lower, lower)
		if v != 0 {
			t.Fatalf("Expected comp(%v, %v) to equal 0 but got %v\n", lower, lower, v)
		}

		v = comp(lower, higher)
		if v >= 0 {
			t.Fatalf("Expected comp(%v, %v) to be < 0 but got %v\n", lower, higher, v)
		}

		v = comp(higher, lower)
		if v <= 0 {
			t.Fatalf("Expected comp(%v, %v) to be > 0 but got %v\n", higher, lower, v)
		}
	}
}

func TestComparators(t *testing.T) {
	t.Run("Int", testComparator(IntComparator, int(-100), int(100)))
	t.Run("Int8", testComparator(Int8Comparator, int8(-100), int8(100)))
	t.Run("Int16", testComparator(Int16Comparator, int16(-100), int16(100)))
	t.Run("Int32", testComparator(Int32Comparator, int32(-100), int32(100)))
	t.Run("Int64", testComparator(Int64Comparator, int64(-100), int64(100)))
	t.Run("Uint", testComparator(UintComparator, uint(100), uint(200)))
	t.Run("Uint8", testComparator(Uint8Comparator, uint8(100), uint8(200)))
	t.Run("Uint16", testComparator(Uint16Comparator, uint16(100), uint16(200)))
	t.Run("Uint32", testComparator(Uint32Comparator, uint32(100), uint32(200)))
	t.Run("Uint64", testComparator(Uint64Comparator, uint64(100), uint64(200)))
	t.Run("Float32", testComparator(Float32Comparator, float32(1.234), float32(1.456)))
	t.Run("Float64", testComparator(Float64Comparator, float64(1.234), float64(1.456)))
	t.Run("String", testComparator(StringComparator, "AAA", "ZZZ"))
}
