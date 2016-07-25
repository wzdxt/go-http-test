package main

import (
)
import (
	"fmt"
	"github.com/wzdxt/go-http-test/framwork"
	"reflect"
)

func main() {
	//client := framework.Client{}
	//client.Run()
	a := framework.Params{}
	//a := map[string]int{}
	fmt.Printf("%T: %v\n", reflect.TypeOf(a), reflect.TypeOf(a) )
	fmt.Printf("%T: %v\n", reflect.TypeOf(framework.Params{}), reflect.TypeOf(framework.Params{}) )
	a["a"] = 1
	fmt.Println(a["a"])
}
