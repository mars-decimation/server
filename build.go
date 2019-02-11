package main

import (
	"fmt"

	"./build"
)

// main writes the build configuration
func main() {
	err := build.WriteConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Build configured")
}
