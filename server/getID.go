package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need argument")
		return
	}
	URL := arguments[1]

}
