package client

import (
	"github.com/wzdxt/go-http-test/framework/cellphone"
	"github.com/wzdxt/go-http-test/framework/command"
	"github.com/wzdxt/go-http-test/framework/common"
	"github.com/wzdxt/go-http-test/framework/request"
)

//client
type Client struct {
	Platform int
	Version  string
	Os       string
	Model    string
	Channel  string

	app      cellphone.App
	customer cellphone.Customer
}

func (this Client) Set(ps common.Params) Client {
	this.customer.SetCommand(command.SetParamsCommand{
		App:    this.GetApp(),
		Params: ps,
	})
	return this
}

func (this Client) Save(keys []string) Client {
	this.customer.SetCommand(command.SaveCommand{
		App:  this.GetApp(),
		Keys: keys,
	})
	return this
}

func (this Client) Load(keys []string) Client {
	this.customer.SetCommand(command.LoadCommand{
		App:  this.GetApp(),
		Keys: keys,
	})
	return this
}

func (this Client) Request(req request.Request) Client {
	this.customer.SetCommand(command.RequestCommand{
		App: this.GetApp(),
		Req: req,
	})
	return this
}

func (this Client) Validate(validator common.Validator) Client {
	this.customer.SetCommand(command.ValidateCommand{
		App:       this.GetApp(),
		Validator: validator,
	})
	return this
}

func (this Client) SaveAs(keyMap map[string]string) Client {
	this.customer.SetCommand(command.SaveAsCommand{
		App:    this.GetApp(),
		KeyMap: keyMap,
	})
	return this
}

func (this Client) LoadAs(keyMap map[string]string) Client {
	this.customer.SetCommand(command.LoadAsCommand{
		App:    this.GetApp(),
		KeyMap: keyMap,
	})
	return this
}

func (this Client) Only(keys []string) Client {
	this.customer.SetCommand(command.OnlyCommand{
		App:  this.GetApp(),
		Keys: keys,
	})
	return this
}

func (this Client) Except(keys []string) Client {
	this.customer.SetCommand(command.ExceptCommand{
		App:  this.GetApp(),
		Keys: keys,
	})
	return this
}

func (this Client) Run() cellphone.App {
	this.GetCustomer().Call()
	return this.app
}

func (this Client) GetCustomer() cellphone.Customer {
	if this.customer == nil {
		this.customer = cellphone.Customer{
			App: this.GetApp(),
		}
	}
	return this.customer
}

func (this Client) GetApp() cellphone.App {
	if this.app == nil {
		this.app = cellphone.App{
			Platform: this.Platform,
			Version:  this.Version,
			Os:       this.Os,
			Model:    this.Model,
			Channel:  this.Channel,
		}
	}
	return this.app
}

func Continue(app cellphone.App) {

}
