package request

import (
	"fmt"
	"github.com/wzdxt/go-http-test/framework/cellphone"
	"github.com/wzdxt/go-http-test/framework/common"
)

type Request interface {
	Request(app cellphone.App) common.ResponseResult
	Validate(common.Response)
}

type Json struct {
	Method string
	Uri    string
}

func (this Json) Request(app cellphone.App) common.ResponseResult {
	//request
	url := fmt.Sprintf("https://%s/%s", app.Host, this.Uri)
	fmt.Printf("request %s\n", url)
	//response
	resp := common.ResponseResult{make(common.Params)}
	resp[0] = url
	//cookies

	return resp
}
