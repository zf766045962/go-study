package main

import (
	"fmt"
)

func main() {
	f := closure(10)
	res1 := f(1)
	fmt.Println(res1)
	res2 := f(2)
	fmt.Println(res2)

}

//闭包函数命名方式
func closure(x int) func(int) int {
	//fmt.Printf("%p \n", &x)
	return func(y int) int {
		//fmt.Printf("%p \n", &x)
		return x + y
	}
}
