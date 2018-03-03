package main

import "fmt"

func scal(str string) (strCount, spaceCount, numCount int) {
	utfChars := []rune(str)

	for i := 0; i < len(utfChars); i++ {
		if utfChars[i] >= 'a' && utfChars[i] <= 'z' || utfChars[i] >= 'A' && utfChars[i] <= 'Z' {
			strCount++
		}
		if utfChars[i] == ' ' {
			spaceCount++
		}
		if utfChars[i] >= '0' && utfChars[i] <= '9' {
			numCount++
		}
	}

	return
}

//统计一条字符串中出现的字符数，空格数量和数字数量
func main() {
	str := "adasd123123123 asdassdsadas 0     asdasdasdas 我在测试数据"
	strCount, spaceCount, numCount := scal(str)
	fmt.Printf("该字符串中有%d个字符，%d个空格，%d个数字", strCount, spaceCount, numCount)

}
