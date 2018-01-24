package main

import (
	"fmt"
	"strings"
)

//switch

var content string

func main(){
	switch content{
		default:
			fmt.Println("Unknow Language")
		case "Python":
			fmt.Println("An interpreted Language")
		case "Go":
			fmt.Println("A compiled langugae")
		case "Php":
			fmt.Println("The best language")
	}


	switch lang:=strings.TrimSpace(content);lang{
		default:
			fmt.Println("Unknow Language")
		case "Python":
			fmt.Println("An interpreted Language")
		case "Go":
			fmt.Println("A compiled langugae")
		case "Php":
			fmt.Println("The best language")
	}

	switch lang:=strings.TrimSpace(content);lang{
		case "Ruby"
			fallthrough
		case "Python":
			fmt.Println("An interpreted Language")
		case "Go","Java","C":
			fmt.Println("A compiled langugae")
			break
		case "Php":
			fmt.Println("The best language")
		default:
			fmt.Println("Unknow Language")
	}

	var v interface{}

	switch v.(type){
	case string:
			fmt.Printf("The string is '%s'.\n",v.(string))
	case int,uint,int8,uint8,int16,uint16,int32,uint32:
			fmt.Printf("The integer is %d.\n",v)
	default :
		fmt.Printf("Unsupported value. (type=%T)\n",v)
	}

	switch i:=v.(type){
	case string:
		fmt.Printf("The string is '%s'.\n",i)
	case int,uint,int8,uint8,int16,uint16,int32,uint32:
		fmt.Printf("The integer is %d.\n",i)
	default :
		fmt.Printf("Unsupported value. (type=%T)\n",i)
	}

}