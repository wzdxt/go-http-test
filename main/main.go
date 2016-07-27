package main

import (
	"fmt"
	"github.com/wzdxt/go-http-test/framework/common"
	"reflect"
)

func main() {
	//client := framework.Client{}
	//client.Run()
	a := common.Params{}
	//a := map[string]int{}
	fmt.Printf("%T: %v\n", reflect.TypeOf(a), reflect.TypeOf(a))
	fmt.Printf("%T: %v\n", reflect.TypeOf(common.Params{}), reflect.TypeOf(common.Params{}))
	fmt.Printf("%T: %v\n", common.ResponseResult{}, reflect.TypeOf(common.ResponseResult{}))
	a["a"] = 1
	a.Set("b.c.d", "t")
	r := a.Get("b.c.d")

	fmt.Printf("%T: %v\n", r, r)
	fmt.Println(a.Get("b.c.d"))
}
