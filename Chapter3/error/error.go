package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	err := fmt.Errorf("a height of %0.2f is invalid", -2.33333)
	fmt.Println(err.Error())
	fmt.Println(err)
}

func Err() {
	err := errors.New("a height can't be negative")
	fmt.Println(err.Error())
	log.Fatal(err)
}
