package main

import (
	"fmt"
	"github.com/mintutu/flashtext4go/flashtext"
)

func main()  {
	processor := flashtext.NewKeyWordProcessor(true)
	processor.AddKeyWords("Scala", "Golang")
	processor.AddKeyWords("Java", "C++")
	res := processor.ReplaceKeyWords("I like Scala and Java and Scala")
	fmt.Println(res)
}


