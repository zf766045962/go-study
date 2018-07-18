package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("2")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("3")
	}()
	//test()  panic抛出异常

	testHaha()
}

func test() {
	fmt.Println("1")
	panic("bug")
	fmt.Println("4")
}

func testHaha() {
	fmt.Println("1")
	arr := []string{"test", "test1"}
	fmt.Println(arr[3])
	panic("bug")
}
