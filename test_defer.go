package main

import (
	"fmt"
)

/*
1、defer的执行方式类似其它语言中的析构函数，在函数执行体结束后按照调用顺序的相反顺序逐个执行
2、即使函数发生严重错误也会执行
3、支持匿名函数的调用
4、通常用于资源清理、文件关闭、解锁以及记录时间等操作
5、通过与匿名函数配合可在return之后修改函数计算结果
6、如果函数体内某个变量作为defer时匿名函数的参数，则在定义defer时即已经获得了拷贝，否则则是引用某个变量的地址
7、GO没有异常机制，但有panic/recove模式来处理错误
8、Panic可以在任何地方引发，但recover只有在defer调用的函数中有效

*/

func main() {
	fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	//执行结果 1 3 2

	for i := 0; i < 3; i++ {
		//defer fmt.Println(i)
		defer func() {
			fmt.Println(i)
		}() //调用这个函数
	}
	//刚才直接打印的时候，是作为一个参数传递进去，运行到defer的时候是将这个i的值进行了一个拷贝，所以打印的是 2 1 0
	//这种情况下i一直是一个地址的引用，i一直引用的是局部变量的i，在退出这个循环体的时候 i已经变成了3，在main函数return的时候，开始执行defer语句，defer语句的时候i已经变成了3
}
