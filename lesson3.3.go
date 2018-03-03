package main

import (
	"fmt"
)

func narcissistic(n int) bool {
	third := n % 10
	second := (n / 10) % 10
	first := (n / 100) % 10

	if sumTotal := first*first*first + second*second*second + third*third*third; sumTotal == n {
		return true
	}

	return false
}

//打印出100-1000之间的水仙花数
func main() {
	for i := 101; i < 1000; i++ {
		if narcissistic(i) == true {
			fmt.Printf("[%d] 是立方水仙花数 \n", i)
		}
	}
}
