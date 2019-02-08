package main

import (
	"fmt"

	"./build"
)

func main() {
	version, err := build.GetVersion()
	if err != nil {
		panic(err)
	}
	fmt.Println(version)
}
