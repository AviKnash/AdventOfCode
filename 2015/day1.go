package main

import (
	"fmt"
	"os"
)

const input = "input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	data, err := os.ReadFile(input)
	check(err)
	fmt.Println(string(data))
	floor := findFloor(data)
	fmt.Println(floor)
}

func findFloor(data []byte) int {

}

