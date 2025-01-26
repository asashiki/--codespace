package main

import (
	"fmt"
	"os"
	"strconv"
)

// func main() {
// 	var name = "David"
// 	name = "42" //只能字符串类型,42报错 "42"✓

// 	fmt.Println(name)
// }

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a number")
		return
	}

	input := os.Args[1]
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid number:", input)
		return
	}

	if number > 35 {
		fmt.Println("Too old")
	}else{
		fmt.Println("OK")
	}
}