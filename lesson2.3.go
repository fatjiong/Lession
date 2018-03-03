// 数据类型
package main

const (
	DEFAULT_USE_CONST = "test,asd"
)

// 数组
var ipv4 [4]uint8 = [4]uint8{192, 168, 0, 200}

// 切片就是不定长同一类型的数组
var testSlice = []string{"test1", "test2", "test3", "test4"}

//字典，就是hash格式，key类型为字符串，value类型为bool。不定长
var testMap = map[string]bool{}

//结构体，类似于php的数组
type testStruct struct {
	name string
	talk Talk
	age  int
}
