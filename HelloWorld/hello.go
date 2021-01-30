package main
import (
	"fmt"
	"reflect"
)

func main(){
	fmt.Println("Hello World")
	// 這裡請務必使用 " .." 否則很容易跳出 invalid character literal (more than one character) 的問題

	fmt.Println(reflect.TypeOf("Hello World"))
	fmt.Println(reflect.TypeOf(true))

	var q int
	var len,wid int

	len,wid = 10.0,20.0
	q = 4

	var name string = "Jim"
	fmt.Println(q)
	fmt.Println(len)
	fmt.Println(wid)
	fmt.Println(name)
}