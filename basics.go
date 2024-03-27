// declaring variables
package main

import (
	"fmt"
)

func main() {
	var x string = "Hello World"
	fmt.Println(x)
	var y string
	y = "Hello World"
	fmt.Println(y)
	z := "Hello World"
	fmt.Println(z)
	// z = 5 //error
	z = "Hello"
	fmt.Println(z)
}
