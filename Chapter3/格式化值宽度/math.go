package main

import "fmt"

func main() {
	fmt.Printf("%12s |  %s\n", "protect", "Cost in cents")
	fmt.Printf("%12s |  %2d\n", "hello world", 12)
	fmt.Printf("%12s |  %2d\n", "hello", 1)

	fmt.Printf("%%7.3f: %7.3f\n", 12.3456)
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456)
	fmt.Printf("%%.1f: %.1f\n", 12.3456)
	fmt.Printf("%%.2f: %.2f\n", 12.3456)
}
