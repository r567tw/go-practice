package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// if (true) {
	// 	fmt.Println("true")
	// }
	// fmt.Println("Your Grade is ", input)
	fmt.Printf("Your Graade is %s",input)
}
