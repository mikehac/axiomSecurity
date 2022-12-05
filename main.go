package main

import (
	"axiomsecurity/manifestlib"
	"fmt"
)

func main() {
	manifestlib.InitComboBox()
	message := manifestlib.Hello("Mike")
	fmt.Println(message)
}
