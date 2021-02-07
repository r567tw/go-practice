package main

import (
    "bufio"
	"os"
	"fmt"
)

func main(){
	var numbers []string
    file ,_ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
      content := scanner.Text()
	  numbers = append(numbers,content)
	  if scanner.Err() != nil {
		  fmt.Println("error")
	  }
    }
	file.Close()
	fmt.Println(numbers)
}