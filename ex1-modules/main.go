package main

import (
	"fmt"

	"github.com/porwalameet/golang/modules/nummanip/calc"
)

func main() {
	result := calc.Add(2, 3)
	fmt.Println("Addition", result)
}