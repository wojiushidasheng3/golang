package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	second := rand.New(rand.NewSource(time.Now().UnixNano()))
	target := second.Intn(100) + 1

	fmt.Println("准备猜数字")

	for x := 0; ; x++ {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		grade, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if target > grade {
			fmt.Println("输入的数字小了")
		} else if target < grade {
			fmt.Println("输入的数字大了")
		} else {
			fmt.Println("恭喜,猜中了")
			break
		}
	}
}
