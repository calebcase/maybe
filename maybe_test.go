package maybe_test

import (
	"fmt"

	"github.com/calebcase/maybe"
)

func ExampleValue() {
	// Odd returns true if the integer is odd.
	odd := func(v int) bool {
		return v%2 != 0
	}

	// NOTE: The additional type hinting `maybe.An[int](...)` is currently
	// necessary because of a limitation in Go's type inferencing. The
	// hinting may eventually be unnecessary when/if the type inferencing
	// improves for generics.
	fmt.Println(maybe.Value(false, odd, maybe.An[int](maybe.Just[int]{3})))
	fmt.Println(maybe.Value(false, odd, maybe.An[int](maybe.Nothing[int]{})))

	// These all produce the desired compile time error (because the types
	// are mismatched):
	//
	//fmt.Println(maybe.Value(false, odd, maybe.An[float32](maybe.Nothing[int]{})))
	//fmt.Println(maybe.Value(false, odd, maybe.An[float32](maybe.Nothing[float32]{})))
	//fmt.Println(maybe.Value(false, odd, maybe.An[int](maybe.Just[float32]{3})))

	// str returns the string even or odd for the value.
	str := func(v int) string {
		if v%2 == 0 {
			return "even"
		}

		return "odd"
	}

	fmt.Println(maybe.Value("unknown", str, maybe.An[int](maybe.Just[int]{3})))
	fmt.Println(maybe.Value("unknown", str, maybe.An[int](maybe.Just[int]{4})))
	fmt.Println(maybe.Value("unknown", str, maybe.An[int](maybe.Nothing[int]{})))

	// Output:
	// true
	// false
	// odd
	// even
	// unknown
}
