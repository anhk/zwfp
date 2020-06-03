package main

import (
	"fmt"
	"github.com/anhk/zwfp"
)

func main() {
	data, err := zwfp.Embed("Hello World.", "344433")
	if err != nil {
		panic(err)
	}
	fmt.Println("Embed:", data)

	data, key, err := zwfp.Extract(data)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data:", data)
	fmt.Println("Key:", key)
}
