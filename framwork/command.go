package framework

import (
	"fmt"
)

type Anything interface{}

//command
type RequestCommand struct {
	app *App
}

func (this RequestCommand) Execute() {
	fmt.Println("execute request command on", *this.app)
	//request
	this.app.nextUri = ""
	//response
	this.app.LastResponse = Response{
		lastResponse: &this.app.LastResponse,
		result:       ResponseResult{},
	}
}

//command
type UriCommand struct {
	app *App
	uri string
}

func (this UriCommand) Execute() {
	this.app.nextUri += "/" + this.uri
}

//command
type SetParamsCommand struct {
	app    *App
	params Params
}

func (this SetParamsCommand) Execute() {
	fmt.Println("execute set params command on", *this.app, this.params)
}

//command
type SetParamsFromPreviousCommand struct {
	app     *App
	keyList []string
}

func (this SetParamsFromPreviousCommand) Execute() {
	fmt.Println("execute set params from previous command on", *this.app, this.keyList)
}
