package main

import (
	"fmt"
	"os"
)

func main() {
	str, _ := os.Getwd()
	fmt.Print(str)
}
