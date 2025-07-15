package main

import (
	"fmt"
	"hellogo/src/keyboard"
	"log"
)

func main() {
	fmt.Println("输入成绩")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "pass"
	} else {
		status = "fail"
	}
	fmt.Println("a grade of", grade, "is", status)
}
