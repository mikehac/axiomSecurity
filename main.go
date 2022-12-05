package main

import (
	"axiomsecurity/manifestlib"
	"fmt"
)

func main() {
	message := manifestlib.Hello("Mike")
	fmt.Println(message)
}