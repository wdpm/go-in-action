// Sample program to show how you can't always get the
// address of a value.
package main

import "fmt"

// duration is a type with a base type of int.
type duration int

// format pretty-prints the duration value.
func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

// main is the entry point for the application.
func main() {
	d := duration(42)
	fmt.Println(d.pretty())

	// duration(42).pretty()

	// ./listing46.go:17: cannot call pointer method on duration(42)
	// ./listing46.go:17: cannot take the address of duration(42)
}

// Values        Methods Receivers
// -----------------------------------------------
// T                  (t T)
// *T                 (t T) and (t *T)

// Methods       Receivers Values
// -----------------------------------------------
// (t T)               T and *T
// (t *T)              *T
