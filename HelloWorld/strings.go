package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "HelloWorld"
	replacer := strings.NewReplacer("o","x")
	result := replacer.Replace(data)

	fmt.Println(result)
}
