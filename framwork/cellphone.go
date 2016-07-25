package framework

import (
	"github.com/wzdxt/go-http-test/command"
	"strings"
	"fmt"
	"reflect"
)

//receiver
type App struct {
	platform     int
	version      string
	os           string
	model        string
	channel      string

	host         string
	LastResponse Response
	nextUri      string
	normalParams ResponseResult
	userParams   ResponseResult
}

type Response struct {
	lastResponse *Response
	result       ResponseResult
}

type Params map[string]Anything
type ResponseResult map[string]Anything

func (this ResponseResult) Get(key string) Anything {
	var rr Anything = this
	keys := strings.Split(key, ".")
	for _, k := range keys {
		switch reflect.TypeOf(rr) {
		case reflect.TypeOf(ResponseResult{}):
			rr = rr.(ResponseResult)[k]
		default:
			panic(fmt.Sprintf("not a ResponseResult instance: [%T]%v", rr, rr))
		}
	}
	return rr
}

func (this *ResponseResult) Set(key string, value string) {
	var rr Anything = this
	keys := strings.Split(key, ".")
	for i, k := range keys {
		switch reflect.TypeOf(rr) {
		case reflect.TypeOf(ResponseResult{}):
			if i == len(keys) - 1 {
				rr.(ResponseResult)[k] = &value
			} else {
				if rr.(ResponseResult)[k] == nil {
					rr.(ResponseResult)[k] = ResponseResult{}
				}
				rr = rr.(ResponseResult)[k]
			}
		default:
			panic(fmt.Sprintf("not a ResponseResult instance: [%T]%v", rr, rr))
		}
	}

}

//invoker
type Customer struct {
	app      *App
	commands []command.Command
}

func (this *Customer) SetCommand(command command.Command) {
	this.commands = append(this.commands, command)
}

func (this Customer) Call() {
	for _, command := range this.commands {
		command.Execute()
	}
}

//client
type Client struct {
}

func (this Client) Run() {
	app := App{
		platform: 1,
		version:  "5000000",
		os:       "8.2.1",
		model:    "iOS",
		channel:  "App Store",
	}
	customer := Customer{}
	customer.SetCommand(SetParamsCommand{
		app:    &app,
		params: Params{"userName": 123},
	})
	customer.SetCommand(RequestCommand{app: &app})
	customer.Call()
}
