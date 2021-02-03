package main

import (
	"fmt"
)

func main() {
	// 第一種
	var todos [2]string
	todos[0] = "learning go !"
	todos[1] = "use go to write an app"
	fmt.Println(todos[1])
	// 第二種
	var grades [3]int = [3]int{90, 98, 93}
	fmt.Println(grades[2])
	// 第三種
	heights := [3]int{90,98,93}
	fmt.Println(heights[0])
}
