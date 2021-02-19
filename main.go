package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./test.log")
	if err != nil {
		fmt.Println("Error open file")
		os.Exit(-1)
	}
	file.Close()
	InitPlatform()
	menu()
}
