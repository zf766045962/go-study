package main

import (
	"fmt"
)

func main() {
	test(1, 2, 3, 4, 5, 6)

	/*mapvalue := make(map[int]int)

	mapvalue[2] = 3
	mapvalue[3] = 4
	mapvalue[4] = 5
	testmap(mapvalue)*/

	/*numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	testslice(numbers)

	s1 := []int{1, 2, 3, 4}*/
	a, b := 1, 2
	testslice(a, b)

	s1 := []int{1, 2, 3, 4}
	B(s1)
	fmt.Println(s1)
	slice := []int{1, 2, 3, 4, 5}
	sliceModify(&slice)
	fmt.Println(slice)

	f := F
	f() //这个时候是将A的函数类型赋值给a，在go语言当中一切皆是类型啊

	fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")
}

func test(a ...int) {
	fmt.Println(a)
}

func testmap(a ...int) {
	fmt.Println(a)
}

func testslice(a ...int) {
	//这里传进来的实际上是一个slice,引用类型
	a[0] = 3
	a[1] = 4
	//尽管我们在函数A当中接收到的是一个slice，但是它得到的是一个值拷贝
	//和直接传递一个slice的区别看函数B
	fmt.Println(a)
}

func B(s []int) {
	//这里并不是传递一个指针进去，而是对这个slice的内存地址进行了一个拷贝
	//这里还可以看到像int型、string型进行常规的参数传进去的话，只是进行了个值拷贝，slice传进去虽然也是拷贝，但是它是内存地址的拷贝
	s[0] = 4
	s[1] = 5
	s[2] = 6
	s[3] = 7
	fmt.Println(s)
	//在这里 我们看到我们在函数B当中的修改，实际上影响到了我们main函数当中的变量s1
	//如果直接传递一个slice，它的修改就会影响到这个slice的本身

}

//传递地址
func sliceModify(slice *[]int) {
	*slice = append(*slice, 6)
}

func F() {
	fmt.Println("func F")
}
