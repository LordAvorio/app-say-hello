package main

import (
	"fmt"
	go_say_hello "github.com/LordAvorio/go-say-hello/v2"
) 

func main(){
	var myName = "Gading"
	fmt.Println(go_say_hello.SayHello(myName))
}