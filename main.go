package main

import (
	"fmt"
	"github.com/mikeyhu/mekkanism/parser"
)

func main() {
	exp, _ := parser.Parse("(+(- 2 1) 3)")

	fmt.Println("Result: ", exp.Evaluate())

}
