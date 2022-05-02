package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	flag.Parse()
	var input []int
	for i := 0; i < len(flag.Args()); i++ {
		a, _ := strconv.Atoi(flag.Arg(i))
		input = append(input, a)
		if flag.Arg(i)[0] == 'n' {
			input[len(input)-1] *= -1
		}
	}
	fmt.Println(3*input[0] + input[1])
}
