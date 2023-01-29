package command

import (
	"bufio"
	"fmt"
	"os/exec"
)

type ModCmdRunner struct{}

func NewModCmdRunner() *ModCmdRunner {
	this := new(ModCmdRunner)

	return this
}

func (this ModCmdRunner) Run(cmd *exec.Cmd) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	cmd.Wait()
}
