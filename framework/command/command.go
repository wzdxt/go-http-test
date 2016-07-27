package command

import (
	"github.com/wzdxt/go-http-test/framework/cellphone"
	"github.com/wzdxt/go-http-test/framework/common"
	"github.com/wzdxt/go-http-test/framework/request"
)

//command
type SetParamsCommand struct {
	App    cellphone.App
	Params common.Params
}

func (this SetParamsCommand) Execute() {
	for k, v := range this.Params {
		this.App.SetParam(k, v)
	}
}

//command
type RequestCommand struct {
	App cellphone.App
	Req request.Request
}

func (this RequestCommand) Execute() {
	resp := this.Req.Request(this.App)
	this.App.SetResponse(resp)
}

//command
type SaveCommand struct {
	App  cellphone.App
	Keys []string
}

func (this SaveCommand) Execute() {
	for _, k := range this.Keys {
		this.App.Save(k)
	}
}

//command
type LoadCommand struct {
	App  cellphone.App
	Keys []string
}

func (this LoadCommand) Execute() {
	for _, k := range this.Keys {
		this.App.Load(k)
	}
}

//command
type SaveAsCommand struct {
	App    cellphone.App
	KeyMap map[string]string
}

func (this SaveAsCommand) Execute() {
	for from, to := range this.KeyMap {
		this.App.SaveAs(from, to)
	}
}

//command
type LoadAsCommand struct {
	App    cellphone.App
	KeyMap map[string]string
}

func (this LoadAsCommand) Execute() {
	for from, to := range this.KeyMap {
		this.App.LoadAs(from, to)
	}
}

//command
type OnlyCommand struct {
	App  cellphone.App
	Keys []string
}

func (this OnlyCommand) Execute() {
	this.App.Only(this.Keys)
}

//command
type ExceptCommand struct {
	App  cellphone.App
	Keys []string
}

func (this ExceptCommand) Execute() {
	this.App.Except(this.Keys)
}

//command
type ValidateCommand struct {
	App       cellphone.App
	Validator common.Validator
}

func (this ValidateCommand) Execute() {
	this.Validator.Validate(this.App.LastResponse)
}
