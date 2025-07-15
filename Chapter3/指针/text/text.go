package main

import "fmt"

func negate(myBoolean *bool) {
	*myBoolean = !(*myBoolean)
}

func main() {
	trush := true
	negate(&trush)
	fmt.Println(trush)

	lies := false
	negate(&lies)
	fmt.Println(lies)
}
