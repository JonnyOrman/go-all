package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/mod/modfile"
)

type Command struct {
	action string
	args   string
}

func NewCommand(action string, args string) *Command {
	this := new(Command)

	this.action = action
	this.args = args

	return this
}

func (this Command) ExecuteAllIn(directory string) {
	modFile := directory + "/go.mod"
	_, modFileErr := os.Stat(modFile)

	if modFileErr == nil {
		modFileBytes, _ := ioutil.ReadFile(modFile)
		modName := modfile.ModulePath(modFileBytes)
		fmt.Println(this.action + " " + modName)

		goCmd := exec.Command("go", strings.Split(this.args, " ")...)
		goCmd.Dir = directory

		goCmd.Start()

		goCmd.Wait()
	}

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}

	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			this.ExecuteAllIn(directory + "/" + fileInfo.Name())
		}
	}
}
