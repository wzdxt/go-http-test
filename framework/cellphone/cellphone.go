package cellphone

import (
	"github.com/wzdxt/go-http-test/framework/common"
	"github.com/wzdxt/go-util/array"
)

//receiver
type App struct {
	Platform int
	Version  string
	Os       string
	Model    string
	Channel  string

	Host         string
	Cookies      common.Params
	LastResponse common.Response
	Params       common.Params
	SavedParams  common.Params
	OnlyKeys     []string
	ExceptKeys   []string
}

func (this *App) SetResponse(res common.ResponseResult) {
	this.LastResponse = common.Response{
		Previous: &this.LastResponse,
		Result:   res,
	}
}

func (this *App) SetParam(key string, v common.Anything) {
	this.Params.Set(key, v)
}

func (this App) Save(key string) {
	this.SavedParams.Set(key, this.Params.Get(key))
}

func (this App) Load(key string) {
	this.Params.Set(key, this.SavedParams.Get(key))
}

func (this App) SaveAs(from, to string) {
	this.SavedParams.Set(to, this.Params.Get(from))
}

func (this App) LoadAs(from, to string) {
	this.Params.Set(to, this.SavedParams.Get(from))
}

func (this App) Only(keys []string) {
	this.OnlyKeys = keys
}

func (this App) Except(keys []string) {
	this.ExceptKeys = keys
}

func (this App) GetRequestParams(keys []string) common.Params {
	switch {
	case len(this.OnlyKeys) > 0:
		kArr := array.Intersect(keys, this.OnlyKeys)
		ret := make(common.Params, len(kArr))
		for _, key := range kArr {
			ret[key] = this.Params.Get(key)
		}
		return ret
	default:
		kArr := array.Minus(keys, this.ExceptKeys)
		ret := make(common.Params, len(kArr))
		for _, key := range kArr {
			ret[key] = this.Params.Get(key)
		}
		return ret
	}
}

//invoker
type Customer struct {
	App      *App
	commands []common.Command
}

func (this *Customer) SetCommand(command common.Command) {
	this.commands = append(this.commands, command)
}

func (this Customer) Call() {
	for _, command := range this.commands {
		command.Execute()
	}
}
