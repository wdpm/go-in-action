// Sample benchmarks to test which function is better for converting
// an integer into a string. First using the fmt.Sprintf function,
// then the strconv.FormatInt function and then strconv.Itoa.
package listing05_test

import (
	"fmt"
	"strconv"
	"testing"
)

// BenchmarkSprintf provides performance numbers for the
// fmt.Sprintf function.
func BenchmarkSprintf(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}

// BenchmarkFormat provides performance numbers for the
// strconv.FormatInt function.
func BenchmarkFormat(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

// BenchmarkItoa provides performance numbers for the
// strconv.Itoa function.
func BenchmarkItoa(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}

// BenchmarkSprintf
// BenchmarkSprintf-8      12498866                85.24 ns/op
// BenchmarkFormat
// BenchmarkFormat-8       434653472                3.161 ns/op
// BenchmarkItoa
// BenchmarkItoa-8         378260010                2.702 ns/op
