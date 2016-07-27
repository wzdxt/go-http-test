package common

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Response struct {
	Previous *Response
	Result   ResponseResult
}

type Anything interface{}
type Params map[string]Anything
type ResponseResult Params

func (this Params) Get(key string) Anything {
	var rr Anything = this
	keys := strings.Split(key, ".")
	for _, k := range keys {
		if i, err := strconv.Atoi(k); err == nil {
			rr = rr.([]Params)[i]
		} else {
			switch reflect.TypeOf(rr) {
			case reflect.TypeOf(ResponseResult{}), reflect.TypeOf(Params{}):
				rr = rr.(Params)[k]
			default:
				panic(fmt.Sprintf("not a ResponseResult instance: [%T]%v", rr, rr))
			}
		}
	}
	return rr
}

func (this *Params) Set(key string, value Anything) {
	var rr Anything = this
	keys := strings.Split(key, ".")
	for i, k := range keys {
		switch reflect.TypeOf(rr) {
		case reflect.TypeOf(new(ResponseResult)), reflect.TypeOf(new(Params)):
			if i == len(keys)-1 {
				(*rr.(*Params))[k] = value
			} else {
				if (*(rr.(*Params)))[k] == nil {
					m := make(Params)
					(*rr.(*Params))[k] = m
				}
				m := (*rr.(*Params))[k].(Params)
				rr = &m
			}
		default:
			panic(fmt.Sprintf("not a ResponseResult instance: [%T]%v", rr, rr))
		}
	}

}
