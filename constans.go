// Go supports _constants_ of character, string, boolean,
// and numeric values.

package main

import "fmt"
import _ "math"

// `const` declares a constant value.
const s string = "constant"

func main() {
	fmt.Println(s)

	// A `const` statement can appear anywhere a `var`
	// statement can.
	const n = 500000000

	// Constant expressions perform arithmetic with
	// arbitrary precision.
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it's given