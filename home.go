// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func main() {
	// fmt.Println(add(1, 3))
	// fmt.Println(divmod(233, 212))

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	doubleAt(values, 4)
	fmt.Println(values)
}
