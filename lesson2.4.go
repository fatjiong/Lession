//函数和方法
package main

// 函数
func test(testTitle string, testResult bool) (result bool, err error) {
	result = testTitle != ""
	return
}

//接口
type Talk interface {
	Hello(username string) string
	Talk(heard string) (saying string, end bool, err error)
}
