package main

import ()
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
	fmt.Printf("%T: %v\n", reflect.TypeOf(a), reflect.TypeOf(a))
	fmt.Printf("%T: %v\n", reflect.TypeOf(framework.Params{}), reflect.TypeOf(framework.Params{}))
	fmt.Printf("%T: %v\n", framework.ResponseResult{}, reflect.TypeOf(framework.ResponseResult{}))
	a["a"] = 1
	a.Set("b.c.d", "t")
	r := a.Get("b.c.d")

	fmt.Printf("%T: %v\n", r, r)
	fmt.Println(a.Get("b.c.d"))
}
