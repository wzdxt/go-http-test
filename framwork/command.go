package framework

import (
	"fmt"
)

//command
type RequestCommand struct {
	app *App
}

func (this RequestCommand) Execute() {
	fmt.Println("execute request command on", *this.app)
	//request
	this.app.nextUrl = ""
	//response
	this.app.lastResponse = Response{
		lastResponse: this.app.lastResponse,
		result:       ResponseResult{},
	}
}

//command
type Params map[string]interface{}
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
