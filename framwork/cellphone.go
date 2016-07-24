package framework

import (
	"github.com/wzdxt/go-http-test/command"
)

//receiver
type App struct {
	platform int
	version  string
	os       string
	model    string
	channel  string

	lastResponse Response
	nextUrl      string
}

type Response struct {
	lastResponse Response
	result       ResponseResult
}

type ResponseResult map[string]interface{}

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
