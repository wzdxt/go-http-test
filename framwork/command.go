package framework

import (
	"fmt"
)

//command
type RequestCommand struct {
	app    App
	method string
	uri    string
}

func (this RequestCommand) Execute() {
	fmt.Println("execute request command on", this.app)
	//request
	url := fmt.Sprintf("https://%s%s", this.app.Host, this.uri)
	//response
	this.app.SetResponse(ResponseResult{})
	//cookies
}

//command
type SetParamsCommand struct {
	app    App
	params Params
}

func (this SetParamsCommand) Execute() {
	fmt.Println("execute set params command on", this.app, this.params)
}

//command
type SaveCommand struct {
	app  App
	keys []string
}

func (this SaveCommand) Execute() {
	this.app.Save(this.keys)
}

//command
type LoadCommand struct {
	app  App
	keys []string
}

func (this LoadCommand) Execute() {
	this.app.Load(this.keys)
}

//command
type SaveAsCommand struct {
	app    App
	keyMap map[string]string
}

func (this SaveAsCommand) Execute() {
	this.app.SaveAs(this.keyMap)
}

//command
type LoadAsCommand struct {
	app    App
	keyMap map[string]string
}

func (this LoadAsCommand) Execute() {
	this.app.LoadAs(this.keyMap)
}

//command
type OnlyCommand struct {
	app  App
	keys []string
}

func (this OnlyCommand) Execute() {
	this.app.Only(this.keys)
}

//command
type ExceptCommand struct {
	app  App
	keys []string
}

func (this ExceptCommand) Execute() {
	this.app.Except(this.keys)
}

