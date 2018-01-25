package main

import "fmt"

func main(){
	ints := []int{1,2,3,4}
	for i,d:=range ints {
		fmt.Printf("Index:%d,Value:%d\n",i,d)
	}
}
