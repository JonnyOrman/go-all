package command

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"golang.org/x/mod/modfile"
)

type ModExecutor struct {
	config *Config
}

func NewModExecutor(config *Config) *ModExecutor {
	this := new(ModExecutor)

	this.config = config

	return this
}

func (this ModExecutor) Execute(modFile string, directory string) {
	modFileBytes, _ := ioutil.ReadFile(modFile)
	modName := modfile.ModulePath(modFileBytes)
	fmt.Println(this.config.Action + " " + modName)

	goCmd := exec.Command("go", strings.Split(this.config.Args, " ")...)
	goCmd.Dir = directory

	goCmd.Start()

	goCmd.Wait()
}
