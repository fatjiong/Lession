package main

import "flag"

var chatbotMap = map[string]Chatbot{}

func Register(chatbot Chatbot) error{
		chatbotMap[name] = chatbot
		reurn nil
}

func Get(name string)Chatbot{
	return chatbotMap[name]
}



var chatbotName string

func init(){
	flag.StringVar(&chatbotName,"chatbot","simmple.en","The chatbot's name for dialouge")
}