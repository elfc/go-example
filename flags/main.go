package main

import (
	"flag"
	"fmt"
)

func main() {
	b := flag.Bool("b", true, "布尔值")
	str := flag.String("str", "test", "字符串")

	flag.Parse() // Scan the arguments list

	fmt.Println("b: ", *b)
	fmt.Println("str: ", *str)
}
