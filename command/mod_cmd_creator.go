package command

import (
	"os/exec"
	"strings"
)

type ModCmdCreator struct {
	config *Config
}

func NewModCmdCreator(
	config *Config,
) *ModCmdCreator {
	this := new(ModCmdCreator)

	this.config = config

	return this
}

func (this ModCmdCreator) Create(modFile string, directory string) *exec.Cmd {
	goCmd := exec.Command("go", strings.Split(this.config.Args, " ")...)
	goCmd.Dir = directory

	return goCmd
}
