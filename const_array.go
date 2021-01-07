package quiz

// Source: https://twitter.com/davecheney/status/1346617723212582912
// A constant is untyped value (see https://blog.golang.org/constants) which should be determined in compiled time.
// The builtin `len(a)` only returns the int in compiled iff a is of string/array/pointer of array type.
// See https://golang.org/pkg/builtin/#len for reference
func ConstArray() int {
	// array/slice can utilize map-like notation to assign value in the specific index
	// e.g,
	// fmt.Println([]int{4:0, 2:1})
	// {0, 0, 1, 0}
	var a = [...]int{0, 1: 0}
	const len = len(a)
	return len
}
