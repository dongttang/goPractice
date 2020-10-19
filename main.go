package main

import (
	"fmt"
	"github.com/matishsiao/goInfo"
)

func main() {

	fmt.Println("hello, world!")

	gi := goInfo.GetInfo()
	gi.VarDump()

}
