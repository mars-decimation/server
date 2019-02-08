package main

import (
	"fmt"

	"./build"
)

func main() {
	err := build.WriteConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Build configured")
}
