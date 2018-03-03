package main

import "fmt"

//defer

func outerFunc() {
	defer fmt.Println("函数执行结束前打印")
	fmt.Println("实际第一个打印")
}

func main() {
	outerFunc()
	printNunmbers()
}

func printNunmbers() {
	for i := 0; i < 5; i++ {
		defer func(n int) {
			fmt.Printf("%d", n)
		}(i * 2)
	}
}
