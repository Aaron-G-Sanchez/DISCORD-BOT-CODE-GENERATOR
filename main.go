package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println(dir)
}
