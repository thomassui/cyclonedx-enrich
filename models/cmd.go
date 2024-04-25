package models

type Command struct {
	Flag        string
	Description string
	Handler     func(value string) error
	NeedsValue  bool
}

type Commandable interface {
	Commands() []Command
}
