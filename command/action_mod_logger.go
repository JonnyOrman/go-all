package command

import "fmt"

type ActionModLogger struct {
	modNameProvider ModNameProvider
	config          *Config
}

func NewActionModLogger(
	modNameProvider ModNameProvider,
	config *Config,
) *ActionModLogger {
	this := new(ActionModLogger)

	this.modNameProvider = modNameProvider
	this.config = config

	return this
}

func (this ActionModLogger) Log(modFile string) {
	modName := this.modNameProvider.Get(modFile)
	fmt.Println(this.config.Action + " " + modName)
}
