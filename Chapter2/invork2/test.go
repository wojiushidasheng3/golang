package main

import (
	"fmt"
	"strings"
)

func main() {
	broke := "G#  r#cks!"
	replaces := strings.NewReplacer("#", "o")
	fixed := replaces.Replace(broke)
	fmt.Println(fixed)
}
