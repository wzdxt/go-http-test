package common

type Receiver interface {
}

type Command interface {
	Execute()
}

type Invoker interface {
	SetCommand(*Command)
	Call()
}

type Client interface {
	Run()
}

type Validator interface {
	Validate(Response)
}
