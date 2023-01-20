package command

type Config struct {
	Action string
	Args   string
}

func NewConfig(
	action string,
	args string) *Config {
	this := new(Config)

	this.Action = action
	this.Args = args

	return this
}
