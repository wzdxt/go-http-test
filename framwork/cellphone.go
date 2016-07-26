package framework

import (
	"github.com/wzdxt/go-http-test/command"
	"strings"
	"fmt"
	"reflect"
	"strconv"
)

//receiver
type App struct {
	platform     int
	version      string
	os           string
	model        string
	channel      string

	Host         string
	Cookies      Params
	LastResponse Response
	Params       Params
	SavedParams  Params
	OnlyKeys     []string
	ExceptKeys   []string
}

func (this *App)SetResponse(res ResponseResult) {
	this.LastResponse = Response{
		lastResponse:&this.LastResponse,
		result:res,
	}
}

func (this *App)SetParams(params Params) {
	for k, v := range params {
		this.Params[k] = v
	}
}

func (this App)Save(keys []string) {
	for _, key := range keys {
		this.SavedParams[key] = this.Params[key]
	}
}

func (this App)Load(keys []string) {
	for _, key := range keys {
		this.Params[key] = this.SavedParams[key]
	}
}

func (this App)SaveAs(keyMap map[string]string) {
	for from, to := range keyMap {
		this.SavedParams[to] = this.Params[from]
	}
}

func (this App)LoadAs(keyMap map[string]string) {
	for from, to := range keyMap {
		this.Params[to] = this.SavedParams[from]
	}
}

func (this App)Only(keys []string) {
	this.OnlyKeys = keys
}

func (this App)Except(keys []string) {
	this.ExceptKeys = keys
}

func (this App)GetRequestParams() Params {
	switch  {
	case len(this.OnlyKeys) > 0:
		ret := make(Params, len(this.OnlyKeys))
		for _, key := range this.OnlyKeys {
			ret[key] = this.Params[key]
		}
		return ret
	default:
		ret := this.Params //todo: 获取uri对应参数
		for _, key := range this.ExceptKeys {
			delete(ret, key)
		}
		return ret
	}
}

type Response struct {
	lastResponse *Response
	result       ResponseResult
}

type Anything interface{}
type Params map[string]Anything
type ResponseResult struct {
	Params
}

func (this Params) Get(key string) Anything {
	var rr Anything = this
	keys := strings.Split(key, ".")
	for _, k := range keys {
		if i, err := strconv.Atoi(k); err == nil {
			rr = rr[i]
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

func (this *Params) Set(key string, value string) {
	var rr Anything = this
	keys := strings.Split(key, ".")
	for i, k := range keys {
		switch reflect.TypeOf(rr) {
		case reflect.TypeOf(new(ResponseResult)), reflect.TypeOf(new(Params)):
			if i == len(keys) - 1 {
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
