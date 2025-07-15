package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv" //使用ParseFloat
	"strings" //使用TrimSpace
)

func main() {

	fmt.Println("enter")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "pass"
	} else {
		status = "fail"
	}
	fmt.Println("enter is ", grade, " is ", status)
}
